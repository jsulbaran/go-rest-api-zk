package main

import (
	"RestService/config"
	"RestService/routes"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
)

func Routes(db *gorm.DB) *chi.Mux {
	router := chi.NewRouter()
	router.Use(render.SetContentType(render.ContentTypeJSON), middleware.Logger, middleware.DefaultCompress, middleware.RedirectSlashes, middleware.Recoverer)
	router.Route("/api/v1", func(r chi.Router) {
		r.Mount("/users", routes.UserRoutes(db))          // gestión de usuarios
		r.Mount("/config", routes.ConfigRoutes())         // configuración
		r.Mount("/printer", routes.PrinterRoutes(db))     // eventos
		r.Mount("/eventlogs", routes.EventlogsRoutes(db)) // impresora
	})
	return router
}

func main() {

	configuration := config.NewConfig()
	db, error2 := config.ConnectDatabase(configuration)
	if error2 != nil {
		panic(error2)
	}
	router := Routes(db)
	walkFunc := func(method string, route string, handler http.Handler, middlewares ...func(handler2 http.Handler) http.Handler) error {
		log.Printf("%s %s\n", method, route)
		return nil
	}
	if err := chi.Walk(router, walkFunc); err != nil {
		log.Panicf("Loggin err :%s\n", err.Error())
	}
	log.Fatal(http.ListenAndServe(":"+configuration.Port, router))
}
