package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"io"
	"io/ioutil"
	"net/http"
)

func helloGetHandle(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	fmt.Println("-----------Request Header---------")
	for k, v := range r.Header {
		for ind,_ := range v{
			fmt.Println(k + " : " + v[ind])
		}
	}

	fmt.Println("-----------Request Body---------")
	rb, _ := ioutil.ReadAll(r.Body)
	fmt.Println(string(rb))

	w.WriteHeader(http.StatusOK)
	io.WriteString(w, "Get Hello world")
}

func helloPostHandle(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	fmt.Println("-----------Request Header---------")
	for k, v := range r.Header {
		for ind,_ := range v{
			fmt.Println(k + " : " + v[ind])
		}
	}

	fmt.Println("-----------Request Body---------")
	rb, _ := ioutil.ReadAll(r.Body)
	fmt.Println(string(rb))
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, "Post Hello world")
}

func RegisterHandler() *httprouter.Router {
	r := httprouter.New()
	r.GET("/hellog", helloGetHandle)
	r.POST("/hellop", helloPostHandle)
	return r
}

// curl -XGET localhost:9000/hellog
// curl -XPOST localhost:9000/hellop -d "name=luo&age=11"
// curl -XPOST localhost:9000/hellop  -H "Content-Type: application/json" -d '{"name":"luo","age":111}'
func main() {
	r := RegisterHandler()
	http.ListenAndServe(":9000", r)
}
