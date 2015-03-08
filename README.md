# uploader

Upload image to storage, retrieve and delete.
Bundle with File storage but any storage implements
ImageStorage shold works, please create pull request
if you implement one.

## http request methods
### GET
for getting image
### POST/PUT
for creating new image
### DELETE
for deleting an image

## Examle
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
    
    // curl -i -X PUT -F file=@data/kth.jpg http://127.0.0.1:8080/upload
