package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	uploadDir := "uploads"
	if err := os.MkdirAll(uploadDir, 0o755); err != nil {
		log.Fatalf("failed to create upload dir: %v", err)
	}

	http.Handle("/upload", uploadHandler(uploadDir))

	log.Println("listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func uploadHandler(uploadDir string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// Limit body to 10MB to avoid abuse.
		r.Body = http.MaxBytesReader(w, r.Body, 10*1024*1024)
		data, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "failed to read body", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		tmp, err := os.CreateTemp(uploadDir, "upload-*")
		if err != nil {
			http.Error(w, "failed to create file", http.StatusInternalServerError)
			return
		}
		defer tmp.Close()

		if _, err := tmp.Write(data); err != nil {
			http.Error(w, "failed to write file", http.StatusInternalServerError)
			return
		}

		resp := map[string]any{
			"path":  filepath.Join(uploadDir, filepath.Base(tmp.Name())),
			"bytes": len(data),
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			http.Error(w, "failed to encode response", http.StatusInternalServerError)
			return
		}
	}
}
