package main

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
)

// 透传struct
type ApiBody struct {
	Url         string `json:"url"`
	Method      string `json:"method"`
	RequestBody string `json:"request_body"`
}

var httpClient *http.Client

func init() {
	httpClient = &http.Client{}
}

// 透传
func realServerApiClient(body *ApiBody, w http.ResponseWriter, r *http.Request) {
	switch {
	case body.Method == "GET":
		req, err := http.NewRequest(http.MethodGet, body.Url, nil)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			io.WriteString(w, "realServerApiClient is failed")
			return
		}
		req.Header = r.Header
		resp, err := httpClient.Do(req)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			io.WriteString(w, "response io failed!")
			return
		}
		rightResponse(w, resp)

	case body.Method == "POST":
		req, err := http.NewRequest(http.MethodPost, body.Url, bytes.NewBuffer([]byte(body.RequestBody)))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			io.WriteString(w, "realServerApiClient is failed")
			return
		}
		req.Header = r.Header
		resp, err := httpClient.Do(req)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			io.WriteString(w, "response io failed!")
			return
		}
		rightResponse(w, resp)

	case body.Method == "DELETE":
		req, err := http.NewRequest(http.MethodDelete, body.Url, nil)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			io.WriteString(w, "realServerApiClient is failed")
			return
		}
		req.Header = r.Header
		resp, err := httpClient.Do(req)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			io.WriteString(w, "response io failed!")
			return
		}
		rightResponse(w, resp)

	case body.Method == "PUT":
		req, err := http.NewRequest(http.MethodPut, body.Url, bytes.NewBuffer([]byte(body.RequestBody)))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			io.WriteString(w, "realServerApiClient is failed")
			return
		}
		req.Header = r.Header
		resp, err := httpClient.Do(req)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			io.WriteString(w, "response io failed!")
			return
		}
		rightResponse(w, resp)
	default:
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, "api has no specify method")
	}
}

func rightResponse(w http.ResponseWriter, response *http.Response) {
	real_server_response, err := ioutil.ReadAll(response.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, "response io failed!")
		return
	}
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, string(real_server_response))
}
