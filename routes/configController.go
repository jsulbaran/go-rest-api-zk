package routes

import (
	"RestService/config"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"net/http"
)

func ConfigRoutes(configuration config.Config) *chi.Mux {
	router := chi.NewRouter()
	router.Get("/", getConfig)
	return router
}

func getConfig(writer http.ResponseWriter, request *http.Request) {
	response := make(map[string]string)
	response["status"] = "OK"
	render.JSON(writer, request, response)
}
