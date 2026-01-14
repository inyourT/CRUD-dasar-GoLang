package impl

import (
	"context"
	"crud-dasar/helper"
	"crud-dasar/model/entity"
	"crud-dasar/repository"
	"database/sql"
	"errors"
)

type MahasiswaRepositoryImpl struct {
}

func NewMahasiswaRepository() repository.MahasiswaRepository {
	return &MahasiswaRepositoryImpl{}
}

// FindAll implements [repository.MahasiswaRepository].
func (m *MahasiswaRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []entity.Mahasiswa {
	SQL := "select id, nama, nim, jenis_kelamin, alamat from mahasiswa"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)

	defer rows.Close()

	var mahasiswas []entity.Mahasiswa
	for rows.Next() {
		mahasiswa := entity.Mahasiswa{}
		err := rows.Scan(&mahasiswa.ID, &mahasiswa.Nim, &mahasiswa.Nama, &mahasiswa.Alamat, &mahasiswa.JenisKelamin)
		helper.PanicIfError(err)

		mahasiswas = append(mahasiswas, mahasiswa)
	}
	return mahasiswas
}

// FindByID implements [repository.MahasiswaRepository].
func (m *MahasiswaRepositoryImpl) FindByID(ctx context.Context, tx *sql.Tx, id int) (entity.Mahasiswa, error) {
	SQL := "select id, nama, nim, jenis_kelamin, alamat from mahasiswa where id = ?"
	rows, err := tx.QueryContext(ctx, SQL, id)
	helper.PanicIfError(err)
	defer rows.Close()

	mahasiswa := entity.Mahasiswa{}

	if rows.Next() {
		err := rows.Scan(&mahasiswa.ID, &mahasiswa.Nim, &mahasiswa.Nama, &mahasiswa.Alamat, &mahasiswa.JenisKelamin)
		helper.PanicIfError(err)
		return mahasiswa, nil
	} else {
		return mahasiswa, errors.New("mahasiswa not found")
	}
}

// Save implements [repository.MahasiswaRepository].
func (m *MahasiswaRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, mahasiswa entity.Mahasiswa) entity.Mahasiswa {
	SQL := "insert into mahasiswa(nama, nim, jenis_kelamin, alamat) values (?, ?, ?, ?)"
	result, err := tx.ExecContext(ctx, SQL,
		mahasiswa.Nama,
		mahasiswa.Nim,
		mahasiswa.JenisKelamin,
		mahasiswa.Alamat,
	)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	mahasiswa.ID = int(id)
	return mahasiswa
}

// Update implements [repository.MahasiswaRepository].
func (m *MahasiswaRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, mahasiswa entity.Mahasiswa) entity.Mahasiswa {
	SQL := "update mahasiswa set nama=?, nim=?, jenis_kelamin=?, alamat=? where id=?"
	_, err := tx.ExecContext(ctx, SQL, mahasiswa)
	helper.PanicIfError(err)
	return mahasiswa
}

// Delete implements [repository.MahasiswaRepository].
func (m *MahasiswaRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, mahasiswa entity.Mahasiswa) {
	SQL := "delete from mahasiswa where id = ?"
	_, err := tx.ExecContext(ctx, SQL, mahasiswa.ID)

	helper.PanicIfError(err)

}
