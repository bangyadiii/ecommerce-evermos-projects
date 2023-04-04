package dto

import "mime/multipart"

type TokoResp struct {
	ID       uint    `json:"id"`
	UserID   uint    `json:"id_user,omitempty"`
	Nama     string  `json:"nama_toko"`
	PhotoUrl *string `json:"url_foto"`
}

type TokoReqUpdate struct {
	Nama  string                `json:"nama_toko" form:"nama_toko"`
	Photo *multipart.FileHeader `form:"-"`
}

type FilterToko struct {
	Nama  string `json:"nama_toko"`
	Limit int    `json:"limit"`
	Page  int    `json:"page"`
}
