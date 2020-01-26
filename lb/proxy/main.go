package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"github.com/luoruofeng/config"
)

type proxyHandler struct {
	lbAddr *url.URL
}

func (ph proxyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	reverseProxy := httputil.NewSingleHostReverseProxy(ph.lbAddr)
	reverseProxy.ServeHTTP(w, r)
}

func NewProxyHandler(lbAddr *url.URL) http.Handler {
	return proxyHandler{
		lbAddr: lbAddr,
	}
}

// curl -XGET http://localhost:8000/hello

func main() {
	url, _ := url.Parse(config.GetLbAddr())
	ph := NewProxyHandler(url)
	http.ListenAndServe(":8000", ph)
	fmt.Println("Done!")
}
