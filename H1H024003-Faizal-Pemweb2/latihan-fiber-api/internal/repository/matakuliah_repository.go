package repository

import (
    "latihan-fiber-api/internal/config"
    "latihan-fiber-api/internal/model"
)

func AmbilSemuaMataKuliah() ([]model.MataKuliah, error) {
    var daftar []model.MataKuliah

    err := config.DB.
        Order("semester ASC").
        Find(&daftar).Error

    return daftar, err
}

func AmbilMataKuliahPerId(id uint) (model.MataKuliah, error) {
    var matakuliah model.MataKuliah

    err := config.DB.First(&matakuliah, id).Error

    return matakuliah, err
}

func SimpanMataKuliah(matakuliah *model.MataKuliah) error {
    return config.DB.Create(matakuliah).Error
}

func PerbaruiMataKuliah(matakuliah *model.MataKuliah) error {
    return config.DB.Save(matakuliah).Error
}

func HapusMataKuliah(matakuliah *model.MataKuliah) error {
    return config.DB.Delete(matakuliah).Error
}