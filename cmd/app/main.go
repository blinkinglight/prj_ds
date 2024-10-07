package main

import (
	"log"
	"net/http"

	"github.com/blinkinglight/prj_ds/pkg/routes"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {

	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	routes.SetupTest(router)
	log.Printf("Startgin server on port 8999")
	http.ListenAndServe(":8999", router)
}
