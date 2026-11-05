package config

import (
    "os"

    "github.com/joho/godotenv"
)

type Konfigurasi struct {
    NamaAplikasi string
    Port         string
    Lingkungan   string
    ApiKey       string

    DbHost  string
    DbPort  string
    DbNama  string
    DbUser  string
    DbSandi string
}

func ambil(kunci, nilaiDefault string) string {
    nilai := os.Getenv(kunci)

    if nilai == "" {
        return nilaiDefault
    }

    return nilai
}

func Muat() Konfigurasi {
    godotenv.Load()

    return Konfigurasi{
        NamaAplikasi: ambil("APP_NAME", "Fiber API"),
        Port:         ambil("APP_PORT", "3000"),
        Lingkungan:   ambil("APP_ENV", "local"),
        ApiKey:       ambil("API_KEY", "belajar-fiber-123"),

        DbHost:  ambil("DB_HOST", "127.0.0.1"),
        DbPort:  ambil("DB_PORT", "3306"),
        DbNama:  ambil("DB_NAME", "pemweb2_fiber"),
        DbUser:  ambil("DB_USER", "root"),
        DbSandi: ambil("DB_PASS", ""),
    }
}