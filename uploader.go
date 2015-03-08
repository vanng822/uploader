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

func (u *Uploader) typeValid(imageData []byte) (bool, string) {
	contentType := strings.ToLower(http.DetectContentType(imageData))

	ext, allowed := u.allowedContentTypes[contentType]
	
	return allowed, ext
}

func (u *Uploader) Store(imageData []byte) (filename string, err error) {

	allowed, ext := u.typeValid(imageData)

	if !allowed {
		err = fmt.Errorf("Image type is not allowed")
		return
	}

	uu, err := uuid.NewV4()
	if err != nil {
		return
	}

	filename = fmt.Sprintf("%s.%s", uu.String(), ext)

	err = u.storage.Put(filename, imageData)

	return
}

func (u *Uploader) Put(filename string, imageData []byte) error {
	allowed, _ := u.typeValid(imageData)
	
	if !allowed {
		return fmt.Errorf("Image type is not allowed")
	}
	
	return u.storage.Put(filename, imageData)
}
func (u *Uploader) Delete(filename string) error {
	return u.storage.Delete(filename)
}

func (u *Uploader) Get(filename string) ([]byte, error) {
	return u.storage.Get(filename)
}

func (u *Uploader) Exists(filename string) bool {
	return u.storage.Exists(filename)
}