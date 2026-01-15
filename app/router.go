package app

import (
	"crud-dasar/controller"
	"crud-dasar/exeption"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(mahasiswaController controller.MahasiswaController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/api/mahasiswa", mahasiswaController.FindAll)
	router.GET("/api/mahasiswa/:ID", mahasiswaController.FindByID)
	router.POST("/api/mahasiswas", mahasiswaController.Craete)
	router.PUT("/api/mahasiswas/:ID", mahasiswaController.Update)
	router.DELETE("/api/mahasiswas/:ID", mahasiswaController.Delete)

	router.PanicHandler = exeption.ErrorHandler

	return router
}
