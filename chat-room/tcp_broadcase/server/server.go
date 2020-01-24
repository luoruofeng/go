package main

import (
	"bufio"
	"encoding/json"
	"io"
	"log"
	"net"
	"time"
)

type Info struct {
	Name       string    `json:"name"`
	Message    string    `json:"message"`
	CreateTime time.Time `json:"create_time"`
}

type Client chan Info

var (
	entering = make(chan Client)
	leaving  = make(chan Client)
	messages = make(chan Info)
)

func broadcaster() {
	clients := make(map[Client]bool)
	for {
		select {
		case c := <-entering:
			clients[c] = true
		case c := <-leaving:
			close(c)
			delete(clients, c)
		case info := <-messages:
			for c := range clients {
				c <- info
			}
		}
	}
}

func NewInfo(name string, message string, createTime time.Time) Info {
	return Info{
		Name:       name,
		Message:    message,
		CreateTime: createTime,
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	name := conn.RemoteAddr().String()
	c := make(Client)
	go clientWriter(conn, c)
	c <- NewInfo(name, "hello "+name, time.Now())

	entering <- c
	messages <- NewInfo(name, name+" is coming", time.Now())

	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- NewInfo(name, name+" : "+input.Text(), time.Now())
	}

	leaving <- c
	messages <- NewInfo(name, "bye "+name, time.Now())

}

func clientWriter(conn net.Conn, ch <-chan Info) {
	for info := range ch {
		text, _ := json.Marshal(info)
		io.WriteString(conn, string(text)+"\n")
	}
}

func main() {
	go broadcaster()

	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Println("lisent 8080 failed")
		panic(1)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println("lisent 8080 failed")
			continue
		}
		go handleConnection(conn)
	}

}
