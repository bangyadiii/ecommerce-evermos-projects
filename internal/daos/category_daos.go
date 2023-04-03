package daos

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Nama string `json:"nama_category" gorm:"column:nama_category;type:varchar(255)"`
}
