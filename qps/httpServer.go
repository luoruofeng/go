package qps

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

// wrk 测试结果
// wrk -c 10 -d 10s http://localhost:9000/hello

//   2 threads and 10 connections
//   Thread Stats   Avg      Stdev     Max   +/- Stdev
//     Latency     4.89ms   15.68ms 208.76ms   93.54%
//     Req/Sec    10.16k     6.11k   24.37k    62.37%

//   201100 requests in 10.10s, 24.55MB read
//   Requests/sec:  19906.57
//   Transfer/sec:      2.43MB

func QpsApp() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {

		go func() {
			fmt.Println("127.0.0.1 - -  [" + time.Now().Format("2006-01-02 15:04:05") + "]  \"GET /hello HTTP/1.1\" 200")
		}()

		io.WriteString(w, "hello world")
	})
	http.ListenAndServe(":9000", nil)
}
