package main

import (
	"fmt"
	"github.com/vanng822/uploader"
	"github.com/vanng822/uploader/storage/mongodb"
	"github.com/vanng822/uploader/storage/file"
)

const (
	STORAGE_TYPE_FILE    = "file"
	STORAGE_TYPE_MONGODB = "mongodb"
)

func GetStorage(config *StorageConfig) ImageStorage {
	var storage ImageStorage
	switch config.Type {
	case STORAGE_TYPE_FILE:
		storage = storage_file.New(config)
	case STORAGE_TYPE_MONGODB:
		storage = storage_mongodb.New(config)
	default:
		panic(fmt.Sprintf("Unsupported storage type %s", config.Type))
	}
	return storage
}
