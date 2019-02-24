package main

import (
	"RestService/config"
	"RestService/routes"
	"RestService/service"
	"github.com/go-chi/chi"
	"github.com/tkanos/gonfig"
	"log"
	"net/http"
)

func main() {

	//configuration := config.NewConfig()
	var configuration config.Config

	err := gonfig.GetConf("configrest.json", &configuration)
	if err != nil {
		panic(err)
	}

	db, error2 := config.ConnectDatabase(configuration)
	configuration.DeviceSerial = service.GetDeviceSerial(configuration.SystemDatabasePath)
	if error2 != nil {
		panic(error2)
	}
	router := routes.Routes(db, configuration)
	walkFunc := func(method string, route string, handler http.Handler, middlewares ...func(handler2 http.Handler) http.Handler) error {
		log.Printf("%s %s\n", method, route)
		return nil
	}
	if err := chi.Walk(router, walkFunc); err != nil {
		log.Panicf("Loggin err :%s\n", err.Error())
	}
	log.Fatal(http.ListenAndServe(":"+configuration.Port, router))
}
