package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type Foo struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

/* Define the interface for a FooList */
type FooList interface {
	FindFoo(id string) (Foo, error)
	SaveFoo(name string) Foo
	DeleteFoo(id string)
}

/*
	FooServer is a handler because it satisfies the ServerHTTP interface
*/
type FooServer struct {
	list FooList
}

/*
  Defines the ServeHTTP method on the FooServer receiver which satisfies the interface

	server := &FooServer{}      // new instance of an empty struct
	server.ServeHTTP(res, req)
*/
func (f *FooServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Routing
	switch r.Method {
	case http.MethodPost:
		f.createFoo(w, r)
	case http.MethodGet:
		f.showFoo(w, r)
	case http.MethodDelete:
		f.deleteFoo(w, r)
	}
}

func (f *FooServer) showFoo(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/foo/")

	if id == "/" {
		fmt.Fprintln(w, "Hello, Rakuten!")
		return
	}

	foo, err := f.list.FindFoo(id)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(foo)
}

func (f *FooServer) createFoo(w http.ResponseWriter, r *http.Request) {
	var foo Foo
	err := json.NewDecoder(r.Body).Decode(&foo)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		// http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	foo = f.list.SaveFoo(foo.Name)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(foo)
}

func (f *FooServer) deleteFoo(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/foo/")
	fmt.Printf("Debug: id = %q", id)

	f.list.DeleteFoo(id)

	// if err != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	return
	// }

	w.WriteHeader(http.StatusNoContent)
}
