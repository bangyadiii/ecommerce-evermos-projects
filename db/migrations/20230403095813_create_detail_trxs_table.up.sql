CREATE TABLE detail_trxs (
    id BIGINT unsigned AUTO_INCREMENT PRIMARY KEY,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    id_trx BIGINT unsigned,
    id_log_produk BIGINT unsigned,
    kuantitas INT,
    harga_total INT,
    FOREIGN KEY (id_trx) REFERENCES trxs(id) ON DELETE CASCADE,
    FOREIGN KEY (id_log_produk) REFERENCES log_produks(id) ON DELETE CASCADE
);