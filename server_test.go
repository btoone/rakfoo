package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

// GET /foo
// GET /foo/{id}
func TestGETFoo(t *testing.T) {
	list := StubFooList{
		map[string]string{
			"123": "Jack",
		},
	}

	server := &FooServer{&list}

	t.Run("returns a list of Foo items", func(t *testing.T) {
		t.Skip("To be implemented")
	})

	t.Run("returns Foo when an ID is specified", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/foo/123", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		actual := response.Body.String()
		expected := "{\"id\":\"123\",\"name\":\"Jack\"}\n"

		assertResponseBody(t, actual, expected)
	})

}

// GET /foo/000
func TestNotFound(t *testing.T) {
	list := StubFooListWithError{}

	server := &FooServer{&list}

	t.Run("returns 404 when Foo does not exist", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/foo/000", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		actual := response.Code
		expected := http.StatusNotFound

		assertStatus(t, actual, expected)
		assertResponseBody(t, response.Body.String(), "")
	})
}

// GET /foo
func TestGETRoot(t *testing.T) {
	server := &FooServer{}

	t.Run("returns a greeting", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		actual := response.Body.String()
		expected := "Hello, Rakuten!\n"

		assertResponseBody(t, actual, expected)
	})
}

// POST /foo -d '{"name": "Mike"}'
func TestCreateFoo(t *testing.T) {
	list := StubFooList{make(map[string]string)}
	server := &FooServer{&list}

	t.Run("responds with 200", func(t *testing.T) {
		body, _ := json.Marshal(Foo{"123", "Mike"})
		request, _ := http.NewRequest(http.MethodPost, "/foo", bytes.NewBuffer(body))
		request.Header.Set("Content-Type", "application/json; charset=UTF-8")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		actual := response.Code
		expected := http.StatusOK

		assertStatus(t, actual, expected)
	})
}

// POST /foo -d '{"name": "Mike"}'
func TestSaveFoo(t *testing.T) {
	list := StubFooList{make(map[string]string)}
	server := &FooServer{&list}

	t.Run("saves Foo to the list", func(t *testing.T) {
		foo := Foo{"123", "Mike"}
		body, _ := json.Marshal(foo)

		request, _ := http.NewRequest(http.MethodPost, "/foo", bytes.NewBuffer(body))
		request.Header.Set("Content-Type", "application/json; charset=UTF-8")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		actual := len(list.items)
		expected := 1

		if actual != expected {
			t.Errorf("Failure/Error: expected list size to be %d, got %d instead\n", expected, actual)
		}

	})
}

/*
Helpers
-----------------------
*/

func assertResponseBody(t testing.TB, actual, expected string) {
	t.Helper()

	if actual != expected {
		t.Errorf("Failure/Error: expected response to be %q, got %q instead\n", expected, actual)
	}
}

func assertStatus(t testing.TB, actual, expected int) {
	t.Helper()

	if actual != expected {
		t.Errorf("Failure/Error: expected status to = %d, got %d instead\n", expected, actual)
	}
}

/*
Stubs
-----------------------
*/

type StubFooList struct {
	items map[string]string
}

func (s *StubFooList) FindFoo(id string) (Foo, error) {
	foo := Foo{
		Id:   id,
		Name: s.items[id],
	}

	return foo, nil
}

func (s *StubFooList) SaveFoo(name string) Foo {
	foo := Foo{"000", "Stub"}

	// Persist Foo's attributes on our list
	s.items[foo.Id] = foo.Name

	return foo
}

/* Create a new list to stub with an error response */
type StubFooListWithError struct{}

func (s *StubFooListWithError) FindFoo(id string) (Foo, error) {
	return Foo{"000", "With Error"}, errors.New("Stub Error")
}

func (s *StubFooListWithError) SaveFoo(name string) Foo {
	return Foo{"000", "With Error"}
}
