package internal

import (
	"fmt"
	"log"
	"net/http"
)

// Middleware for enforcing a query parameter
func EnforceQueryParam(param string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()

		params, ok := q[param]
		if !ok || len(params[0]) < 1 {
			log.Printf("Url Param '%s' is missing!", param)
			errorPage := defaultHandler("Parameter missing!", fmt.Sprintf("Url param '%s' is not included in the URL provided.", param), true)
			errorPage(w, r)
			return
		}

		next.ServeHTTP(w, r)
	})
}
