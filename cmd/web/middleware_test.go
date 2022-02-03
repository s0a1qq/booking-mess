package main

import (
	"fmt"
	"net/http"
	"testing"
)

func TestNoSurf(t *testing.T) {
	var myH myHandler

	h := NoSrufe(&myH)

	switch v := h.(type) {
	case http.Handler:

	default:
		t.Error(fmt.Sprintf("type is not http.handler but %T", v))
	}
}

func TestSessionLoad(t *testing.T) {
	var myH myHandler

	h := SessionLoad(&myH)

	switch v := h.(type) {
	case http.Handler:

	default:
		t.Error(fmt.Sprintf("type is not http.handler but %T", v))
	}
}
