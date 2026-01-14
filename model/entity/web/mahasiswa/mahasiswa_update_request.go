package mahasiswa

type MahasiswaUpdateRequest struct {
	ID           int    `validate:"required" json:"id"`
	Nama         string `validate:"max=100, min=3" json:"name"`
	Nim          int    `validate:"max=100, min=3" json:"nim"`
	Alamat       string `validate:"max=100, min=3" json:"alamat"`
	JenisKelamin string `validate:"max=100, min=3" json:"jenis_kelamin"`
}
