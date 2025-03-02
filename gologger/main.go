package main

import (
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type DecisionLog struct {
	Labels struct {
		ID      string `json:"id"`
		Version string `json:"version"`
	} `json:"labels"`
	DecisionID string `json:"decision_id"`
	Path       string `json:"path"`
	Input      struct {
		User string `json:"user"`
	} `json:"input"`
	Result      bool      `json:"result"`
	Erased      []string  `json:"erased"`
	RequestedBy string    `json:"requested_by"`
	Timestamp   time.Time `json:"timestamp"`
	Metrics     struct {
		CounterServerQueryCacheHit int   `json:"counter_server_query_cache_hit"`
		TimerRegoInputParseNs      int64 `json:"timer_rego_input_parse_ns"`
		TimerRegoQueryCompileNs    int64 `json:"timer_rego_query_compile_ns"`
		TimerRegoQueryEvalNs       int64 `json:"timer_rego_query_eval_ns"`
		TimerRegoQueryParseNs      int64 `json:"timer_rego_query_parse_ns"`
		TimerServerHandlerNs       int64 `json:"timer_server_handler_ns"`
	} `json:"metrics"`
	ReqID int `json:"req_id"`
}

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
		log.Printf("Error creating gzip reader: %v", err)
		http.Error(w, "Invalid gzip data", http.StatusBadRequest)
		return
	}
	defer gzReader.Close()

	body, err := io.ReadAll(gzReader)
	if err != nil {
		log.Printf("Error reading decompressed body: %v", err)
		http.Error(w, "Error decompressing data", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	//log.Printf("Received logs: %s", string(body))

	var decisionLogs []DecisionLog
	if err := json.Unmarshal(body, &decisionLogs); err != nil {
		log.Printf("Error unmarshalling JSON: %v", err)
		http.Error(w, "Invalid JSON data", http.StatusBadRequest)
		return
	}

	for _, logEntry := range decisionLogs {
		log.Printf("Received logs: %+v", logEntry)
	}

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
