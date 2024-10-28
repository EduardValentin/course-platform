package main

import (
	"log"
	"net/http"

	"github.com/EduardValentin/course-platform/feature/auth/authenticator"
	"github.com/EduardValentin/course-platform/renderer"
	"github.com/EduardValentin/course-platform/router"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Failed to load the env vars: %v", err)
	}

	authenticator, err := authenticator.New()
	if err != nil {
		log.Fatalf("Failed to initialize the authenticator: %v", err)
	}

	router := router.Instance(authenticator)

	ginHtmlRenderer := router.HTMLRender
	router.HTMLRender = &renderer.HTMLTemplRenderer{FallbackHtmlRenderer: ginHtmlRenderer}

	router.SetTrustedProxies(nil)

	router.Static("/assets", "./assets")

	log.Println("Server listening on http://localhost:8080/")
	if err := http.ListenAndServe("0.0.0.0:8080", router); err != nil {
		log.Fatalf("There was an error with the http server: %v", err)
	}
}
