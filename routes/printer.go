package routes

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"net/http"
)

func PrinterRoutes() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/", getPrinter)
	return router
}

func getPrinter(writer http.ResponseWriter, request *http.Request) {
	response := make(map[string]string)
	response["status"] = "OK"
	render.JSON(writer, request, response)
}
