package repository

import (
    "latihan-fiber-api/internal/config"
    "latihan-fiber-api/internal/model"
)

type HasilProgramStudi struct {
    ID             uint   `json:"id"`
    Kode           string `json:"kode"`
    Nama           string `json:"nama"`
    Jenjang        string `json:"jenjang"`
    JumlahMahasiswa int64 `json:"jumlah_mahasiswa"`
}

func AmbilProgramStudi() ([]HasilProgramStudi, error) {
    var daftar []HasilProgramStudi

    err := config.DB.
        Model(&model.ProgramStudi{}).
        Select(`
            program_studis.id,
            program_studis.kode,
            program_studis.nama,
            program_studis.jenjang,
            COUNT(mahasiswas.id) AS jumlah_mahasiswa
        `).
        Joins(`
            LEFT JOIN mahasiswas
            ON mahasiswas.program_studi_id = program_studis.id
        `).
        Group(`
            program_studis.id,
            program_studis.kode,
            program_studis.nama,
            program_studis.jenjang
        `).
        Find(&daftar).Error

    return daftar, err
}