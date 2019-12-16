package middleware

import (
	"fmt"
	"net/http"
	"time"
)

func TimeMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		now := time.Now()

		next.ServeHTTP(writer, request)

		processTime := time.Since(now)
		fmt.Println("process time %v", processTime)
	})

}

func hello(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("hello world"))
}

//Handle方法第二个参数为Handler接口，HandlerFunc方法实现了Handler接口
//HandlerFunc(f)如果f为一个有适当签名的方法，将会调用f
//所谓适当签名就是func(http.ResponseWriter,*http.Request)
func SimpleApp() {
	http.Handle("/hello", TimeMiddleware(http.HandlerFunc(hello)))
	http.ListenAndServe(":8080", nil)
}
