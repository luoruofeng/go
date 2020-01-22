package main

import (
	"net/http"
	"testing"
)

func TestMain(m *testing.M) {
	main()
	m.Run()
}

func BenchmarkRoomGet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		resp, err := http.Get("localhost:8080/room/popo")
		defer resp.Body.Close()
		if err != nil {
			b.Errorf("connect failed!!")
		}
		b.Logf(resp.Status)
	}
}
