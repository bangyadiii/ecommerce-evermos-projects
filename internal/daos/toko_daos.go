package daos

import "gorm.io/gorm"

type Toko struct {
	gorm.Model
	UserID   uint   `json:"id_user" gorm:"id_user"`
	Name     string `json:"nama_toko" gorm:"nama_toko"`
	PhotoUrl *string `json:"url_foto" gorm:"url_foto"`
	User     User   `gorm:"constraint:OnDelete:CASCADE,OnUpdate:Cascade"`
}
type FilterToko struct {
	Limit, Offset int
	Name          string `json:"nama_toko"`
}
