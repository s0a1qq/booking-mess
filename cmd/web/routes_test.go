package main

import (
	"fmt"
	"testing"

	"github.com/go-chi/chi"
	"github.com/s0a1qq/booking-mess/internal/config"
)

func TestRoutes(t *testing.T) {
	var app config.AppConfig

	mux := routes(&app)

	switch v := mux.(type) {
	case *chi.Mux:

	default:
		t.Error(fmt.Sprintf("type is not chi.Mux but %T", v))
	}
}
