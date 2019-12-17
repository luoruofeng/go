from http.server import HTTPServer,SimpleHTTPRequestHandler
from http import HTTPStatus
import socketserver

# wrk 测试结果
# wrk -c 10 -d 10s http://localhost:8000/hello
#   2 threads and 10 connections

#   Thread Stats   Avg      Stdev     Max   +/- Stdev
#     Latency     2.08ms    2.03ms  18.95ms   85.90%
#     Req/Sec   255.34    194.74     1.60k    83.16%

#   5042 requests in 10.03s, 635.30KB read

#   Socket errors: connect 0, read 5093, write 13, timeout 0
# Requests/sec:    502.72
# Transfer/sec:     63.34KB

class MyHTTPRequestHandler(SimpleHTTPRequestHandler):
    def do_GET(self):
        path = str(self.path)
        if(path == "/hello"):
            self.send_response(HTTPStatus.OK)
            self.send_header("Content-type","text/html")
            self.end_headers()
            self.wfile.write("hello world".encode())

def run(server_class=HTTPServer, handler_class=MyHTTPRequestHandler):
    server_address = ('', 8000)
    httpd = server_class(server_address, handler_class)
    httpd.serve_forever()

if __name__ == "__main__":
    run()