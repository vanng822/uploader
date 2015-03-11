package uploader

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStorageFactoryFile(t *testing.T) {
	config := make(map[string]interface{})
	config["directory"] = "./data"
	conf := &StorageConfig{
		Type:           STORAGE_TYPE_FILE,
		Configurations: config}

	assert.Implements(t, new(ImageStorage), GetStorage(conf))
}

func TestStorageFactoryNotSupported(t *testing.T) {
	config := make(map[string]interface{})
	config["directory"] = "./data"
	conf := &StorageConfig{
		Type:           "blablabla",
		Configurations: config}
	
	assert.Panics(t, func() {
		GetStorage(conf)
	})
}
