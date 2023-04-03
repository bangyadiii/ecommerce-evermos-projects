package dto

import "ecommerce-evermos-projects/internal/daos"

type AlamatRes struct {
	ID           uint   `json:"id"`
	UserID       uint   `json:"id_user"`
	Judul        string `json:"judul_alamat"`
	NamaPenerima string `json:"nama_penerima"`
	NoTelp       string `json:"no_telp"`
	DetailAlamat string `json:"detail_alamat"`
}

type AlamatReqCreate struct {
	UserID       uint   `json:"id_user,omitempty"`
	Judul        string `json:"judul_alamat" validate:"required"`
	NamaPenerima string `json:"nama_penerima" validate:"required"`
	NoTelp       string `json:"no_telp" validate:"required"`
	DetailAlamat string `json:"detail_alamat" validate:"required"`
}

type AlamatReqUpdate struct {
	Judul        string `json:"judul_alamat"`
	NamaPenerima string `json:"nama_penerima"`
	NoTelp       string `json:"no_telp"`
	DetailAlamat string `json:"detail_alamat"`
}

func AlamatDaosToDto(alamat daos.Alamat) AlamatRes {
	return AlamatRes{
		ID:           alamat.ID,
		UserID:       alamat.UserID,
		Judul:        alamat.Judul,
		NamaPenerima: alamat.NamaPenerima,
		NoTelp:       alamat.NoTelp,
		DetailAlamat: alamat.DetailAlamat,
	}
}
func AlamatSliceDaosToDto(alamat []*daos.Alamat) (res []*AlamatRes) {
	for _, v := range alamat {
		var val AlamatRes = AlamatDaosToDto(*v)
		res = append(res, &val)
	}
	return res
}
