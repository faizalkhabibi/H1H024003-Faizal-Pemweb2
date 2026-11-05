package model

import "time"

type MataKuliah struct {
    ID        uint      `gorm:"primaryKey" json:"id"`
    Kode      string    `gorm:"size:15;uniqueIndex" json:"kode"`
    Nama      string    `gorm:"size:100" json:"nama"`
    SKS       int       `json:"sks"`
    Semester  int       `json:"semester"`
    CreatedAt time.Time `json:"dibuat_pada"`
    UpdatedAt time.Time `json:"diubah_pada"`
}