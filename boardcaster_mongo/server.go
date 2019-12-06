package boardcaster_mongo

import (
	"bufio"
	"encoding/json"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"log"
	"math/rand"
	"net"
)

type message_chan chan Message

type User struct {
	Name string
	Id   primitive.ObjectID `bson:"_id"`
	Addr net.Addr
	Conn net.Conn
	Msgs message_chan
}

type Message struct {
	Content string
	User    *User
	Time    time.Time
}

var (
	entering        = make(chan *User)
	leaving         = make(chan *User)
	add_user_for_id = make(chan *User)
	message         = make(chan Message)
)

func dencode(obj interface{}) string {
	bytesobj, err := json.Marshal(obj)
	if err != nil {
		log.Fatal(err)
	}
	return string(bytesobj)
}

func encode(data []byte, obj interface{}) interface{} {
	err := json.Unmarshal(data, obj)
	if err != nil {
		log.Fatal(err)
	}
	return obj
}

var users map[*User]bool

func boardcaster() {
	MongoInit()
	users = make(map[*User]bool)
	for {
		select {
		case userP := <-entering:
			go func() {
				InsertUser(userP)
			}()
			users[userP] = true
		case userP := <-leaving:
			go func() {
				DeleteUser(userP)
			}()

			users[userP] = false
			delete(users, userP)
		case mes := <-message:
			go func() {
				InsertMessage(mes)
			}()
			for user := range users {
				user.Msgs <- mes
			}
		}
	}
}

func NewUserPointer(conn net.Conn) *User {
	rand.Seed(time.Now().Local().UnixNano())
	names := []string{
		"tony",
		"tom",
		"jack",
		"trump",
		"neo",
		"tide",
		"susan",
		"vetory",
		"rose",
		"jams",
		"kobe",
		"joden",
	}
	msgs := make(chan Message)
	username := names[rand.Intn(len(names))]
	objid := primitive.NewObjectID()
	return &User{Name: username, Addr: conn.RemoteAddr(), Conn: conn, Msgs: msgs, Id: objid}
}

func NewMessage(user *User, content string) Message {
	return Message{
		User:    user,
		Content: content,
		Time:    time.Now()}
}

func handleConnection(conn net.Conn) {

	userP := NewUserPointer(conn)
	go userP.clientWriter()

	entering <- userP

	userP.Msgs <- NewMessage(userP, "welcome "+userP.Name)
	message <- NewMessage(userP, userP.Name+" arriving")

	input := bufio.NewScanner(conn)
	for input.Scan() {
		message <- NewMessage(userP, userP.Name+" : "+input.Text())
	}

	leaving <- userP
	message <- NewMessage(userP, userP.Name+" left")

	conn.Close()
}

func (user *User) clientWriter() {
	for message := range user.Msgs {
		fmt.Fprintln(user.Conn, message.Content)
	}
}

func ServerApp() {
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	go boardcaster()
	go web()

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Fatal(err)
			continue
		}
		go handleConnection(conn)
	}
}
