package boardcaster_mongo

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func ClientApp() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		input := bufio.NewScanner(os.Stdin)
		for input.Scan() {
			content := input.Text()
			io.WriteString(conn, content+"\r\n")
		}
	}()

	output := bufio.NewScanner(conn)
	for output.Scan() {
		fmt.Println(output.Text())
	}

}
