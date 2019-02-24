package routes

import (
	"RestService/config"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/jinzhu/gorm"
)

func Routes(db *gorm.DB, config config.Config) *chi.Mux {
	router := chi.NewRouter()
	router.Use(render.SetContentType(render.ContentTypeJSON), middleware.Logger, middleware.DefaultCompress, middleware.RedirectSlashes, middleware.Recoverer, middleware.AllowContentType("application/json"))
	router.Route("/api/v1", func(r chi.Router) {
		r.Mount("/users", UserRoutes(db))                  // gestión de usuarios
		r.Mount("/config", ConfigRoutes(config))           // configuración
		r.Mount("/printer", PrinterRoutes(db))             // eventos
		r.Mount("/eventlogs", EventlogsRoutes(db, config)) // impresora
	})
	return router
}
