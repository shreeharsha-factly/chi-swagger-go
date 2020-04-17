// Data Portal API.
//
// the purpose of this application is to provide an application
// that is using plain go code to define an API
//
// This should demonstrate all the possible comment annotations
// that are available to turn go code into a fully compliant swagger 2.0 spec
//
//     Schemes: http, https
//     Host: localhost:1323
//     BasePath: /
//     Version: 0.0.1
//
// swagger:meta
package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-openapi/runtime/middleware"
	"github.com/rs/cors"
	"github.com/shreeharsha-factly/chi-swagger-go/actions"
	"github.com/shreeharsha-factly/chi-swagger-go/models"
)

func main() {
	// db setup
	models.SetupDB()

	// db migrations
	models.DB.AutoMigrate(&models.User{})

	r := chi.NewRouter()

	r.Route("/users", func(r chi.Router) {
		r.Post("/", actions.CreateUser)
		r.Route("/{userId}", func(r chi.Router) {
			r.Get("/", actions.GetUser)
			r.Delete("/", actions.DeleteUser)
			r.Put("/", actions.UpdateUser)
		})
	})
	opts := middleware.RedocOpts{SpecURL: "/swagger.yaml"}
	sh := middleware.Redoc(opts, nil)

	r.Handle("/docs", sh)
	r.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))
	log.Fatal(http.ListenAndServe(":1323", setupGlobalMiddleware(r)))
}

func setupGlobalMiddleware(handler http.Handler) http.Handler {
	handleCORS := cors.Default().Handler

	return handleCORS(handler)
}
