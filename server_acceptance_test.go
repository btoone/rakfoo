package main

import (
	"os/exec"
	"testing"
)

/*
Using curl command for an acceptance test.
*/
func TestCurlRoot(t *testing.T) {
	// curl -i -X GET http://localhost:8080/
	curl := exec.Command("curl", "-X", "GET", "http://localhost:8080/")
	out, err := curl.Output()

	if err != nil {
		t.Fatal(err)
	}

	actual := string(out)
	exp := "Hello, Rakuten!\n"

	if exp != actual {
		t.Error(err)
		t.Errorf("Expected %q, got %q instead.\n", exp, actual)
	}
}
