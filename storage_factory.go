package uploader

import (
	"fmt"
)

const (
	STORAGE_TYPE_FILE    = "file"
	STORAGE_TYPE_MONGODB = "mongodb"
)

type StorageConfig struct {
	Type           string
	Configurations map[string]interface{}
}

func GetStorage(config *StorageConfig) ImageStorage {
	var storage ImageStorage
	switch config.Type {
	case STORAGE_TYPE_FILE:
		// can convert to struct if more complex config
		directory, ok := config.Configurations["directory"]
		dir := directory.(string)
		if !ok || dir == "" {
			panic("File storage configuration needs to have a directory")
		}
		storage = NewImageStorageFile(dir)
	case STORAGE_TYPE_MONGODB:
		url := config.Configurations["url"].(string)
		prefix := config.Configurations["prefix"].(string)
		if url == "" || prefix == "" {
			panic("You need to configure 'url' with database and 'prefix'")
		}
		storage = NewImageStorageMongodb(prefix, url)
	default:
		panic(fmt.Sprintf("Unsupported storage type %s", config.Type))
	}
	return storage
}
