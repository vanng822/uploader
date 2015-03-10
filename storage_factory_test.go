package uploader

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStorageFactoryFile(t *testing.T) {
	assert.Implements(t, new(ImageStorage), GetStorage("file", "./data"))
}

func TestStorageFactoryNotSupported(t *testing.T) {
	assert.Panics(t, func() {
		GetStorage("blablabla", "something")
	})
}
