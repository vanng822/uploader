// Package uploader for uploading image to storage, retrieve and delete
// 
//	package main
// 
//	import (
//	    "github.com/vanng822/uploader"
//	    "net/http"
//	)
//
//	func main() {
//	    storage := uploader.NewImageStorageFile("./data")
//	    u := uploader.NewUploader(storage)
//	    handler := uploader.UploadHandler(u, "image", "filename")
//	    http.HandleFunc("/upload", handler)
//	    http.ListenAndServe(":8080", nil)
//	} 
//	
//	// curl -i -X POST -F image=@data/kth.jpg http://127.0.0.1:8080/upload
//	// => new file 40a0eb02-1e16-44e2-4694-2db2c159d452.jpg
//	// curl -i -X PUT -F image=@data/kth.jpg http://127.0.0.1:8080/upload?filename=40a0eb02-1e16-44e2-4694-2db2c159d452.jpg
//	// curl -i http://127.0.0.1:8080/upload?filename=40a0eb02-1e16-44e2-4694-2db2c159d452.jpg
//	// curl -i -x DELETE http://127.0.0.1:8080/upload?filename=40a0eb02-1e16-44e2-4694-2db2c159d452.jpg
package uploader
