package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/vanng822/r2router"
	"github.com/vanng822/uploader"
	storage_mongodb "github.com/vanng822/uploader/storage/mongodb"
	"github.com/vanng822/uploader/storage/s3"
	"net/http"
	"os"
)

const (
	STORAGE_TYPE_FILE    = "file"
	STORAGE_TYPE_MONGODB = "mongodb"
	STORAGE_TYPE_S3      = "s3"
)

type StorageConfig struct {
	Type           string
	Configurations map[string]string
}

// Configuration format
// {
// 	"Host": "127.0.0.1",
// 	"Port": 8080,
// 	"Endpoints": [{
// 		"Endpoint": "/storage",
// 		"FileField": "image",
// 		"Storage": {
// 			"Type": "file",
// 			"Configurations": {
// 				"directory": "./data"
// 			}
// 		}
// 	},{
// 		"Endpoint": "/mongodb",
// 		"FileField": "image",
// 		"Storage": {
// 			"Type": "mongodb",
// 			"Configurations": {
// 				"url": "localhost:27017/uploader",
// 				"prefix": "test_uploader"
// 			}
// 		}
// 	}, {
// 		"Endpoint": "/s3",
// 		"FileField": "image",
// 		"Storage": {
// 			"Type": "s3",
// 			"Configurations": {
// 				"regionName": "eu-west-1",
// 				"accessKey":  "your_access_key",
// 				"secretKey":  "your_secret_key",
// 				"bucketName": "some_bucket"
// 			}
// 		}
// 	}]
// }
func GetStorage(config *StorageConfig) uploader.ImageStorage {
	var storage uploader.ImageStorage
	switch config.Type {
	case STORAGE_TYPE_FILE:
		storage = uploader.NewImageStorageFile(config.Configurations)
	case STORAGE_TYPE_MONGODB:
		storage = storage_mongodb.New(config.Configurations)
	case STORAGE_TYPE_S3:
		storage = s3.New(config.Configurations)
	default:
		panic(fmt.Sprintf("Unsupported storage type %s", config.Type))
	}
	return storage
}

type EndpointConfig struct {
	Endpoint  string
	FileField string
	Storage   *StorageConfig
}

type Config struct {
	Host      string
	Port      int
	Endpoints []*EndpointConfig
}

func LoadConfig(filename string) *Config {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	conf := Config{}
	err = decoder.Decode(&conf)
	if err != nil {
		panic(err)
	}
	return &conf
}

func main() {

	var (
		config string
		host   string
		port   int
	)

	flag.StringVar(&host, "h", "", "Host to bind to")
	flag.IntVar(&port, "p", 0, "Port number to listen on")
	flag.StringVar(&config, "c", "./config/app.json", "Path to configurations")
	flag.Parse()

	conf := LoadConfig(config)

	if host != "" {
		conf.Host = host
	}

	if port != 0 {
		conf.Port = port
	}

	if len(conf.Endpoints) == 0 {
		panic("There is no endpoint configured")
	}

	app := r2router.NewSeeforRouter()

	for _, endpoint := range conf.Endpoints {

		go func(endpoint *EndpointConfig) {
			u := uploader.NewUploader(GetStorage(endpoint.Storage))
			handler := uploader.NewHandler(u)

			app.Group(endpoint.Endpoint, func(r *r2router.GroupRouter) {
				r.Get("/:filename", func(res http.ResponseWriter, req *http.Request, params r2router.Params) {
					handler.HandleGet(res, params.Get("filename"))
				})
				r.Post("/", func(res http.ResponseWriter, req *http.Request, params r2router.Params) {
					file, _, err := req.FormFile(endpoint.FileField)
					if err != nil {
						res.WriteHeader(http.StatusBadRequest)
						return
					}
					handler.HandlePost(res, file)
				})
				r.Put("/:filename", func(res http.ResponseWriter, req *http.Request, params r2router.Params) {
					file, _, err := req.FormFile(endpoint.FileField)
					if err != nil {
						res.WriteHeader(http.StatusBadRequest)
						return
					}
					handler.HandlePut(res, file, params.Get("filename"))
				})
				r.Delete("/:filename", func(res http.ResponseWriter, req *http.Request, params r2router.Params) {
					handler.HandleDelete(res, params.Get("filename"))
				})
			})
		}(endpoint)
	}

	http.ListenAndServe(fmt.Sprintf("%s:%d", conf.Host, conf.Port), app)
}
