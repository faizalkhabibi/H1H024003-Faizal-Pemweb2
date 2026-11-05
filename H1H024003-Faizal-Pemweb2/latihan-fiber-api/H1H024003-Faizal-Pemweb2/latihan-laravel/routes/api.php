<?php

use App\Http\Controllers\Api\AuthController;
use App\Http\Controllers\Api\MahasiswaController;
use App\Http\Middleware\PeranAdmin;
use Illuminate\Support\Facades\Route;

Route::post('/auth/register', [AuthController::class, 'register']);

// Tambahkan middleware throttle:5,1 pada rute login
Route::post('/auth/login', [AuthController::class, 'login'])->middleware('throttle:5,1');

Route::middleware('auth:sanctum')->group(function () {
    Route::get('/auth/profil', [AuthController::class, 'profil']);
    Route::put('/auth/password', [AuthController::class, 'updatePassword']); // Rute Tugas 1
    Route::post('/auth/logout', [AuthController::class, 'logout']);
    Route::post('/auth/logout-semua', [AuthController::class, 'logoutSemua']);

    Route::get('/mahasiswa', [MahasiswaController::class, 'index']);
    Route::get('/mahasiswa/{mahasiswa}', [MahasiswaController::class, 'show']);

    Route::middleware('ability:mahasiswa:tulis')->group(function () {
        Route::post('/mahasiswa', [MahasiswaController::class, 'store']);
        Route::put('/mahasiswa/{mahasiswa}', [MahasiswaController::class, 'update']);
        Route::patch('/mahasiswa/{mahasiswa}', [MahasiswaController::class, 'update']);
        
        // Rute Tugas 2 dengan Middleware Kustom PeranAdmin
        Route::delete('/mahasiswa/{mahasiswa}', [MahasiswaController::class, 'destroy'])
            ->middleware(PeranAdmin::class);
    });
});