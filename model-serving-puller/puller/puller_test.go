// Copyright 2021 IBM Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package puller

import (
	"fmt"
	"path/filepath"
	"testing"

	"github.com/go-logr/logr"
	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/kserve/modelmesh-runtime-adapter/model-serving-puller/generated/mocks"

	"sigs.k8s.io/controller-runtime/pkg/log/zap"
)

var (
	TestDataDir         = "../testdata"
	StorageConfigDir    = filepath.Join(TestDataDir, "/storage-config")
	RootModelDir        = filepath.Join(TestDataDir, "/models")
	StorageKeyTest      = "storage_with_data" // uses file server/testdata/storage-config
	StorageKeyTestEmpty = "myStorage"
	StorageConfigTest   = StorageConfiguration{
		StorageType:     "s3",
		AccessKeyID:     "56421a5768b68451c86546d87654e467",
		SecretAccessKey: "abc643587654876546457985647863547865431786540521",
		EndpointURL:     "https://some-url.appdomain.cloud",
		Region:          "us-south",
		DefaultBucket:   "",
		Certificate:     "random-cert"}
)

func newPullerWithMock(t *testing.T) (*Puller, *mocks.MockS3Downloader, *gomock.Controller) {
	log := zap.New()

	mockCtrl := gomock.NewController(t)
	mockDownloader := mocks.NewMockS3Downloader(mockCtrl)

	rootModelDir, err := filepath.Abs(RootModelDir)
	if err != nil {
		t.Fatalf("Failed to get absolute path of testdata dir %v", err)
	}

	pc := PullerConfiguration{
		RootModelDir:            rootModelDir,
		StorageConfigurationDir: StorageConfigDir,
		S3DownloadConcurrency:   0,
	}
	p := NewPullerFromConfig(log, &pc)
	// inject mock downloader
	p.NewS3DownloaderFromConfig = func(*StorageConfiguration, int, logr.Logger) (S3Downloader, error) { return mockDownloader, nil }
	p.Log = log

	return p, mockDownloader, mockCtrl
}

func Test_GetS3Downloader_AddDownloaderToCache_AndNoConfigChange(t *testing.T) {
	expectedDownloader := s3Downloader{config: &StorageConfigTest}

	p, _, mockCtrl := newPullerWithMock(t)
	p.NewS3DownloaderFromConfig = func(config *StorageConfiguration, downloadConcurrency int, log logr.Logger) (S3Downloader, error) {
		return &s3Downloader{config: config}, nil
	}
	defer mockCtrl.Finish()

	// StorageConfig not needed because will pull from disk on first pass
	downloader, err := p.getS3Downloader(StorageKeyTest, nil)

	assert.Nil(t, err)
	assert.Equal(t, &expectedDownloader, downloader)
	assert.True(t, downloader.IsSameConfig(&StorageConfigTest))
	assert.Equal(t, 1, len(p.s3DownloaderCache))
	assert.Equal(t, downloader, p.s3DownloaderCache[StorageKeyTest])

	downloader2, err := p.getS3Downloader(StorageKeyTest, &StorageConfigTest)
	assert.Nil(t, err)
	assert.Same(t, downloader, downloader2)
	assert.Equal(t, &expectedDownloader, downloader2)
	assert.Equal(t, 1, len(p.s3DownloaderCache))
}

func Test_GetS3Downloader_StorageConfigChange(t *testing.T) {
	StorageConfigTestDiffAccessKey := StorageConfiguration{
		StorageType:     "s3",
		AccessKeyID:     "46421a5768b68451c86546d876542369",
		SecretAccessKey: "abc643587654876546457985647863547865431786540521",
		EndpointURL:     "https://some-url.appdomain.cloud",
		Region:          "us-south",
		DefaultBucket:   ""}

	firstDownloader := s3Downloader{config: &StorageConfigTestDiffAccessKey}
	expectedDownloader := s3Downloader{config: &StorageConfigTest}

	p, _, mockCtrl := newPullerWithMock(t)
	p.NewS3DownloaderFromConfig = func(config *StorageConfiguration, downloadConcurrency int, log logr.Logger) (S3Downloader, error) {
		return &s3Downloader{config: config}, nil
	}
	defer mockCtrl.Finish()

	p.s3DownloaderCache[StorageKeyTest] = &firstDownloader
	assert.Equal(t, 1, len(p.s3DownloaderCache))
	assert.Equal(t, &firstDownloader, p.s3DownloaderCache[StorageKeyTest])

	// overwrites existing downloader in cache
	downloader, err := p.getS3Downloader(StorageKeyTest, &StorageConfigTest)

	assert.Nil(t, err)
	assert.Equal(t, &expectedDownloader, downloader)
	assert.NotEqual(t, downloader, firstDownloader)
	assert.False(t, downloader.IsSameConfig(&StorageConfigTestDiffAccessKey))
	assert.Equal(t, 1, len(p.s3DownloaderCache))
	assert.Equal(t, &expectedDownloader, p.s3DownloaderCache[StorageKeyTest])
}

