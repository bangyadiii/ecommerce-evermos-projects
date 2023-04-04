package daos

import "gorm.io/gorm"

type Toko struct {
	gorm.Model
	UserID   uint    `json:"id_user" gorm:"column:id_user;foreignKey:id_user;constrait:OnDelete:CASCADE"`
	Name     string  `json:"nama_toko" gorm:"column:nama_toko;type:text"`
	PhotoUrl *string `json:"url_foto" gorm:"column:url_foto;type:text"`
	User     User    ``
}
type FilterToko struct {
	Limit, Offset int
	Name          string `json:"nama_toko"`
}
