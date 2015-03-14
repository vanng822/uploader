package mongodb

import (
	"github.com/vanng822/uploader"
	"gopkg.in/mgo.v2"
	"io/ioutil"
)

var original_session *mgo.Session

func dial(url string) *mgo.Session {
	session, err := mgo.Dial(url)
	if err != nil {
		panic(err)
	}
	return session
}

func getSession(url string) *mgo.Session {
	if original_session == nil {
		original_session = dial(url)
	}
	return original_session.New()
}

type imageStorageMongodb struct {
	url    string
	prefix string
}

func New(config map[string]interface{}) uploader.ImageStorage {
	url := config["url"].(string)
	prefix := config["prefix"].(string)
	if url == "" || prefix == "" {
		panic("You need to configure 'url' with database and 'prefix'")
	}

	return &imageStorageMongodb{
		url:    url,
		prefix: prefix,
	}
}

func (is *imageStorageMongodb) getGridFS(session *mgo.Session) *mgo.GridFS {
	return session.DB("").GridFS(is.prefix)
}

func (is *imageStorageMongodb) Put(filename string, imageData []byte) error {
	session := getSession(is.url)
	defer session.Close()
	gridfs := is.getGridFS(session)
	fd, err := gridfs.Create(filename)
	if err != nil {
		return err
	}
	defer fd.Close()

	_, err = fd.Write(imageData)
	return err
}

func (is *imageStorageMongodb) Delete(filename string) error {
	session := getSession(is.url)
	defer session.Close()
	
	return is.getGridFS(session).Remove(filename)
}

func (is *imageStorageMongodb) Get(filename string) ([]byte, error) {
	session := getSession(is.url)
	defer session.Close()
	gridfs := is.getGridFS(session)

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
	session := getSession(is.url)
	defer session.Close()
	gridfs := is.getGridFS(session)
	
	fd, err := gridfs.Open(filename)
	if err != nil {
		// mgo.ErrNotFound
		return false
	}
	defer fd.Close()
	return true
}
