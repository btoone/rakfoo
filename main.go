package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
)

type InMemoryFooList struct {
	items map[string]string
}

func (i *InMemoryFooList) FindFoo(id string) (Foo, error) {
	item := i.items[id]

	if item == "" {
		return Foo{}, errors.New("No item found in list")
	} else {
		foo := Foo{
			Id:   id,
			Name: i.items[id],
		}
		return foo, nil
	}

}

func (i *InMemoryFooList) SaveFoo(name string) Foo {
	uuid := uuid.New().String()

	foo := Foo{
		Id:   uuid,
		Name: name,
	}

	// Persist Foo's attributes on our list
	i.items[uuid] = foo.Name

	// Return Foo struct
	return foo
}

// Run using: `go run .`
func main() {
	fmt.Println("Starting server...")

	server := &FooServer{&InMemoryFooList{make(map[string]string)}}
	err := http.ListenAndServe("localhost:8080", server)

	log.Fatal(err)
}
