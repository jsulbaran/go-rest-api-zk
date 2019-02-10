package main

import (
	"RestService/routes"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"log"
	"net/http"
)

func Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Use(render.SetContentType(render.ContentTypeJSON), middleware.Logger, middleware.DefaultCompress, middleware.RedirectSlashes, middleware.Recoverer)
	router.Route("/api/v1", func(r chi.Router) {
		r.Mount("/users", routes.UserRoutes())          // gestión de usuarios
		r.Mount("/config", routes.ConfigRoutes())       // configuración
		r.Mount("/printer", routes.PrinterRoutes())     // eventos
		r.Mount("/eventlogs", routes.EventlogsRoutes()) // impresora
	})
	return router
}

func main() {
	router := Routes()
	walkFunc := func(method string, route string, handler http.Handler, middlewares ...func(handler2 http.Handler) http.Handler) error {
		log.Printf("%s %s\n", method, route)
		return nil
	}
	if err := chi.Walk(router, walkFunc); err != nil {
		log.Panicf("Loggin err :%s\n", err.Error())
	}
	log.Fatal(http.ListenAndServe(":8000", router))

}
