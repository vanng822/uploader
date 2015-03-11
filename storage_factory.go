package uploader

import (
	"fmt"
)

const (
	STORAGE_TYPE_FILE = "file"
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
	default:
		panic(fmt.Sprintf("Unsupported storage type %s", config.Type))
	}
	return storage
}
