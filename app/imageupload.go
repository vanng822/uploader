package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/go-martini/martini"
	"github.com/vanng822/uploader"
	"github.com/vanng822/uploader/storage/mongodb"
	"net/http"
	"os"
)

const (
	STORAGE_TYPE_FILE    = "file"
	STORAGE_TYPE_MONGODB = "mongodb"
)

type StorageConfig struct {
	Type           string
	Configurations map[string]interface{}
}

func GetStorage(config *uploader.StorageConfig) uploader.ImageStorage {
	var storage uploader.ImageStorage
	switch config.Type {
	case STORAGE_TYPE_FILE:
		storage = uploader.NewImageStorageFile(config.Configurations)
	case STORAGE_TYPE_MONGODB:
		storage = storage_mongodb.New(config.Configurations)
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

	m := martini.Classic()

	for _, endpoint := range conf.Endpoints {

		go func(endpoint *EndpointConfig) {
			u := uploader.NewUploader(GetStorage(endpoint.Storage))
			handler := uploader.NewHandler(u)

			m.Group(endpoint.Endpoint, func(r martini.Router) {
				r.Get("/:filename", func(res http.ResponseWriter, req *http.Request, params martini.Params) {
					handler.HandleGet(res, params["filename"])
				})
				r.Post("/", func(res http.ResponseWriter, req *http.Request, params martini.Params) {
					file, _, err := req.FormFile(endpoint.FileField)
					if err != nil {
						res.WriteHeader(http.StatusBadRequest)
						return
					}
					handler.HandlePost(res, file)
				})
				r.Put("/:filename", func(res http.ResponseWriter, req *http.Request, params martini.Params) {
					file, _, err := req.FormFile(endpoint.FileField)
					if err != nil {
						res.WriteHeader(http.StatusBadRequest)
						return
					}
					handler.HandlePut(res, file, params["filename"])
				})
				r.Delete("/:filename", func(res http.ResponseWriter, req *http.Request, params martini.Params) {
					handler.HandleDelete(res, params["filename"])
				})
			})
		}(endpoint)
	}

	http.ListenAndServe(fmt.Sprintf("%s:%d", conf.Host, conf.Port), m)
}
