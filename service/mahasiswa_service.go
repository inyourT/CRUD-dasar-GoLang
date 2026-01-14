package service

import (
	"context"
	"crud-dasar/model/entity/web/mahasiswa"
)

type MahasiswaService interface {
	Create(ctx context.Context, request mahasiswa.MahasiswaCreateRequest) mahasiswa.MahasiswaResponse
	Update(ctx context.Context, request mahasiswa.MahasiswaUpdateRequest) mahasiswa.MahasiswaResponse
	Delete(ctx context.Context, id int)
	FindByID(ctx context.Context, id int) mahasiswa.MahasiswaResponse
	FindAll(ctx context.Context) []mahasiswa.MahasiswaResponse
}
