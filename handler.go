package uploader

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func UploadHandler(uploader *Uploader, uploadField, filenameField string) http.HandlerFunc {

	return func(res http.ResponseWriter, req *http.Request) {
		req.ParseForm()

		switch req.Method {
		case "GET":
			filename := req.FormValue(filenameField)
			if filename == "" {
				res.WriteHeader(http.StatusBadRequest)
				return
			}
			imageData, err := uploader.Get(filename)
			if err != nil {
				if !uploader.Has(filename) {
					res.WriteHeader(http.StatusNotFound)
					return
				}
				res.WriteHeader(http.StatusInternalServerError)
				return
			}
			res.WriteHeader(http.StatusOK)
			res.Write(imageData)
			return
		case "PUT":
			// same for now
			fallthrough
		case "POST":
			file, fileinfo, err := req.FormFile(uploadField)
			if err != nil {
				res.WriteHeader(http.StatusBadRequest)
				return
			}
			fmt.Println(file, fileinfo)
			imageData, err := ioutil.ReadAll(file)
			if err != nil {
				res.WriteHeader(http.StatusInternalServerError)
				return
			}
			
			filename, err := uploader.Store(imageData)
			if err != nil {
				res.WriteHeader(http.StatusInternalServerError)
				return
			}
			res.WriteHeader(http.StatusOK)
			res.Write([]byte(fmt.Sprintf("{\"status\": \"OK\", \"filename\": \"%s\"}", filename)))
			return
		case "DELETE":
			filename := req.FormValue(filenameField)
			if filename == "" {
				res.WriteHeader(http.StatusBadRequest)
				return
			}
			err := uploader.Delete(filename)
			if err != nil {
				if !uploader.Has(filename) {
					res.WriteHeader(http.StatusNotFound)
					return
				}
				res.WriteHeader(http.StatusInternalServerError)
				return
			}
			res.WriteHeader(http.StatusOK)
			res.Write([]byte("{\"status\": \"OK\"}"))
		default:
			res.WriteHeader(http.StatusMethodNotAllowed)
		}

	}
}
