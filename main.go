package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/EduardValentin/course-platform/ui"
	"github.com/EduardValentin/course-platform/utils"
	"github.com/a-h/templ"
)

var dev = true

func disableCacheInDevMode(next http.Handler) http.Handler {
	if !dev {
		return next
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "no-store")
		next.ServeHTTP(w, r)
	})
}

func withNonce(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		nonce := utils.GenerateSecureNonce(128)
		w.Header().Add("Content-Security-Policy", fmt.Sprintf("script-src 'nonce-%s'", nonce))

		ctx := templ.WithNonce(r.Context(), nonce)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/styles/", disableCacheInDevMode(http.StripPrefix("/styles/", http.FileServer(http.Dir("ui/styles")))))

	mux.Handle("/", templ.Handler(ui.Hello()))

	withNonceMux := withNonce(mux)

	fmt.Println("Listening on :3000")
	if err := http.ListenAndServe(":3000", withNonceMux); err != nil {
		log.Printf("error listening: %v", err)
	}
}
