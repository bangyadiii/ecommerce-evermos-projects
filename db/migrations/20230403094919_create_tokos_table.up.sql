CREATE TABLE tokos (
    id BIGINT unsigned AUTO_INCREMENT PRIMARY KEY,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    id_user BIGINT unsigned,
    nama_toko TEXT,
    url_foto TEXT,
    FOREIGN KEY (id_user) REFERENCES users(id) ON DELETE CASCADE
);