@extends('layouts.app')

@section('judul', 'Detail Matakuliah')

@section('konten')
<h1 class="h3 mb-4">Detail Matakuliah</h1>
<div class="card">
    <div class="card-body">
        <p><strong>Kode:</strong> {{ $matakuliah['kode'] }}</p>
        <p><strong>Nama Matakuliah:</strong> {{ $matakuliah['nama'] }}</p>
        <p><strong>Jumlah SKS:</strong> <x-badge-sks :sks="$matakuliah['sks']" /></p>
    </div>
</div>
<a href="{{ route('matakuliah.index') }}" class="btn btn-secondary mt-3">Kembali</a>
@endsection