# uploader

Upload image to storage, retrieve and delete.
Bundle with File storage but any storage implements
ImageStorage should works, please create pull request
if you implement one.

## http request methods
### GET
for getting image
### POST/PUT
for creating new image
### DELETE
for deleting an image

## Examle

Take look in "app" if you want something more ready to run

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
    
    // curl -i -X POST -F file=@data/kth.jpg http://127.0.0.1:8080/upload
    // curl -i -X PUT -F file=@data/kth.jpg http://127.0.0.1:8080/upload?filename=40a0eb02-1e16-44e2-4694-2db2c159d452.jpg
    // curl -i http://127.0.0.1:8080/upload?filename=40a0eb02-1e16-44e2-4694-2db2c159d452.jpg
