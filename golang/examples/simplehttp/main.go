package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

const (
	UploadFolder  = "uploads"
	FileSizeLimit = 10 * 1024 * 1024 // 10MB
)

type FileuploadResponse struct {
	Name string `json:"file_name"`
	Size int64  `json:"file_size"`
}

func main() {

	os.MkdirAll(UploadFolder, 0o640)

	http.Handle("/upload", uploadFileHandle())
	http.ListenAndServe(":8080", nil)
}

func uploadFileHandle() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "bad method", http.StatusBadRequest)
			return
		}

		reader := http.MaxBytesReader(w, r.Body, FileSizeLimit)
		body, err := io.ReadAll(reader)
		if err != nil {
			fmt.Println("Failed to read file: ", err)
			http.Error(w, "failed to read files", http.StatusBadRequest)
		}

		tmp, err := os.CreateTemp(UploadFolder, "upload-*")
		if err != nil {
			fmt.Println("Failed to create file")
			http.Error(w, "failed to create file", http.StatusInternalServerError)
			return
		}

		fileSize, err := tmp.Write(body)
		if err != nil {
			fmt.Println("Failed to write file")
			http.Error(w, "failed to create file", http.StatusInternalServerError)
			return
		}

		response := &FileuploadResponse{
			Name: tmp.Name(),
			Size: int64(fileSize),
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(response); err != nil {
			fmt.Printf("Failed to write response")
			return
		}

	}
}