func Test_GetS3Downloader_MultipleStorageKeys(t *testing.T) {
	expectedDownloader := s3Downloader{config: &StorageConfigTest}

	p, _, mockCtrl := newPullerWithMock(t)
	p.NewS3DownloaderFromConfig = func(config *StorageConfiguration, downloadConcurrency int, log logr.Logger) (S3Downloader, error) {
		return &s3Downloader{config: config}, nil
	}
	defer mockCtrl.Finish()

	// pulls StorageConfig from disk based on key name
	downloader, err := p.getS3Downloader(StorageKeyTest, nil)

	assert.Nil(t, err)
	assert.Equal(t, &expectedDownloader, downloader)
	assert.Equal(t, 1, len(p.s3DownloaderCache))

	// pulls StorageConfig from disk based on key name
	secondDownloader, err := p.getS3Downloader(StorageKeyTestEmpty, nil)
	assert.Nil(t, err)
	assert.Equal(t, 2, len(p.s3DownloaderCache))
	assert.NotEqual(t, &expectedDownloader, secondDownloader)
}

func Test_GetS3Downloader_Error(t *testing.T) {
	expectedError := fmt.Errorf("Invalid config")

	p, _, mockCtrl := newPullerWithMock(t)
	defer mockCtrl.Finish()

	p.NewS3DownloaderFromConfig = func(*StorageConfiguration, int, logr.Logger) (S3Downloader, error) { return nil, expectedError }

	downloader, err := p.getS3Downloader(StorageKeyTest, &StorageConfigTest)
	assert.Nil(t, downloader)
	assert.Equal(t, expectedError, err)
	assert.Equal(t, 0, len(p.s3DownloaderCache))
}

func Test_GetS3Downloader_ErrorDeletesCache(t *testing.T) {
	expectedError := fmt.Errorf("Invalid config")
	emptyDownloader := s3Downloader{config: &StorageConfiguration{}}

	p, _, mockCtrl := newPullerWithMock(t)
	defer mockCtrl.Finish()

	p.NewS3DownloaderFromConfig = func(*StorageConfiguration, int, logr.Logger) (S3Downloader, error) { return nil, expectedError }
	p.s3DownloaderCache[StorageKeyTest] = &emptyDownloader
	assert.Equal(t, 1, len(p.s3DownloaderCache))

	downloader, err := p.getS3Downloader(StorageKeyTest, &StorageConfigTest)
	assert.Nil(t, downloader)
	assert.Equal(t, expectedError, err)
	assert.Equal(t, 0, len(p.s3DownloaderCache))
}

func Test_CleanCache_DeletesFakeKeys(t *testing.T) {
	emptyDownloader := s3Downloader{config: &StorageConfiguration{}}

	p, _, mockCtrl := newPullerWithMock(t)
	defer mockCtrl.Finish()

	p.s3DownloaderCache["fake-key"] = &emptyDownloader
	p.s3DownloaderCache["another-fake-key"] = &emptyDownloader
	p.s3DownloaderCache[StorageKeyTest] = &emptyDownloader
	p.s3DownloaderCache[StorageKeyTestEmpty] = &emptyDownloader
	assert.Equal(t, 4, len(p.s3DownloaderCache))

	p.CleanCache()
	assert.Equal(t, 2, len(p.s3DownloaderCache))
}
