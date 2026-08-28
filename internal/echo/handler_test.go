package echo_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"strings"
	"testing"
	"testing/iotest"

	"github.com/dgsamper/echo-service/internal/echo"
)

type response struct {
	Headers http.Header `json:"headers"`
	Params  url.Values  `json:"params"`
	Body    string      `json:"body"`
	Path    string      `json:"path"`
}

func TestHandlerEchoesRequest(t *testing.T) {
	body := "{\"message\":\"hello\"}\n"
	req := httptest.NewRequest(http.MethodPost, "/orders/42?tag=go&tag=devops", strings.NewReader(body))
	req.Header.Add("X-Example", "first")
	req.Header.Add("X-Example", "second")
	rec := httptest.NewRecorder()
	echo.Handler(rec, req)

	if rec.Code != http.StatusOK || rec.Header().Get("Content-Type") != "application/json" {
		t.Fatalf("unexpected response: status=%d, headers=%v", rec.Code, rec.Header())
	}
	var got response
	if err := json.Unmarshal(rec.Body.Bytes(), &got); err != nil {
		t.Fatal(err)
	}
	want := response{
		Headers: http.Header{"X-Example": {"first", "second"}},
		Params:  url.Values{"tag": {"go", "devops"}},
		Body:    body,
		Path:    "/orders/42",
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("response = %#v, want %#v", got, want)
	}
}

func TestHandlerEmptyRequest(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	echo.Handler(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusOK)
	}
	var got response
	if err := json.Unmarshal(rec.Body.Bytes(), &got); err != nil {
		t.Fatal(err)
	}
	want := response{Headers: http.Header{}, Params: url.Values{}, Path: "/"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("response = %#v, want %#v", got, want)
	}
}

func TestHandlerBodyReadError(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/", iotest.ErrReader(errors.New("read failed")))
	rec := httptest.NewRecorder()
	echo.Handler(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Errorf("status = %d, want %d", rec.Code, http.StatusBadRequest)
	}
	var got map[string]string
	if err := json.Unmarshal(rec.Body.Bytes(), &got); err != nil {
		t.Fatal(err)
	}
	if got["error"] != "could not read request body" {
		t.Errorf("unexpected error response: %v", got)
	}
}
