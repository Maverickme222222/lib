// Package web contains help code
package web

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// CORS adding CORS
func CORS(r *mux.Router) http.Handler {
	corsOpts := cors.New(cors.Options{
		AllowedOrigins: []string{"*"}, //you service is available and allowed for this base url
		AllowedMethods: []string{
			http.MethodGet, //http methods for your app
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
			http.MethodOptions,
			http.MethodHead,
		},

		AllowedHeaders: []string{
			"Origin",
			"Content-Type",
			"Accept",
			"Authorization",
			"X-Entity-Token",
		},
		AllowCredentials: true,
	})

	return corsOpts.Handler(r)
}
