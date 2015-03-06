package uploader

import (
	"fmt"
	"github.com/nu7hatch/gouuid"
	"net/http"
	"strings"
)

type Uploader struct {
	storage             ImageStorage
	allowedContentTypes map[string]string
}

func NewUploader(storage ImageStorage) *Uploader {
	u := &Uploader{}
	u.storage = storage
	u.allowedContentTypes = map[string]string{"image/jpeg": "jpg", "image/png": "png", "image/gif": "gif"}

	return u
}

func (u *Uploader) SetAllowedImageType(allowedContentTypes map[string]string) {
	if len(allowedContentTypes) == 0 {
		return
	}
	
	allowedTypes := make(map[string]string)

	for t, ext := range allowedContentTypes {
		allowedTypes[strings.ToLower(t)] = ext
	}
	u.allowedContentTypes = allowedTypes
}

func (u *Uploader) Store(imageData []byte) (filename string, err error) {

	contentType := strings.ToLower(http.DetectContentType(imageData))

	ext, allowed := u.allowedContentTypes[contentType]

	if !allowed {
		err = fmt.Errorf("Image type is not allowed")
		return
	}

	uu, err := uuid.NewV4()
	if err != nil {
		return
	}

	filename = fmt.Sprintf("%s.%s", uu.String(), ext)

	err = u.Put(filename, imageData)

	return
}

func (u *Uploader) Put(path string, imageData []byte) error {
	return u.storage.Put(path, imageData)
}
func (u *Uploader) Delete(path string) error {
	return u.storage.Delete(path)
}

func (u *Uploader) Get(path string) ([]byte, error) {
	return u.storage.Get(path)
}
