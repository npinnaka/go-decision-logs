package main

import (
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

//,{"labels":{"id":"f9541d80-0c8d-4a91-9730-b19af3f2f97e","version":"1.2.0"},"decision_id":"daeea5e6-272d-4be0-b569-edb6a549dd77","bundles":{"/bundle/bundle.tar.gz":{}},"path":"authz/allow","input":{"role":"user","state":"TX"},"result":false,"erased":["/input/password"],"requested_by":"192.168.65.1:47265","timestamp":"2025-03-02T23:18:41.443278322Z","metrics":{"counter_server_query_cache_hit":1,"timer_rego_input_parse_ns":60879,"timer_rego_query_eval_ns":50525,"timer_server_handler_ns":142487},"req_id":20}

type DecisionLog struct {
	Labels      map[string]string  `json:"labels"`
	DecisionID  string             `json:"decision_id"`
	Bundles     json.RawMessage    `json:"bundles,omitempty"`
	Path        string             `json:"path,omitempty"`
	Result      bool               `json:"result,omitempty"`
	Erased      []string           `json:"erased,omitempty"`
	RequestedBy string             `json:"requested_by,omitempty"`
	Timestamp   string             `json:"timestamp"`
	Metrics     DecisionLogMetrics `json:"metrics,omitempty"`
	ReqID       int64              `json:"req_id"`
}

type DecisionLogMetrics struct {
	CounterServerQueryCacheHit int64 `json:"counter_server_query_cache_hit"`
	TimerRegoInputParseNS      int64 `json:"timer_rego_input_parse_ns"`
	TimerRegoQueryEvalNS       int64 `json:"timer_rego_query_eval_ns"`
	TimerServerHandlerNS       int64 `json:"timer_server_handler_ns"`
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

	var decisionLogs []DecisionLog
	if err := json.Unmarshal(body, &decisionLogs); err != nil {
		http.Error(w, "Failed to decode logs", http.StatusBadRequest)
		return
	}

	for _, decisionLog := range decisionLogs {
=		fmt.Printf("\n###########################################################")
		fmt.Printf("\nDecision ID %s", decisionLog.DecisionID)
		fmt.Printf("\nLabels %v", decisionLog.Labels)
		fmt.Printf("\nBundles %s", decisionLog.Bundles)
		fmt.Printf("\nPath %s", decisionLog.Path)
		fmt.Printf("\nResult %t", decisionLog.Result)
		fmt.Printf("\nErased %v", decisionLog.Erased)
		fmt.Printf("\nRequested By %s", decisionLog.RequestedBy)
		fmt.Printf("\nTimestamp %s", decisionLog.Timestamp)
		fmt.Printf("\nMetrics %+v", decisionLog.Metrics)
		fmt.Printf("\nReq ID %d", decisionLog.ReqID)
		fmt.Printf("\n###########################################################")

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
