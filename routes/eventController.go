package routes

import (
	"RestService/config"
	"RestService/service"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"net/http"
	"time"
)

var eventService *service.EventService

func EventlogsRoutes(orm *gorm.DB, configuration *config.Config) *chi.Mux {
	router := chi.NewRouter()
	eventService = service.NewEventService(orm, configuration)
	router.Get("/byDate", getEvents)
	return router
}

func getEvents(writer http.ResponseWriter, request *http.Request) {
	fromDate := request.URL.Query().Get("fromDate")
	if len(fromDate) < 1 {
		// bad reqquest no esta la fecha fromDate
		render.Status(request, http.StatusBadRequest)
		render.Render(writer, request, ErrInvalidRequest("No se indica parametro fromDate"))
		return
	}
	dateFrom, err := time.Parse("20060102150405", fromDate)
	if err != nil {
		render.Render(writer, request, ErrInvalidRequest("Parametro fromDate en formato invalido"))
		return
	}

	toDate := request.URL.Query().Get("toDate")
	var dateTo = time.Now()
	if len(toDate) > 0 {
		t2, err2 := time.Parse("20060102150405", toDate)
		if err != err2 {
			render.Render(writer, request, ErrInvalidRequest("Parametro toDate en formato invalido"))
			return
		}
		dateTo = t2
	}
	var events = eventService.GetEvents(dateFrom, dateTo)
	render.JSON(writer, request, events)
}
