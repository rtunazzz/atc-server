package internal

import (
	"net/http"
	"path/filepath"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

// NewRouter returns a new HTTP handler that implements the main server routes
func NewRouter() http.Handler {
	router := chi.NewRouter()

	// Set up our middleware with sane defaults
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Timeout(60 * time.Second))

	// serve CSS directory
	cssPath, _ := filepath.Abs("./web/cdn/css")
	cssFs := http.FileServer(http.Dir(cssPath))
	router.Handle("/css/*", http.StripPrefix("/css/", cssFs))

	// serve scripts directory
	scriptsPath, _ := filepath.Abs("./web/scripts")
	scriptsFs := http.FileServer(http.Dir(scriptsPath))
	router.Handle("/scripts/*", http.StripPrefix("/scripts/", scriptsFs))

	// cookie routes
	router.HandleFunc("/snipes/set_locale", setLocaleHandler("snipes"))
	router.HandleFunc("/solebox/set_locale", setLocaleHandler("solebox"))

	router.HandleFunc("/nike/set_locale", nikeCookieHandler())

	// atc handling
	atcHandler := router.Middlewares().HandlerFunc(submitFormHandler("addtocart"))
	router.Handle("/atc", EnforceQueryParam("site", atcHandler))

	// wishlist handling
	wishlistHandler := router.Middlewares().HandlerFunc(submitFormHandler("wishlist"))
	router.Handle("/wishlist", EnforceQueryParam("site", wishlistHandler))
	router.Handle("/redirect", EnforceQueryParam("url", redirectHandler()))

	router.HandleFunc("/", defaultHandler("Welcome to BONZAY ATC links!", "", false))

	// Wildcard redirect for any other route than above
	router.HandleFunc("/*", defaultHandler("Not found.", "", true))

	return router
}
