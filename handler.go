package uploader

import (
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
)

type Handler struct {
	uploader *Uploader
}

func NewHandler(uploader *Uploader) *Handler {
	return &Handler{uploader: uploader}
}

func (h *Handler) HandleGet(res http.ResponseWriter, filename string) {
	if filename == "" {
		res.WriteHeader(http.StatusBadRequest)
		return
	}
	imageData, err := h.uploader.Get(filename)
	if err != nil {
		if !h.uploader.Has(filename) {
			res.WriteHeader(http.StatusNotFound)
			return
		}
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
	res.WriteHeader(http.StatusOK)
	res.Write(imageData)
}

func (h *Handler) HandlePost(res http.ResponseWriter, file multipart.File) {
	imageData, err := ioutil.ReadAll(file)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	filename, err := h.uploader.Store(imageData)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
	res.WriteHeader(http.StatusCreated)
	res.Write([]byte(fmt.Sprintf("{\"status\": \"OK\", \"filename\": \"%s\"}", filename)))
}

func (h *Handler) HandleDelete(res http.ResponseWriter, filename string) {
	if filename == "" {
		res.WriteHeader(http.StatusBadRequest)
		return
	}
	if !h.uploader.Has(filename) {
		res.WriteHeader(http.StatusNotFound)
		return
	}
	err := h.uploader.Delete(filename)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
	res.WriteHeader(http.StatusOK)
	res.Write([]byte("{\"status\": \"OK\"}"))
}

func UploadHandler(uploader *Uploader, uploadField, filenameField string) http.HandlerFunc {
	handler := NewHandler(uploader)
	return func(res http.ResponseWriter, req *http.Request) {
		req.ParseForm()

		switch req.Method {
		case "GET":
			handler.HandleGet(res, req.FormValue(filenameField))
		case "PUT":
			// same for now
			fallthrough
		case "POST":
			file, _, err := req.FormFile(uploadField)
			if err != nil {
				res.WriteHeader(http.StatusBadRequest)
				return
			}
			handler.HandlePost(res, file)
		case "DELETE":
			handler.HandleDelete(res, req.FormValue(filenameField))
		default:
			res.WriteHeader(http.StatusMethodNotAllowed)
		}

	}
}
