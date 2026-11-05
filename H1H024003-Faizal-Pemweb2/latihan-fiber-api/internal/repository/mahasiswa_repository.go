
package repository

import (
	"latihan-fiber-api/internal/config"
	"latihan-fiber-api/internal/model"
)

// Filter untuk pencarian dan pagination
type FilterMahasiswa struct {
	Cari       string
	Angkatan   int
	Halaman    int
	PerHalaman int
}

// Mengambil seluruh mahasiswa dengan filter dan pagination
func Ambil(filter FilterMahasiswa) ([]model.Mahasiswa, int64, error) {
	var daftar []model.Mahasiswa
	var total int64

	kueri := config.DB.Model(&model.Mahasiswa{})

	if filter.Cari != "" {
		pola := "%" + filter.Cari + "%"
		kueri = kueri.Where(
			"nama LIKE ? OR nim LIKE ?",
			pola,
			pola,
		)
	}

	if filter.Angkatan > 0 {
		kueri = kueri.Where("angkatan = ?", filter.Angkatan)
	}

	err := kueri.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	lompat := (filter.Halaman - 1) * filter.PerHalaman

	err = kueri.
		Preload("ProgramStudi").
		Order("nama ASC").
		Limit(filter.PerHalaman).
		Offset(lompat).
		Find(&daftar).Error

	if err != nil {
		return nil, 0, err
	}

	return daftar, total, nil
}

// Mengambil mahasiswa berdasarkan ID
func AmbilPerId(id uint) (model.Mahasiswa, error) {
	var mahasiswa model.Mahasiswa

	err := config.DB.
		Preload("ProgramStudi").
		First(&mahasiswa, id).Error

	return mahasiswa, err
}

// Mencari mahasiswa berdasarkan NIM
func CariBerdasarkanNIM(nim string) (model.Mahasiswa, error) {
	var mahasiswa model.Mahasiswa

	err := config.DB.
		Preload("ProgramStudi").
		Where("nim = ?", nim).
		First(&mahasiswa).Error

	return mahasiswa, err
}

// Menyimpan mahasiswa baru
func Simpan(mahasiswa *model.Mahasiswa) error {
	return config.DB.Create(mahasiswa).Error
}

// Memperbarui data mahasiswa
func Perbarui(mahasiswa *model.Mahasiswa) error {
	return config.DB.Save(mahasiswa).Error
}

// Menghapus mahasiswa
func Hapus(mahasiswa *model.Mahasiswa) error {
	return config.DB.Delete(mahasiswa).Error
}

// Mengambil mahasiswa yang sudah dihapus (termasuk soft delete)
func AmbilTerhapus(id uint) (model.Mahasiswa, error) {
	var mahasiswa model.Mahasiswa

	err := config.DB.
		Unscoped().
		Preload("ProgramStudi").
		First(&mahasiswa, id).Error

	return mahasiswa, err
}

// Memulihkan mahasiswa yang sudah dihapus
func Pulihkan(mahasiswa *model.Mahasiswa) error {
	return config.DB.
		Model(mahasiswa).
		Unscoped().
		Update("deleted_at", nil).Error
}

// Mengambil seluruh mahasiswa yang sudah dihapus
func AmbilSemuaTerhapus() ([]model.Mahasiswa, error) {
	var daftar []model.Mahasiswa

	err := config.DB.
		Unscoped().
		Preload("ProgramStudi").
		Where("deleted_at IS NOT NULL").
		Find(&daftar).Error

	return daftar, err
}
