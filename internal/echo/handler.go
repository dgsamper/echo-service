package echo

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
)

type echoResponse struct {
	Headers http.Header `json:"headers"`
	Params  url.Values  `json:"params"`
	Body    string      `json:"body"`
	Path    string      `json:"path"`
}

// Handler echoes the request headers, query parameters, body, and path as JSON.
func Handler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "could not read request body"})
		return
	}

	writeJSON(w, http.StatusOK, echoResponse{
		Headers: r.Header,
		Params:  r.URL.Query(),
		Body:    string(body),
		Path:    r.URL.Path,
	})
}

func writeJSON(w http.ResponseWriter, status int, value any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(value); err != nil {
		log.Printf("write response: %v", err)
	}
}
