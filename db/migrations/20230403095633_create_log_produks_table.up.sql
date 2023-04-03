CREATE TABLE log_produks (
    id BIGINT unsigned AUTO_INCREMENT PRIMARY KEY,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    id_produk BIGINT unsigned,
    nama_produk TEXT,
    slug TEXT,
    harga_reseller TEXT,
    harga_konsumen TEXT,
    stok TEXT,
    deskripsi TEXT,
    id_toko BIGINT unsigned,
    id_cateogory BIGINT unsigned,
    FOREIGN KEY (id_produk) REFERENCES produks(id) ON DELETE CASCADE,
    FOREIGN KEY (id_toko) REFERENCES tokos(id) ON DELETE CASCADE,
    FOREIGN KEY (id_cateogory) REFERENCES categories(id) ON DELETE CASCADE
)