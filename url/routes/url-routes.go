package routes

import (
	"github.com/asahasrabuddhe/go-api-base/router"
	"github.com/asahasrabuddhe/url-shortener/url/handlers"
	"net/http"
)

func CreateUrlRoutes() {
	urlApiRouter := router.ApiRouter.PathPrefix("/url").Subrouter()
	urlRouter := router.Router.PathPrefix("").Subrouter()

	urlApiRouter.HandleFunc("/shrink", handlers.ShrinkURL).Methods(http.MethodPost)
	urlRouter.HandleFunc("/{key}", handlers.ResolveURL).Methods(http.MethodGet)
}
