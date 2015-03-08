package uploader

import (
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
				res.WriteHeader(http.StatusInternalServerError)
				return
			}
			res.WriteHeader(http.StatusOK)
			res.Write(imageData)
			return
		case "POST":
			//file, fileinfo, err := req.FormFile(uploadField)
			
		case "PUT":

		case "DELETE":
			filename := req.FormValue(filenameField)
			if filename == "" {
				res.WriteHeader(http.StatusBadRequest)
				return
			}
			err := uploader.Delete(filename)
			if err != nil {
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
