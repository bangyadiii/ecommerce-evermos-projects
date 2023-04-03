CREATE TABLE alamats (
    id BIGINT unsigned AUTO_INCREMENT PRIMARY KEY,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    id_user BIGINT unsigned,
    judul TEXT,
    nama_penerima TEXT,
    no_telp VARCHAR(20) UNIQUE,
    detail_alamat TEXT,
    FOREIGN KEY (id_user) REFERENCES users(id)
);
