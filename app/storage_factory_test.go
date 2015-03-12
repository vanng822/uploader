package main

import (
	"github.com/stretchr/testify/assert"
	"github.com/vanng822/uploader"
	"testing"
)

func TestStorageFactoryFile(t *testing.T) {
	config := make(map[string]interface{})
	config["directory"] = "./data"
	conf := &uploader.StorageConfig{
		Type:           STORAGE_TYPE_FILE,
		Configurations: config}

	assert.Implements(t, new(ImageStorage), GetStorage(conf))
}

func TestStorageFactoryNotSupported(t *testing.T) {
	config := make(map[string]interface{})
	config["directory"] = "./data"
	conf := &uploader.StorageConfig{
		Type:           "blablabla",
		Configurations: config}

	assert.Panics(t, func() {
		GetStorage(conf)
	})
}

func TestStorageFactoryEmptyDir(t *testing.T) {
	config := make(map[string]interface{})
	config["directory"] = ""
	conf := &uploader.StorageConfig{
		Type:           STORAGE_TYPE_FILE,
		Configurations: config}

	assert.Panics(t, func() {
		GetStorage(conf)
	})
}
