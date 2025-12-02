package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

type FileResponse struct {
	Name string `json:"file_name"`
	Len  int    `json:"length"`
}

const (
	upload_dir    = "uploads"
	max_file_size = 10 * 1024 * 1024
)

func main() {
	if err := os.MkdirAll(upload_dir, 0o755); err != nil {
		fmt.Println("Failed to create dir")
	}

	http.Handle("/upload", uploadHandler())

	http.ListenAndServe(":8080", nil)

}
func uploadHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		reader := http.MaxBytesReader(w, r.Body, max_file_size)
		body, err := io.ReadAll(reader)
		if err != nil {
			fmt.Println("failed to read body")
			http.Error(w, "failed to create temp file", http.StatusBadRequest)
			return
		}

		temp, err := os.CreateTemp(upload_dir, "uploads")
		if err != nil {
			fmt.Println("failed to create temp file")
			http.Error(w, "failed to create temp file", http.StatusInternalServerError)
			return
		}
		if _, err := temp.Write(body); err != nil {
			fmt.Println("failed to write files")
			http.Error(w, "failed to write files", http.StatusInternalServerError)
			return
		}

		response := &FileResponse{
			Name: filepath.Base(temp.Name()),
			Len:  len(body),
		}

		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(response); err != nil {
			fmt.Println("Failed to write response")
			return
		}
	}
}
