package routes

import (
	"RestService/domain"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"net/http"
)

func UserRoutes() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/", getUsers)
	return router
}

func getUsers(writer http.ResponseWriter, request *http.Request) {
	users := []domain.User{{InternalId: 1,
		Name:   "user1",
		UserId: "24773769"}, {UserId: "24780326", Name: "user2", InternalId: 2}}
	render.JSON(writer, request, users)
}
