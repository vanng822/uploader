package mongodb

import (
	"github.com/stretchr/testify/assert"
	"github.com/vanng822/uploader"
	"testing"
)

func TestNewImageStorageFile(t *testing.T) {
	config := map[string]interface{}{
		"url":    testUrl,
		"prefix": testPrefix,
	}
	s := New(config)
	assert.Implements(t, new(uploader.ImageStorage), s)
}

func TestPutDelete(t *testing.T) {
	s := imageStorageMongodb{
		url:    testUrl,
		prefix: testPrefix,
	}
	filename := testGetFilename()
	imageData := testGetImageByte()
	assert.Nil(t, s.Put(filename, imageData))
	assert.Nil(t, s.Delete(filename))
}

func TestGet(t *testing.T) {
	s := imageStorageMongodb{
		url:    testUrl,
		prefix: testPrefix,
	}
	filename := testGetFilename()
	imageData := testGetImageByte()
	assert.Nil(t, s.Put(filename, imageData))

	imageData, err := s.Get(filename)
	assert.Nil(t, err)
	assert.NotNil(t, imageData)
	assert.Equal(t, len(imageData), 29429)

	assert.Nil(t, s.Delete(filename))
}

func TestHas(t *testing.T) {
	s := imageStorageMongodb{
		url:    testUrl,
		prefix: testPrefix,
	}
	filename := testGetFilename()
	imageData := testGetImageByte()
	assert.Nil(t, s.Put(filename, imageData))

	assert.True(t, s.Exists(filename))
	assert.False(t, s.Exists("blabla.jpg"))
	assert.Nil(t, s.Delete(filename))
}
