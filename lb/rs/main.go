package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	r := httprouter.New()
	r.GET("/hello", func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		io.WriteString(w, "hello, here is real-server")
	})
	http.ListenAndServe(":8080", r)
	fmt.Println("Done!")
}
