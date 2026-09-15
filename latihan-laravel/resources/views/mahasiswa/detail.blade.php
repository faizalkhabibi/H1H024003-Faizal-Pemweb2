@extends('layouts.app')
@section('judul', 'Detail Mahasiswa')
@section('konten')
<h1 class="h3 mb-4">Detail Mahasiswa: {{ $mahasiswa->nama }}</h1>
<p><strong>NIM:</strong> {{ $mahasiswa->nim }}</p>
<p><strong>Program Studi:</strong> {{ $mahasiswa->programStudi->nama }}</p>

<table class="table table-bordered bg-white mt-3">
    <thead>
        <tr>
            <th>Kode MK</th>
            <th>Nama Mata Kuliah</th>
            <th>SKS</th>
            <th>Nilai</th>
        </tr>
    </thead>
    <tbody>
        @foreach ($mahasiswa->matakuliah as $mk)
        <tr>
            <td>{{ $mk->kode }}</td>
            <td>{{ $mk->nama }}</td>
            <td>{{ $mk->sks }}</td>
            <td>{{ $mk->pivot->nilai ?? '-' }}</td>
        </tr>
        @endforeach
    </tbody>
</table>
@endsection