CREATE TABLE produks (
    id BIGINT unsigned AUTO_INCREMENT PRIMARY KEY,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    nama_produk TEXT,
    slug TEXT,
    harga_reseller TEXT,
    harga_konsumen TEXT,
    stok TEXT,
    deskripsi TEXT,
    id_toko BIGINT unsigned,
    category_id BIGINT unsigned,
    FOREIGN KEY (id_toko) REFERENCES tokos(id) ON DELETE CASCADE,
    FOREIGN KEY (category_id) REFERENCES categories(id) ON DELETE CASCADE
);