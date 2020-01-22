package main

import (
	"bytes"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"net/url"
)

const (
	REAL_SERVER = "http://127.0.0.1:9000"
)

/*---------对所有url进行代理----------*/
type ProxyHandler struct {
	realServer string
}

func NewProxyHandler() http.Handler {
	handler := ProxyHandler{
		realServer: REAL_SERVER,
	}
	return handler
}

func (mw ProxyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	realUrl, err := url.Parse(mw.realServer)
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		io.WriteString(w, "Real Server URL is wrong!")
		panic(1)
	}
	proxy := httputil.NewSingleHostReverseProxy(realUrl)
	proxy.ServeHTTP(w, r)
}

/*---------对某个url进行代理----------*/
func onlyOneUrlUseProxy(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	realUrl, err := url.Parse(REAL_SERVER)
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		io.WriteString(w, "Real Server URL is wrong!")
		panic(1)
	}
	proxy := httputil.NewSingleHostReverseProxy(realUrl)
	proxy.ServeHTTP(w, r)
}

func realServerApi(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	request_body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		io.WriteString(w, "read request body failed")
		return
	}
	apiBody := &ApiBody{}
	json_err := json.Unmarshal(request_body, apiBody)
	if json_err != nil {
		w.WriteHeader(http.StatusBadGateway)
		io.WriteString(w, "request body unmarshal to object is failed")
		return
	}
	realServerApiClient(apiBody, w, r)
}

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

func RegisterHandler() *httprouter.Router {
	r := httprouter.New()
	// 代理
	r.GET("/hellog", onlyOneUrlUseProxy)
	// 透传
	r.POST("/rs-api", realServerApi)
	return r
}

//测试代理
// curl -XGET localhost:8080/hellog
// curl -XPOST localhost:8080/hellop -d "name=luo&age=11"
// curl -XPOST localhost:8080/hellop  -H "Content-Type: application/json" -d '{"name":"luo","age":111}'

//测试透传
// curl -XPOST localhost:8080/rs-api  -H "Content-Type: application/json" -d '{"url":"http://localhost:9000/hellop","method":"POST","request_body":"{\"name\":\"luo\",\"age\":111}"}'
func main() {
	r := RegisterHandler()
	// 测试代理单个url和透传
	http.ListenAndServe(":8080", r)

	// 测试代理所有url
	// http.ListenAndServe(":8080", NewProxyHandler())
}
