package helper

import (
	"crud-dasar/model/entity"
	response "crud-dasar/model/entity/web/mahasiswa"
)

func ToMahasiswaResponse(mahasiswa entity.Mahasiswa) response.MahasiswaResponse {
	return response.MahasiswaResponse{
		ID:           mahasiswa.ID,
		Nama:         mahasiswa.Nama,
		Nim:          mahasiswa.Nim,
		JenisKelamin: mahasiswa.JenisKelamin,
		Alamat:       mahasiswa.Alamat,
	}
}

func ToMahasiswaResponses(mahasiswas []entity.Mahasiswa) []response.MahasiswaResponse {
	var mahasiswaResponse []response.MahasiswaResponse
	for _, mahasiswa := range mahasiswas {
		mahasiswaResponse = append(mahasiswaResponse, ToMahasiswaResponse(mahasiswa))
	}

	return mahasiswaResponse
}
