package impl

import (
	"context"
	"crud-dasar/exeption"
	"crud-dasar/helper"
	"crud-dasar/model/entity"
	"crud-dasar/model/entity/web/mahasiswa"
	"crud-dasar/repository"
	"crud-dasar/service"
	"database/sql"

	"github.com/go-playground/validator/v10"
)

type MahasiswaServiceImpl struct {
	MahasiswaRepository repository.MahasiswaRepository
	DB                  *sql.DB
	validate            *validator.Validate
}

func NewMahasiswaService(
	mahasiswaRepository repository.MahasiswaRepository,
	DB *sql.DB,
	validate *validator.Validate) service.MahasiswaService {
	return &MahasiswaServiceImpl{
		MahasiswaRepository: mahasiswaRepository,
		DB:                  DB,
		validate:            validate,
	}
}

// Create implements [service.MahasiswaService].
func (m *MahasiswaServiceImpl) Create(ctx context.Context, request mahasiswa.MahasiswaCreateRequest) mahasiswa.MahasiswaResponse {
	err := m.validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := m.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	mahasiswa := entity.Mahasiswa{
		Nama:         request.Nama,
		Nim:          request.Nim,
		JenisKelamin: request.JenisKelamin,
		Alamat:       request.Alamat,
	}

	mahasiswa = m.MahasiswaRepository.Save(ctx, tx, mahasiswa)
	return helper.ToMahasiswaResponse(mahasiswa)
}

// Delete implements [service.MahasiswaService].
func (m *MahasiswaServiceImpl) Delete(ctx context.Context, id int) {
	tx, err := m.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	mahasiswa, err := m.MahasiswaRepository.FindByID(ctx, tx, id)
	if err != nil {
		panic(exeption.NewNotFoundError(err.Error()))
	}

	m.MahasiswaRepository.Delete(ctx, tx, mahasiswa)
}

// FindAll implements [service.MahasiswaService].
func (m *MahasiswaServiceImpl) FindAll(ctx context.Context) []mahasiswa.MahasiswaResponse {
	tx, err := m.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	mahasiswa := m.MahasiswaRepository.FindAll(ctx, tx)

	return helper.ToMahasiswaResponses(mahasiswa)
}

// FindID implements [service.MahasiswaService].
func (m *MahasiswaServiceImpl) FindByID(ctx context.Context, id int) mahasiswa.MahasiswaResponse {
	tx, err := m.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	mahasiswa, err := m.MahasiswaRepository.FindByID(ctx, tx, id)
	if err != nil {
		panic(exeption.NewNotFoundError(err.Error()))
	}

	return helper.ToMahasiswaResponse(mahasiswa)

}

// Update implements [service.MahasiswaService].
func (m *MahasiswaServiceImpl) Update(ctx context.Context, request mahasiswa.MahasiswaUpdateRequest) mahasiswa.MahasiswaResponse {
	err := m.validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := m.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	mahasiswa, err := m.MahasiswaRepository.FindByID(ctx, tx, request.ID)
	if err != nil {
		panic(exeption.NewNotFoundError(err.Error()))
	}

	mahasiswa.Alamat = request.Alamat
	mahasiswa.JenisKelamin = request.JenisKelamin
	mahasiswa.Nama = request.Nama
	mahasiswa.Nim = request.Nim

	mahasiswa = m.MahasiswaRepository.Update(ctx, tx, mahasiswa)
	return helper.ToMahasiswaResponse(mahasiswa)
}
