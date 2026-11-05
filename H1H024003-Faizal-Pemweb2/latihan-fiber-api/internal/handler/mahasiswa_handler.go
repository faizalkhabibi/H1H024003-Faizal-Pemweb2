
package handler

import (
	"errors"
	"strconv"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"

	"latihan-fiber-api/internal/model"
	"latihan-fiber-api/internal/repository"
)

// GET /api/mahasiswa
// Menampilkan daftar mahasiswa
func DaftarMahasiswa(c fiber.Ctx) error {
	halaman, _ := strconv.Atoi(c.Query("halaman", "1"))
	perHalaman, _ := strconv.Atoi(c.Query("per_halaman", "10"))
	angkatan, _ := strconv.Atoi(c.Query("angkatan", "0"))

	if halaman < 1 {
		halaman = 1
	}

	if perHalaman < 1 {
		perHalaman = 10
	}

	if perHalaman > 100 {
		perHalaman = 100
	}

	filter := repository.FilterMahasiswa{
		Cari:       c.Query("cari"),
		Angkatan:   angkatan,
		Halaman:    halaman,
		PerHalaman: perHalaman,
	}

	daftar, total, err := repository.Ambil(filter)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"sukses": false,
			"pesan":  "Gagal mengambil data mahasiswa",
		})
	}

	return c.JSON(fiber.Map{
		"sukses": true,
		"data":   daftar,
		"pagination": fiber.Map{
			"halaman":      halaman,
			"per_halaman":  perHalaman,
			"total_data":   total,
			"total_halaman": (total + int64(perHalaman) - 1) / int64(perHalaman),
		},
	})
}

// GET /api/mahasiswa/nim/:nim
// Mencari mahasiswa berdasarkan NIM
func CariNIM(c fiber.Ctx) error {
	nim := c.Params("nim")

	mahasiswa, err := repository.CariBerdasarkanNIM(nim)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"sukses": false,
				"pesan":  "Mahasiswa tidak ditemukan",
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"sukses": false,
			"pesan":  "Gagal mencari mahasiswa",
		})
	}

	return c.JSON(fiber.Map{
		"sukses": true,
		"data":   mahasiswa,
	})
}

// GET /api/mahasiswa/:id
// Menampilkan detail mahasiswa
func DetailMahasiswa(c fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"sukses": false,
			"pesan":  "ID mahasiswa tidak valid",
		})
	}

	mahasiswa, err := repository.AmbilPerId(uint(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"sukses": false,
				"pesan":  "Mahasiswa tidak ditemukan",
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"sukses": false,
			"pesan":  "Gagal mengambil detail mahasiswa",
		})
	}

	return c.JSON(fiber.Map{
		"sukses": true,
		"data":   mahasiswa,
	})
}

// POST /api/mahasiswa
// Menambahkan mahasiswa baru
func BuatMahasiswa(c fiber.Ctx) error {
	var mahasiswa model.Mahasiswa

	if err := c.Bind().Body(&mahasiswa); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"sukses": false,
			"pesan":  "Format JSON tidak valid",
		})
	}

	if mahasiswa.Nim == "" || mahasiswa.Nama == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"sukses": false,
			"pesan":  "NIM dan nama wajib diisi",
		})
	}

	if err := repository.Simpan(&mahasiswa); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"sukses": false,
			"pesan":  "Gagal menyimpan mahasiswa",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"sukses": true,
		"pesan":  "Mahasiswa berhasil ditambahkan",
		"data":   mahasiswa,
	})
}

// PUT /api/mahasiswa/:id
// Memperbarui mahasiswa
func PerbaruiMahasiswa(c fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"sukses": false,
			"pesan":  "ID mahasiswa tidak valid",
		})
	}

	mahasiswa, err := repository.AmbilPerId(uint(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"sukses": false,
				"pesan":  "Mahasiswa tidak ditemukan",
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"sukses": false,
			"pesan":  "Gagal mengambil data mahasiswa",
		})
	}

	if err := c.Bind().Body(&mahasiswa); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"sukses": false,
			"pesan":  "Format JSON tidak valid",
		})
	}

	mahasiswa.ID = uint(id)

	if err := repository.Perbarui(&mahasiswa); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"sukses": false,
			"pesan":  "Gagal memperbarui mahasiswa",
		})
	}

	return c.JSON(fiber.Map{
		"sukses": true,
		"pesan":  "Mahasiswa berhasil diperbarui",
		"data":   mahasiswa,
	})
}

// DELETE /api/mahasiswa/:id
// Menghapus mahasiswa
func HapusMahasiswa(c fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"sukses": false,
			"pesan":  "ID mahasiswa tidak valid",
		})
	}

	mahasiswa, err := repository.AmbilPerId(uint(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"sukses": false,
				"pesan":  "Mahasiswa tidak ditemukan",
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"sukses": false,
			"pesan":  "Gagal mengambil data mahasiswa",
		})
	}

	if err := repository.Hapus(&mahasiswa); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"sukses": false,
			"pesan":  "Gagal menghapus mahasiswa",
		})
	}

	return c.JSON(fiber.Map{
		"sukses": true,
		"pesan":  "Mahasiswa berhasil dihapus",
	})
}

// POST /api/mahasiswa/:id/pulihkan
// Memulihkan mahasiswa yang sudah dihapus
func PulihkanMahasiswa(c fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))

	if err != nil || id < 1 {
		return fiber.NewError(
			fiber.StatusBadRequest,
			"ID tidak valid",
		)
	}

	mahasiswa, err := repository.AmbilTerhapus(uint(id))

	if err != nil {
		return fiber.NewError(
			fiber.StatusNotFound,
			"Data mahasiswa tidak ditemukan",
		)
	}

	if !mahasiswa.DeletedAt.Valid {
		return fiber.NewError(
			fiber.StatusBadRequest,
			"Data mahasiswa belum dihapus",
		)
	}

	err = repository.Pulihkan(&mahasiswa)

	if err != nil {
		return fiber.NewError(
			fiber.StatusInternalServerError,
			"Gagal memulihkan data",
		)
	}

	return c.JSON(fiber.Map{
		"sukses": true,
		"pesan":  "Data mahasiswa berhasil dipulihkan",
	})
}

// GET /api/mahasiswa/terhapus
// Menampilkan data mahasiswa yang sudah dihapus
func DaftarMahasiswaTerhapus(c fiber.Ctx) error {
	daftar, err := repository.AmbilSemuaTerhapus()

	if err != nil {
		return fiber.NewError(
			fiber.StatusInternalServerError,
			"Gagal mengambil data terhapus",
		)
	}

	return c.JSON(fiber.Map{
		"sukses": true,
		"data":   daftar,
	})
}
