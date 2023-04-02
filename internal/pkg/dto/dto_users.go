package dto

type UserReqLogin struct {
	NoTelp   string `json:"no_telp" validate:"required"`
	Password string `json:"kata_sandi" validate:"required"`
}

type UserResLogin struct {
	Nama         string  `json:"nama"`
	Email        string  `json:"email"`
	NoTelp       string  `json:"no_telp"`
	TanggalLahir *string `json:"tanggal_Lahir"`
	Tentang      *string `json:"tentang" `
	Pekerjaan    *string `json:"pekerjaan"`
	Token        string  `json:"token"`
}

type UserReqRegister struct {
	Nama         string  `json:"nama"`
	Email        string  `json:"email"`
	NoTelp       string  `json:"no_telp"`
	TanggalLahir *string `json:"tanggal_Lahir"`
	Tentang      *string `json:"tentang" `
	Pekerjaan    *string `json:"pekerjaan"`
	Password     string  `json:"kata_sandi"`
	ProvinsiID   string  `json:"id_provinsi"`
	KotaID       string  `json:"id_kota"`
}
