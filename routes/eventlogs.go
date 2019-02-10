package routes

import (
	"RestService/domain"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"net/http"
)

func EventlogsRoutes() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/", getEvents)
	return router
}

func getEvents(writer http.ResponseWriter, request *http.Request) {
	db, _ := gorm.Open("sqlite3", "/home/joe/go/src/data/ZKDB.db")
	defer db.Close()

	var events []domain.Event
	//var event domain.Event
	db.Find(&events)
	//db.First(&event, 19)

	//response := make(map[string]string)
	//response["status"] = "OK"
	//response["id"] = event.UserId
	render.JSON(writer, request, events)
}
