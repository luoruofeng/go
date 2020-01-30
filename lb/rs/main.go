package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"strings"

	"github.com/julienschmidt/httprouter"
	"strconv"
)

var port int

func init() {
	flag.IntVar(&port, "port", 8080, "port of server")
}

func main() {
	flag.Parse()

	r := httprouter.New()

	r.HEAD("/hello",func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		w.WriteHeader(http.StatusOK)
		io.WriteString("这个head是为了避免阿里云SLB对ECS的监控监控4XX的错误，如果SLB配置时候勾选了4XX也不需要这个HEAD")
	})

	r.GET("/hello", func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		var builder strings.Builder
		builder.WriteString("hello, here is real-server\n")
		builder.WriteString("port:" + strconv.Itoa(port) + "\n")

		name, err := os.Hostname()
		if err != nil {
			fmt.Printf("Oops: %v\n", err)
			return
		}
		builder.WriteString("hostname:" + name + "\n")

		addrs, err := net.LookupHost(name)
		if err != nil {
			fmt.Printf("Oops: %v\n", err)
			return
		}

		for _, a := range addrs {
			builder.WriteString("IP:" + a + "\n")
		}

		io.WriteString(w, builder.String())
	})
	http.ListenAndServe(":"+strconv.Itoa(port), r)
	fmt.Println("Done!")
}
