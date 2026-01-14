package mahasiswa

type MahasiswaResponse struct {
	ID           int    `json:"id"`
	Nama         string `json:"name"`
	Nim          int    `json:"nim"`
	Alamat       string `json:"alamat"`
	JenisKelamin string `json:"jenis_kelamin"`
}
