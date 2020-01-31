package main

import (
	"html/template"
	"io"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
)

func main() {

	r := httprouter.New()

	//curl -XGET -vvv http://127.0.0.1:9999/posts/my-second-post
	//curl -XGET -vvv http://127.0.0.1:9999/posts/wrong_path
	r.GET("/posts/:title", func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		title := params.ByName("title")
		_, err := os.Stat("static/posts/" + title)
		if err != nil {
			if os.IsNotExist(err) {
				w.WriteHeader(http.StatusNotFound)
				io.WriteString(w, "no path")
				return
			}
		}

		t := template.Must(template.ParseFiles("./static/posts/" + title + "/index.html"))
		t.Execute(w, nil)
	})

	//if you take auto create static site by Hugo.
	//you should set baseURL = "http://127.0.0.1:9999/static/" in 'config.toml' configuration file before you take 'hugo -D' command line for auto creation.
	r.ServeFiles("/static/*filepath", http.Dir("./static"))
	http.ListenAndServe(":9999", r)
}
