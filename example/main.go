package main

import (
	"github.com/vanng822/uploader"
	"net/http"
)

func main() {
	storage := uploader.NewImageStorageFile("./data")
	u := uploader.NewUploader(storage)
	handler := uploader.UploadHandler(u, "file", "filename")
	http.HandleFunc("/upload", handler)
	http.ListenAndServe(":8080", nil)
} 

