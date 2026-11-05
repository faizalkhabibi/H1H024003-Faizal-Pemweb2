package router

import (
	"github.com/gofiber/fiber/v3"
	"latihan-fiber-api/internal/handler"
)

func Daftarkan(app *fiber.App) {
	api := app.Group("/api")

	// Status API
	api.Get("/status", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"sukses": true,
			"pesan":  "API Fiber aktif",
		})
	})

	// Route program studi
	api.Get("/program-studi", handler.DaftarProgramStudi)

	// Route mahasiswa
	mahasiswa := api.Group("/mahasiswa")

	// Menampilkan daftar mahasiswa
	mahasiswa.Get("/", handler.DaftarMahasiswa)

	// Mencari mahasiswa berdasarkan NIM
	mahasiswa.Get("/nim/:nim", handler.CariNIM)

	// Menambahkan mahasiswa
	mahasiswa.Post("/", handler.BuatMahasiswa)

	// Daftar mahasiswa yang sudah dihapus (harus sebelum /:id)
	mahasiswa.Get("/terhapus", handler.DaftarMahasiswaTerhapus)

	// Detail mahasiswa berdasarkan ID
	mahasiswa.Get("/:id", handler.DetailMahasiswa)

	// Memperbarui mahasiswa
	mahasiswa.Put("/:id", handler.PerbaruiMahasiswa)

	// Menghapus mahasiswa
	mahasiswa.Delete("/:id", handler.HapusMahasiswa)

	// Memulihkan mahasiswa yang sudah dihapus
	mahasiswa.Post("/:id/pulihkan", handler.PulihkanMahasiswa)

	// Route matakuliah
	matakuliah := api.Group("/matakuliah")

	matakuliah.Get("/", handler.DaftarMataKuliah)
	matakuliah.Post("/", handler.BuatMataKuliah)
	matakuliah.Get("/:id", handler.DetailMataKuliah)
	matakuliah.Put("/:id", handler.PerbaruiMataKuliah)
	matakuliah.Delete("/:id", handler.HapusMataKuliah)
}
