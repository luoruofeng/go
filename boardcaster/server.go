package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

type client chan<- string

var (
	entering = make(chan client)
	leaving  = make(chan client)
	message  = make(chan string)
)

func boardcaster() {
	clients := make(map[client]bool)
	for {
		select {
		case enterClient := <-entering:
			clients[enterClient] = true
		case leaveClient := <-leaving:
			clients[leaveClient] = false
			delete(clients, leaveClient)
		case mes := <-message:
			for cli := range clients {
				cli <- mes
			}
		}
	}
}

func handleConnection(conn net.Conn) {
	ch := make(chan string)
	go clientWriter(conn, ch)
	who := conn.RemoteAddr().String()

	entering <- ch
	ch <- "welcome " + who
	message <- who + " arriving"

	input := bufio.NewScanner(conn)
	for input.Scan() {
		message <- who + ":" + input.Text()
	}

	leaving <- ch
	message <- who + " left"

	conn.Close()
}

func clientWriter(conn net.Conn, ch chan string) {
	for message := range ch {
		fmt.Fprintln(conn, message)
	}
}

func main() {

	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	go boardcaster()

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Fatal(err)
			continue
		}
		go handleConnection(conn)
	}
}
