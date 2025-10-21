# Cara Menjalankan

## 1. Clone repo
```bash
git clone https://github.com/Alvi19/backend-golang-test-article.git
cd backend-golang-test-article
```

## 2. Setup Makefile
- Copy Makefile.example ke Makefile:
```bash
cp Makefile.example Makefile
```
- Edit Makefile dan sesuaikan koneksi database Anda:
```bash
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=password   # ganti sesuai password postgres Anda
DB_NAME=task_management
DB_SSLMODE=disable
```

## Install dependency
```bash
go mod tidy
```

## 4. Buat database (jika belum ada)
```bash
CREATE DATABASE task_management;
```

## 5. Jalankan migration
```bash
make migrate-up
```

## 6. Jalankan project
```bash
make run
```
## 7. Atau sekalian migrate + run:
```bash
make run-migrate
```
## 8. API Documentation
- Setelah server berjalan, buka [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html) untuk melihat dokumentasi API menggunakan Swagger UI.
- Untuk mengenerate ulang dokumentasi Swagger (jika ada perubahan di kode), jalankan:
```bash
bashmake swagger-gen
```