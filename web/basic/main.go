package main

import (
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

type Item struct {
	Title string
}

type IndexParam struct {
	Title string
	Time  time.Time
	Items []Item
}

const ALLFONT string = "asdfghjklqwertyuiopzxcvbnm1234567890"

func RandomName() string {
	var resultBuilder strings.Builder
	nameLen := rand.Intn(6) + 1
	for i := 0; i < nameLen; i++ {
		resultBuilder.WriteByte(ALLFONT[rand.Intn(len(ALLFONT))])
	}
	return resultBuilder.String()
}

func NewIndexParam() *IndexParam {

	ip := &IndexParam{
		Title: "我是标题",
		Time:  time.Now(),
	}
	items := make([]Item, 15)
	for i := 0; i < 5; i++ {
		items = append(items, Item{
			Title: RandomName(),
		})
	}
	ip.Items = items
	return ip
}

func main() {
	http.HandleFunc("/index", func(w http.ResponseWriter, r *http.Request) {

		funcMap := template.FuncMap{
			// The name "odd" is what the function will be called in the template text.
			"odd": func(i int) bool {
				return (i % 2) == 1
			},
		}

		// Custom functions need to be registered before parsing the templates
		// if you use Funcs and ParseFiles at the same time.You should set argument of New function the same with html file name.
		// if you use Funcs and Parse .You can set any name in argument of New function.
		t, err := template.New("index.html").Funcs(funcMap).ParseFiles("./static/index.html")
		if err != nil {
			fmt.Println(err)
		}

		t.Execute(w, NewIndexParam())
	})

	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		t := template.Must(template.ParseFiles("static/about/hello-world/index.html", "static/tmpl/foot.html"))
		t.Execute(w, map[string]interface{}{"Title": "我是标题", "Body": "我是身体", "IndexParam": NewIndexParam()})
	})

	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("static"))))

	err := http.ListenAndServe(":9999", nil)
	if err != nil {
		log.Fatal(err)
	}
}
