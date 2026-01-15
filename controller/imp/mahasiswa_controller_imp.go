package imp

import (
	"crud-dasar/controller"
	"crud-dasar/helper"
	"crud-dasar/model/entity/web"
	"crud-dasar/model/entity/web/mahasiswa"
	"crud-dasar/service"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type MahasiswaControllerImp struct {
	MahasiswaService service.MahasiswaService
}

func NewMahasiswaController(mahasiswaService service.MahasiswaService) controller.MahasiswaController {
	return &MahasiswaControllerImp{
		MahasiswaService: mahasiswaService,
	}
}

// Craete implements [controller.MahasiswaController].
func (m *MahasiswaControllerImp) Craete(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	mahasiswaCreateRequest := mahasiswa.MahasiswaCreateRequest{}
	helper.ReadFromRequestBody(r, &mahasiswaCreateRequest)

	mahasiswaResponse := m.MahasiswaService.Create(r.Context(), mahasiswaCreateRequest)
	webResponse := web.WebResponse{
		Code:   http.StatusCreated,
		Status: "CREATED",
		Data:   mahasiswaResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}

// Delete implements [controller.MahasiswaController].
func (m *MahasiswaControllerImp) Delete(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	mahasiswaID := p.ByName("mahasiswa")
	id, err := strconv.Atoi(mahasiswaID)
	helper.PanicIfError(err)

	m.MahasiswaService.Delete(r.Context(), id)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
	}

	helper.WriteToResponseBody(w, webResponse)
}

// FindAll implements [controller.MahasiswaController].
func (m *MahasiswaControllerImp) FindAll(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	mahasiswaRespone := m.MahasiswaService.FindAll(r.Context())
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   mahasiswaRespone,
	}

	helper.WriteToResponseBody(w, webResponse)
}

// FindByID implements [controller.MahasiswaController].
func (m *MahasiswaControllerImp) FindByID(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	mahasiswaID := p.ByName("mahasiswa")
	id, err := strconv.Atoi(mahasiswaID)
	helper.PanicIfError(err)

	mahasiswaRespone := m.MahasiswaService.FindByID(r.Context(), id)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   mahasiswaRespone,
	}

	helper.WriteToResponseBody(w, webResponse)
}

// Update implements [controller.MahasiswaController].
func (m *MahasiswaControllerImp) Update(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	mahasiswaUpdateRequest := mahasiswa.MahasiswaUpdateRequest{}
	helper.ReadFromRequestBody(r, &mahasiswaUpdateRequest)

	mahasiswaResponse := m.MahasiswaService.Update(r.Context(), mahasiswaUpdateRequest)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "UPDATED",
		Data:   mahasiswaResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}
