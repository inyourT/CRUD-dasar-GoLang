package mahasiswa

type MahasiswaCreateRequest struct {
	Nama         string `validate:"required, max=100, min=3" json:"name"`
	Nim          int    `validate:"required, max=100, min=3" json:"nim"`
	Alamat       string `validate:"required, max=100, min=3" json:"alamat"`
	JenisKelamin string `validate:"required, max=100, min=3" json:"jenis_kelamin"`
}
