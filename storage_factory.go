package uploader

import (
	"fmt"
)

const (
	STORAGE_TYPE_FILE = "file"
)

func GetStorage(storageType, configuration string) ImageStorage {
	var storage ImageStorage
	switch storageType {
	case STORAGE_TYPE_FILE:
		storage = NewImageStorageFile(configuration)
	default:
		panic(fmt.Sprintf("Unsupported storage type %s", storageType))
	}
	return storage
}
