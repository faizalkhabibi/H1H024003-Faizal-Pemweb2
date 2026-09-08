<?php

namespace App\Http\Controllers;

use Illuminate\Http\Request;

class MatakuliahController extends Controller
{
    private array $matakuliah = [
        ['kode' => 'TK101', 'nama' => 'Pemrograman Web II', 'sks' => 3],
        ['kode' => 'TK102', 'nama' => 'Sistem Kendali', 'sks' => 3],
        ['kode' => 'TK103', 'nama' => 'Keamanan Jaringan Komputer', 'sks' => 2],
        ['kode' => 'TK104', 'nama' => 'Internet of Things', 'sks' => 3],
        ['kode' => 'TK105', 'nama' => 'Metode Numerik', 'sks' => 2],
    ];

    public function index(Request $request)
    {
        $katakunci = $request->query('q', '');
        $data = $this->matakuliah;

        if ($katakunci !== '') {
            $data = array_filter($data, function ($item) use ($katakunci) {
                return stripos($item['nama'], $katakunci) !== false || stripos($item['kode'], $katakunci) !== false;
            });
        }

        return view('matakuliah.index', [
            'daftarMatakuliah' => $data,
            'katakunci' => $katakunci
        ]);
    }

    public function show(string $kode)
    {
        $matakuliahDetail = collect($this->matakuliah)->firstWhere('kode', $kode);

        if (!$matakuliahDetail) {
            abort(404);
        }

        return view('matakuliah.show', ['matakuliah' => $matakuliahDetail]);
    }
}