package protobuf

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"go_demo/protobuf/tutorial"
	"io/ioutil"
	"log"
)

func ProtoApp() {
	p := tutorial.Person{
		Id:    1234,
		Name:  "John Doe",
		Email: "jdoe@example.com",
		Phones: []*tutorial.Person_PhoneNumber{
			{Number: "555-4321", Type: tutorial.Person_HOME},
		},
	}

	book := &tutorial.AddressBook{}
	book.People = append(book.People, &p)

	// Write the new address book back to disk.
	fname := "./protobuf/disk"
	out, err := proto.Marshal(book)
	if err != nil {
		log.Fatalln("Failed to encode address book:", err)
	}
	if err := ioutil.WriteFile(fname, out, 0644); err != nil {
		log.Fatalln("Failed to write address book:", err)
	}

	// Read the existing address book.
	in, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Fatalln("Error reading file:", err)
	}
	readBook := &tutorial.AddressBook{}
	if err := proto.Unmarshal(in, readBook); err != nil {
		log.Fatalln("Failed to parse address book:", err)
	}

	fmt.Println(readBook.People[0].GetEmail())
}
