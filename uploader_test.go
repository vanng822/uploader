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
}