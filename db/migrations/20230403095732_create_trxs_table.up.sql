CREATE TABLE trxs (
    id BIGINT unsigned AUTO_INCREMENT PRIMARY KEY,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    id_user BIGINT unsigned,
    alamat_pengiriman BIGINT unsigned,
    harga_total INT,
    kode_invoice VARCHAR(255),
    metode_bayar VARCHAR(100),
    FOREIGN KEY (id_user) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (alamat_pengiriman) REFERENCES alamats(id)
);