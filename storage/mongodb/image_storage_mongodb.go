package storage_mongodb

import (
	"github.com/vanng822/uploader"
	"gopkg.in/mgo.v2"
	"io/ioutil"
)

var session *mgo.Session

func dial(url string) *mgo.Session {
	session, err := mgo.Dial(url)
	if err != nil {
		panic(err)
	}
	return session
}

func getSession(url string) *mgo.Session {
	if session == nil {
		session = dial(url)
	}
	return session.New()
}

func getDB(url string) *mgo.Database {
	return getSession(url).DB("")
}

type imageStorageMongodb struct {
	url    string
	prefix string
}

func New(config *uploader.StorageConfig) uploader.ImageStorage {
	url := config.Configurations["url"].(string)
	prefix := config.Configurations["prefix"].(string)
	if url == "" || prefix == "" {
		panic("You need to configure 'url' with database and 'prefix'")
	}

	return &imageStorageMongodb{
		url:    url,
		prefix: prefix,
	}
}

func (is *imageStorageMongodb) getGridFS() *mgo.GridFS {
	return getDB(is.url).GridFS(is.prefix)
}

func (is *imageStorageMongodb) Put(filename string, imageData []byte) error {
	gridfs := is.getGridFS()
	fd, err := gridfs.Create(filename)
	if err != nil {
		return err
	}
	defer fd.Close()

	_, err = fd.Write(imageData)
	return err
}

func (is *imageStorageMongodb) Delete(filename string) error {
	return is.getGridFS().Remove(filename)
}

func (is *imageStorageMongodb) Get(filename string) ([]byte, error) {
	gridfs := is.getGridFS()

	fd, err := gridfs.Open(filename)
	if err != nil {
		return nil, err
	}
	defer fd.Close()
	imageData, err := ioutil.ReadAll(fd)
	if err != nil {
		return nil, err
	}
	return imageData, nil
}

func (is *imageStorageMongodb) Exists(filename string) bool {
	gridfs := is.getGridFS()
	fd, err := gridfs.Open(filename)
	if err != nil {
		// mgo.ErrNotFound
		return false
	}
	defer fd.Close()
	return true
}
