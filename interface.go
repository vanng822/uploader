package uploader

import (
	
)

type ImageStorage interface {
	Put(filename string, imageData []byte) error
	Delete(filename string) error
	Get(filename string) ([]byte, error)
	Exists(filename string) bool
}
