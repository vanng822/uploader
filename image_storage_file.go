package uploader

import (
	"os"
	"fmt"
	"strings"
)

type imageStorageFile struct {
	directory string
}

func NewImageStorageFile(directory string) ImageStorage {
	return &imageStorageFile{
		directory: strings.TrimRight(directory, "/"),
	}
}

func (is *imageStorageFile) makeFilename(filename string) string {
	return fmt.Sprintf("%s/%s", is.directory, filename)
}

func (is *imageStorageFile) Put(filename string, imageData []byte) error {
	fd, err := os.Create(is.makeFilename(filename))
	if err != nil {
		return err
	}
	defer fd.Close()
	
	_, err = fd.Write(imageData)
	return err
}

func (is *imageStorageFile) Delete(filename string) error {
	return os.Remove(is.makeFilename(filename))
}

func (is *imageStorageFile) Get(filename string) ([]byte, error) {
	fd, err := os.Open(is.makeFilename(filename))
	if err != nil {
		return nil, err
	}
	defer fd.Close()
	var imageData []byte
	if _, err = fd.Read(imageData); err != nil {
		return nil, err
	}
	return imageData, nil
}
