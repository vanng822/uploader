package uploader

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestNewImageStorageFile(t *testing.T) {
	s := NewImageStorageFile("data")
	assert.Implements(t, new(ImageStorage), s)
} 

func TestImageStorageFile(t *testing.T) {
	s := imageStorageFile{directory: "data"}
	assert.Equal(t, s.makeFilename("testing.jpg"), "data/testing.jpg") 
}

func TestPutDelete(t *testing.T) {
	s := imageStorageFile{directory: "./data"}
	filename := testGetFilename()
	imageData := testGetImageByte()
	assert.Nil(t, s.Put(filename, imageData))
	assert.Nil(t, s.Delete(filename))
	_, err := os.Stat(s.makeFilename(filename))
	assert.NotNil(t, err)
}

func TestGet(t *testing.T) {
	s := imageStorageFile{directory: "./data"}
	imageData, err := s.Get("kth.jpg")
	assert.Nil(t, err)
	assert.NotNil(t, imageData)
	assert.Equal(t, len(imageData), 29429)
}

func TestHas(t *testing.T) {
	s := imageStorageFile{directory: "./data"}
	assert.True(t, s.Exists("kth.jpg"))
	assert.False(t, s.Exists("blabla.jpg"))
}