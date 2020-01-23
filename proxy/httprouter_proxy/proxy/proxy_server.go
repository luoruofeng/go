package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"net/url"
	"text/template"

	"github.com/julienschmidt/httprouter"
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

func RegisterHandler() *httprouter.Router {
	r := httprouter.New()
	// 代理
	r.GET("/hellog", onlyOneUrlUseProxy)
	// 透传
	r.POST("/rs-api", realServerApi)

	// 代理视频上传和下载
	r.GET("/upload", uploadPageHandle)
	r.GET("/videos/:video_name", onlyOneUrlUseProxy)
	r.POST("/upload/:video_name", onlyOneUrlUseProxy)
	return r
}

func uploadPageHandle(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	t, _ := template.ParseFiles("./tmpl/upload.html")
	t.Execute(w, map[string]interface{}{"title": "upload video", "videoname": "testvideo"})
}

//测试代理
// curl -XGET localhost:8080/hellog
// curl -XPOST localhost:8080/hellop -d "name=luo&age=11"
// curl -XPOST localhost:8080/hellop  -H "Content-Type: application/json" -d '{"name":"luo","age":111}'

//测试透传
// curl -XPOST localhost:8080/rs-api  -H "Content-Type: application/json" -d '{"url":"http://localhost:9000/hellop","method":"POST","request_body":"{\"name\":\"luo\",\"age\":111}"}'

//测试视频上传 http://localhost:8080/upload
//视频播放 http://localhost:8080/videos/testvideo
func main() {
	r := RegisterHandler()

	// 测试代理单个url和透传
	http.ListenAndServe(":8080", r)

	// 测试代理所有url
	// http.ListenAndServe(":8080", NewProxyHandler())
}
