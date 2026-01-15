package main

import (
	"crud-dasar/app"
	controller "crud-dasar/controller/imp"
	"crud-dasar/helper"
	repository "crud-dasar/repository/impl"
	service "crud-dasar/service/impl"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func main() {

	db := app.NewDB()
	validate := validator.New()

	mahasiswaRepository := repository.NewMahasiswaRepository()
	mahasiswaService := service.NewMahasiswaService(mahasiswaRepository, db, validate)
	mahasiswaController := controller.NewMahasiswaController(mahasiswaService)

	app.NewRouter(mahasiswaController)

	server := http.Server{
		Addr: "localhost:3000",
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
