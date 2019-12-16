package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func MyTimeMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		now := time.Now()

		next.ServeHTTP(writer, request)

		processTime := time.Since(now)
		fmt.Println("process time %v", processTime)
	})

}

func helloworld(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("hello world"))
}

func ChiApp() {
	r := chi.NewRouter()
	r.Use(MyTimeMiddleware)  //自己实现的middleware
	r.Use(middleware.Logger) //chi的middleware模块自带方法
	r.Get("/hello", helloworld)

	http.ListenAndServe(":8080", r)
}
