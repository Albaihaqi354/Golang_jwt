# Golang JWT

Proyek ini adalah implementasi sederhana dari autentikasi dan otorisasi menggunakan JSON Web Token (JWT) dalam bahasa Go. Proyek ini mencakup pembuatan dan validasi token JWT, serta penggunaan middleware untuk melindungi endpoint API.

## Fitur

- **Pembuatan Token**: Menghasilkan token JWT dengan klaim pengguna.
- **Validasi Token**: Memverifikasi token JWT untuk memastikan keabsahan dan integritasnya.
- **Middleware Autentikasi**: Menyediakan middleware untuk melindungi endpoint dengan autentikasi JWT.

## Struktur Proyek

- `controllers/`: Berisi logika pengendali untuk menangani permintaan HTTP.
- `helpers/`: Berisi fungsi-fungsi bantu seperti pembuatan dan validasi token JWT.
- `middleware/`: Berisi middleware untuk autentikasi dan otorisasi.
- `models/`: Berisi definisi model data yang digunakan dalam aplikasi.
- `routes/`: Berisi pengaturan rute API dan penerapan middleware.

## Instalasi

1. **Clone Repositori**:

   ```bash
   git clone https://github.com/Albaihaqi354/Golang_jwt.git
   cd Golang_jwt

2. **Instalasi Dependensi**
    ```bash
    go mod tidy

