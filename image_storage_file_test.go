package uploader

import (
	"github.com/stretchr/testify/assert"
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