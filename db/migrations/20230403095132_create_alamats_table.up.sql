CREATE TABLE alamats (
    id BIGINT unsigned AUTO_INCREMENT PRIMARY KEY,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    id_user BIGINT unsigned,
    judul VARCHAR(255),
    nama_penerima VARCHAR(255),
    no_telp VARCHAR(20),
    detail_alamat VARCHAR(255),
    FOREIGN KEY (id_user) REFERENCES users(id)
);
