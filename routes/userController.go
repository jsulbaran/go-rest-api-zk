package routes

import (
	"RestService/domain"
	"RestService/service"
	"RestService/util"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/jinzhu/gorm"
	"net/http"
)

var userService *service.UserService

func UserRoutes(orm *gorm.DB) *chi.Mux {
	router := chi.NewRouter()
	userService = service.NewUserService(orm)

	router.Route("/{id}", func(r chi.Router) {
		r.Get("/", getUser)
		r.Put("/", putUser)
		r.Delete("/", deleteUser)
	})
	router.Get("/", getUsers)
	router.Get("/pin/{id}", getUserByPin)
	router.Post("/", postUsers)

	//router.Get("/{id}", getUsers)
	//
	//router.Put("/", putUser)
	//router.Delete("/", deleteUser)
	return router
}
func getUser(writer http.ResponseWriter, request *http.Request) {
	id := chi.URLParam(request, "id")
	if !util.IsNumeric(id) {
		render.Render(writer, request, ErrInvalidRequest("No se indica un id valido"))
		return
	}
	idNum := util.StringToInt(id)
	user := userService.GetUserById(idNum)
	if user.InternalId == 0 {
		render.Render(writer, request, ErrNotFound("No existe el usuario"))
		return
	}
	render.JSON(writer, request, user)

}

func getUserByPin(writer http.ResponseWriter, request *http.Request) {
	id := chi.URLParam(request, "id")
	if !util.IsNumeric(id) {
		render.Render(writer, request, ErrInvalidRequest("No se indica un id valido"))
		return
	}
	user := userService.GetUserByPin(id)
	if user.InternalId == 0 {
		render.Render(writer, request, ErrNotFound("No existe el usuario"))
		return
	}
	render.JSON(writer, request, user)

}

func getUsers(writer http.ResponseWriter, request *http.Request) {
	var users = userService.GetUsersWithTemplates()
	if len(users) < 1 {
		render.Render(writer, request, ErrNotFound("No se hallaron usuarios"))
		return
	}
	render.JSON(writer, request, users)
}

func postUsers(writer http.ResponseWriter, request *http.Request) {
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
func putUser(writer http.ResponseWriter, request *http.Request) {

}
func deleteUser(writer http.ResponseWriter, request *http.Request) {

}

func getTemplates(writer http.ResponseWriter, request *http.Request) {
	db, _ := gorm.Open("sqlite3", "/home/joe/go/src/data/ZKDB.db")
	defer db.Close()
	var templates []domain.Template
	db.Find(&templates)
	render.JSON(writer, request, templates)
}
