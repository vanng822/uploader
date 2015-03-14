package uploader

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewUploader(t *testing.T) {
	u := NewUploader(nil)
	assert.Equal(t, u.storage, nil)
	assert.Equal(t, u.allowedContentTypes, map[string]string{"image/jpeg": "jpg", "image/png": "png", "image/gif": "gif"})
}

func TestSetAllowedImageType(t *testing.T) {
	u := NewUploader(nil)
	u.SetAllowedImageType(map[string]string{"image/jpeg": "jpg"})
	assert.Equal(t, u.allowedContentTypes, map[string]string{"image/jpeg": "jpg"})

	u = NewUploader(nil)
	u.SetAllowedImageType(map[string]string{})
	assert.Equal(t, u.allowedContentTypes, map[string]string{"image/jpeg": "jpg", "image/png": "png", "image/gif": "gif"})
}

func TestUploaderStore(t *testing.T) {
	config := map[string]interface{}{
		"directory": "./data",
	}
	u := NewUploader(NewImageStorageFile(config))
	filename, err := u.Store(testGetImageByte())
	assert.Nil(t, err)
	assert.NotNil(t, filename)
	assert.Nil(t, u.Delete(filename))
}

func TestUploaderStoreFileTypeNotAllowed(t *testing.T) {
	config := map[string]interface{}{
		"directory": "./data",
	}
	u := NewUploader(NewImageStorageFile(config))
	u.SetAllowedImageType(map[string]string{"image/png": "png", "image/gif": "gif"})
	filename, err := u.Store(testGetImageByte())
	assert.NotNil(t, err)
	assert.Empty(t, filename)
	assert.Error(t, err, "Image type is not allowed")
}

func TestUploaderGet(t *testing.T) {
	config := map[string]interface{}{
		"directory": "./data",
	}
	u := NewUploader(NewImageStorageFile(config))
	imageData, err := u.Get("kth.jpg")
	assert.Nil(t, err)
	assert.NotNil(t, imageData)
	assert.Equal(t, len(imageData), 29429)
}
