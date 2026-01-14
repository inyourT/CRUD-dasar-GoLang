package imp

import (
	"crud-dasar/controller"
	"crud-dasar/service"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type MahasiswaControllerImp struct {
	MahasiswaService service.MahasiswaService
}

func NewMahasiswaController(mahasiswaService service.MahasiswaService) controller.MahasiswaController {
	return &MahasiswaControllerImp{}
}

// Craete implements [controller.MahasiswaController].
func (m *MahasiswaControllerImp) Craete(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	panic("unimplemented")
}

// Delete implements [controller.MahasiswaController].
func (m *MahasiswaControllerImp) Delete(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	panic("unimplemented")
}

// FindAll implements [controller.MahasiswaController].
func (m *MahasiswaControllerImp) FindAll(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	panic("unimplemented")
}

// FindByID implements [controller.MahasiswaController].
func (m *MahasiswaControllerImp) FindByID(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	panic("unimplemented")
}

// Update implements [controller.MahasiswaController].
func (m *MahasiswaControllerImp) Update(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	panic("unimplemented")
}

