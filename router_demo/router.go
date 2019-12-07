package router_demo

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"

	"github.com/julienschmidt/httprouter"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

func Show(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	htmlPath := filepath.Join("./router_demo/tem/", ps.ByName("file_name")+".html")

	content, _ := ioutil.ReadFile(htmlPath)
	t, err := template.New("httpr").Parse(string(content))
	if err != nil {
		fmt.Fprintf(w, "404, %s!\n", ps.ByName("file_name"))
	}
	t.Execute(w, "i'm parameters")
}

func Web() {
	router := httprouter.New()

	//static server in net/http.
	//and just used to request of net/http. for example: http.HandleFunc()
	// http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./router_demo/static"))))

	//static server in httprouter
	router.ServeFiles("/static/*filepath", http.Dir("./router_demo/static"))

	router.GET("/hello/:file_name", Show) //http://localhost:8080/hello/httpr

	log.Fatal(http.ListenAndServe(":8080", router))
}
