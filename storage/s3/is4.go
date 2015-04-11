package s3

import (
	"github.com/mitchellh/goamz/aws"
	"github.com/mitchellh/goamz/s3"
	"github.com/vanng822/uploader"
	"net/http"
)

type imageStorageS3 struct {
	bucket *s3.Bucket
	ACL    s3.ACL
}

func New(config map[string]string) uploader.ImageStorage {
	regionName := config["regionName"]
	accessKey := config["accessKey"]
	secretKey := config["secretKey"]
	bucketName := config["bucketName"]
	if regionName == "" || accessKey == "" || secretKey == "" || bucketName == "" {
		panic("Config must contain regionName, accessKey, secretKey and bucketName")
	}

	ACL := config["ACL"]

	auth, err := aws.GetAuth(accessKey, secretKey)
	if err != nil {
		panic("Could not get auth")
	}

	client := s3.New(auth, aws.Regions[regionName])
	bucket := client.Bucket(bucketName)

	return NewIS4(bucket, s3.ACL(ACL))
}

func NewIS4(bucket *s3.Bucket, ACL s3.ACL) uploader.ImageStorage {
	s := &imageStorageS3{
		bucket: bucket,
		ACL:    ACL,
	}

	return s
}

func (is4 *imageStorageS3) Put(filename string, imageData []byte) error {
	return is4.bucket.Put(filename, imageData, http.DetectContentType(imageData), is4.ACL)
}

func (is4 *imageStorageS3) Delete(filename string) error {
	return is4.bucket.Del(filename)
}

func (is4 *imageStorageS3) Get(filename string) ([]byte, error) {
	return is4.bucket.Get(filename)
}

func (is4 *imageStorageS3) Exists(filename string) bool {
	res, err := is4.bucket.Head(filename)

	if err != nil {
		return false
	}

	return res.StatusCode == http.StatusOK
}
