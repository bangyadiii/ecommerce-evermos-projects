package dto

import "ecommerce-evermos-projects/internal/daos"

type CategoryReqCreate struct {
	Nama string `json:"nama_category" validate:"required"`
}

type CategoryReqUpdate struct {
	Nama string `json:"nama_category" validate:"required"`
}

type CategoryRes struct {
	ID   uint   `json:"id"`
	Nama string `json:"nama_category"`
}

func CategoryDaosToDto(data daos.Category) CategoryRes {
	return CategoryRes{
		ID:   data.ID,
		Nama: data.Nama,
	}
}

func CategoriesDaosToDto(data []daos.Category) (res []CategoryRes) {
	var conv []CategoryRes
	for _, v := range data {
		res = append(conv, CategoryDaosToDto(v))
	}

	return res
}
