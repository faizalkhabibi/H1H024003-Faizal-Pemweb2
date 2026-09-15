<?php
namespace Database\Seeders;

use Illuminate\Database\Seeder;
use App\Models\Matakuliah;

class MatakuliahSeeder extends Seeder
{
    public function run(): void
    {
        $daftar = [
            ['kode' => 'MK01', 'nama' => 'Pemrograman Web II', 'sks' => 3, 'semester' => 4],
            ['kode' => 'MK02', 'nama' => 'Sistem Mikrokontroler', 'sks' => 3, 'semester' => 4],
        ];
        foreach ($daftar as $item) {
            Matakuliah::create($item);
        }
    }
}