package routes

import (
	"RestService/domain"
	"RestService/service"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/jinzhu/gorm"
	"net/http"
)

var userService *service.UserService

func UserRoutes(orm *gorm.DB) *chi.Mux {
	router := chi.NewRouter()
	userService = service.NewUserService(orm)
	router.Get("/", getUsers)
	return router
}

func getUsers(writer http.ResponseWriter, request *http.Request) {
	var users = userService.GetUsersWithTemplates()
	//var templates domain.Template
	//db.First(&users, 5).Model(&users).Related(&templates)
	//db.Preload("Biometric").Find(&users, 5)
	//db.First(&users, 5).Model(&users).Related(&templates)
	//query := db.Debug().Model(&users).Related(&templates).Error
	//query1 := db.Debug().Find(&users, 5)
	//query := db.Debug().Model(&users).Related(&templates).Error
	//if query1 != nil {
	//	panic(query1)
	//}
	//if query != nil {
	//	panic(query)
	//}
	render.JSON(writer, request, users)

}
func getTemplates(writer http.ResponseWriter, request *http.Request) {
	db, _ := gorm.Open("sqlite3", "/home/joe/go/src/data/ZKDB.db")
	defer db.Close()
	var templates []domain.Template
	db.Find(&templates)
	render.JSON(writer, request, templates)
}
