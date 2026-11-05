package handler

import (
    "strconv"

    "github.com/gofiber/fiber/v3"
    "gorm.io/gorm"

    "latihan-fiber-api/internal/model"
    "latihan-fiber-api/internal/repository"
)

func DaftarMataKuliah(c fiber.Ctx) error {
    daftar, err := repository.AmbilSemuaMataKuliah()

    if err != nil {
        return fiber.NewError(
            fiber.StatusInternalServerError,
            "Gagal mengambil data mata kuliah",
        )
    }

    return c.JSON(fiber.Map{
        "sukses": true,
        "data": daftar,
    })
}

func DetailMataKuliah(c fiber.Ctx) error {
    id, err := strconv.Atoi(c.Params("id"))

    if err != nil || id < 1 {
        return fiber.NewError(
            fiber.StatusBadRequest,
            "ID tidak valid",
        )
    }

    matakuliah, err := repository.AmbilMataKuliahPerId(uint(id))

    if err != nil {
        if err == gorm.ErrRecordNotFound {
            return fiber.NewError(
                fiber.StatusNotFound,
                "Mata kuliah tidak ditemukan",
            )
        }

        return fiber.NewError(
            fiber.StatusInternalServerError,
            "Gagal mengambil data",
        )
    }

    return c.JSON(fiber.Map{
        "sukses": true,
        "data": matakuliah,
    })
}

func BuatMataKuliah(c fiber.Ctx) error {
    var permintaan PermintaanMataKuliah

    if err := c.Bind().Body(&permintaan); err != nil {
        return fiber.NewError(
            fiber.StatusBadRequest,
            "Format body tidak valid",
        )
    }

    matakuliah := model.MataKuliah{
        Kode:     permintaan.Kode,
        Nama:     permintaan.Nama,
        SKS:      permintaan.SKS,
        Semester: permintaan.Semester,
    }

    if err := repository.SimpanMataKuliah(&matakuliah); err != nil {
        return fiber.NewError(
            fiber.StatusUnprocessableEntity,
            "Gagal menyimpan mata kuliah",
        )
    }

    return c.Status(fiber.StatusCreated).JSON(fiber.Map{
        "sukses": true,
        "pesan": "Mata kuliah berhasil dibuat",
        "data": matakuliah,
    })
}

func PerbaruiMataKuliah(c fiber.Ctx) error {
    id, err := strconv.Atoi(c.Params("id"))

    if err != nil || id < 1 {
        return fiber.NewError(
            fiber.StatusBadRequest,
            "ID tidak valid",
        )
    }

    matakuliah, err := repository.AmbilMataKuliahPerId(uint(id))

    if err != nil {
        return fiber.NewError(
            fiber.StatusNotFound,
            "Mata kuliah tidak ditemukan",
        )
    }

    var permintaan PermintaanMataKuliah

    if err := c.Bind().Body(&permintaan); err != nil {
        return fiber.NewError(
            fiber.StatusBadRequest,
            "Format body tidak valid",
        )
    }

    if permintaan.Kode != "" {
        matakuliah.Kode = permintaan.Kode
    }

    if permintaan.Nama != "" {
        matakuliah.Nama = permintaan.Nama
    }

    if permintaan.SKS > 0 {
        matakuliah.SKS = permintaan.SKS
    }

    if permintaan.Semester > 0 {
        matakuliah.Semester = permintaan.Semester
    }

    if err := repository.PerbaruiMataKuliah(&matakuliah); err != nil {
        return fiber.NewError(
            fiber.StatusUnprocessableEntity,
            "Gagal memperbarui mata kuliah",
        )
    }

    return c.JSON(fiber.Map{
        "sukses": true,
        "pesan": "Mata kuliah berhasil diperbarui",
        "data": matakuliah,
    })
}

func HapusMataKuliah(c fiber.Ctx) error {
    id, err := strconv.Atoi(c.Params("id"))

    if err != nil || id < 1 {
        return fiber.NewError(
            fiber.StatusBadRequest,
            "ID tidak valid",
        )
    }

    matakuliah, err := repository.AmbilMataKuliahPerId(uint(id))

    if err != nil {
        return fiber.NewError(
            fiber.StatusNotFound,
            "Mata kuliah tidak ditemukan",
        )
    }

    if err := repository.HapusMataKuliah(&matakuliah); err != nil {
        return fiber.NewError(
            fiber.StatusInternalServerError,
            "Gagal menghapus mata kuliah",
        )
    }

    return c.JSON(fiber.Map{
        "sukses": true,
        "pesan": "Mata kuliah berhasil dihapus",
    })
}