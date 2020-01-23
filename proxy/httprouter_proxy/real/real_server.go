package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/julienschmidt/httprouter"
)

func helloHandle(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	fmt.Println("-----------Request Header---------")
	for k, v := range r.Header {
		for ind, _ := range v {
			fmt.Println(k + " : " + v[ind])
		}
	}

	fmt.Println("-----------Request Body---------")
	rb, _ := ioutil.ReadAll(r.Body)
	fmt.Println(string(rb))

	w.WriteHeader(http.StatusOK)
	io.WriteString(w, "Hello world")
}

const (
	VIDEO_DIR = "./video/"

	// 100MB
	MAX_VIDEO = 1024 * 1024 * 100
)

func streamHandle(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	videoName := params.ByName("video_name")
	file, err := os.Open(VIDEO_DIR + videoName)
	defer file.Close()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, "video path wrong!")
		return
	}
	w.Header().Set("Content-Type", "video/mp4")
	http.ServeContent(w, r, "", time.Now(), file)
}

func uploadHandle(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	r.Body = http.MaxBytesReader(w, r.Body, MAX_VIDEO)

	err := r.ParseMultipartForm(MAX_VIDEO)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, "video is too big to upload")
		return
	}

	file, _, err := r.FormFile("file")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, "Get video failed")
		return
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, "read video to memory failed")
		return
	}

	vname := params.ByName("video_name")
	err = ioutil.WriteFile(VIDEO_DIR+vname, data, 0664)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, "write video to dick failed")
		return
	}
	w.WriteHeader(http.StatusCreated)
	io.WriteString(w, "uploaded successfully")

}

func RegisterHandler() *httprouter.Router {
	r := httprouter.New()
	r.GET("/hellog", helloHandle)
	r.POST("/hellop", helloHandle)
	r.GET("/videos/:video_name", streamHandle)
	r.POST("/upload/:video_name", uploadHandle)
	return r
}

// curl -XGET localhost:9000/hellog
// curl -XPOST localhost:9000/hellop -d "name=luo&age=11"
// curl -XPOST localhost:9000/hellop  -H "Content-Type: application/json" -d '{"name":"luo","age":111}'
func main() {
	r := RegisterHandler()
	http.ListenAndServe(":9000", r)
}
