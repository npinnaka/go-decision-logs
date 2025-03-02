package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"net/http"
)

func logsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost || r.URL.Path != "/logs" {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}
	if r.Header.Get("Content-Encoding") != "gzip" {
		http.Error(w, "Only gzip-encoded content is accepted", http.StatusUnsupportedMediaType)
		return
	}
	gzReader, err := gzip.NewReader(r.Body)
	if err != nil {
		log.Printf("Gzip error: %v", err)
		http.Error(w, "Invalid gzip data", http.StatusBadRequest)
		return
	}
	defer gzReader.Close()
	body, err := io.ReadAll(gzReader)
	if err != nil {
		log.Printf("Decompression error: %v", err)
		http.Error(w, "Error decompressing data", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	log.Printf("Received logs: %s", string(body))
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{"status": "success"}`)
}

func main() {
	http.HandleFunc("/logs", logsHandler)
	log.Println("Starting logging server on :3001")
	err := http.ListenAndServe(":3001", nil)
	if err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
