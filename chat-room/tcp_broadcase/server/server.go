package main

import (
	"bufio"
	"io"
	"log"
	"net"
)

type Client chan string

var (
	entering = make(chan Client)
	leaving  = make(chan Client)
	messages = make(chan string)
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
		case m := <-messages:
			for c := range clients {
				c <- m + "\n"
			}
		}
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	name := conn.RemoteAddr().String()
	c := make(Client)
	go clientWriter(conn, c)
	c <- "hello " + name

	entering <- c
	messages <- name + " is coming"

	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- name + " : " + input.Text()
	}

	leaving <- c
	messages <- "bye " + name

}

func clientWriter(conn net.Conn, ch <-chan string) {
	for m := range ch {
		io.WriteString(conn, m)
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
