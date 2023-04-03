CREATE TABLE foto_produks (
    id BIGINT unsigned AUTO_INCREMENT PRIMARY KEY,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    id_produk BIGINT unsigned,
    FOREIGN KEY (id_produk) REFERENCES produks(id) ON DELETE CASCADE
);