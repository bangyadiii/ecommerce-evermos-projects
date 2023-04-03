CREATE TABLE users (
    id BIGINT unsigned AUTO_INCREMENT PRIMARY KEY,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    nama TEXT,
    email VARCHAR(255) UNIQUE,
    no_telp VARCHAR(20) UNIQUE,
    tanggal_lahir TEXT,
    tentang TEXT,
    pekerjaan TEXT,
    kata_sandi TEXT,
    id_provinsi TEXT,
    id_kota TEXT,
    is_admin BOOLEAN
);
