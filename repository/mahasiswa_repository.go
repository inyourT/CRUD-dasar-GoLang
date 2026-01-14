package repository

import (
	"context"
	"crud-dasar/model/entity"
	"database/sql"
)

type MahasiswaRepository interface {
	Save(ctx context.Context, tx *sql.Tx, mahasiswa entity.Mahasiswa) entity.Mahasiswa
	Update(ctx context.Context, tx *sql.Tx, mahasiswa entity.Mahasiswa) entity.Mahasiswa
	Delete(ctx context.Context, tx *sql.Tx, mahasiswa entity.Mahasiswa)
	FindByID(ctx context.Context, tx *sql.Tx, id int) (entity.Mahasiswa, error)
	FindAll(ctx context.Context, tx *sql.Tx) []entity.Mahasiswa
}
