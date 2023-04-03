package dto

type UserReqLogin struct {
	NoTelp    string `json:"no_telp" validate:"required"`
	KataSandi string `json:"kata_sandi" validate:"required"`
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
	Nama         string  `json:"nama" validate:"required"`
	Email        string  `json:"email" validate:"required"`
	NoTelp       string  `json:"no_telp" validate:"required"`
	TanggalLahir *string `json:"tanggal_Lahir" validate:"required"`
	Tentang      *string `json:"tentang" validate:""`
	Pekerjaan    *string `json:"pekerjaan" validate:""`
	KataSandi    string  `json:"kata_sandi" validate:"required"`
	ProvinsiID   string  `json:"id_provinsi" `
	KotaID       string  `json:"id_kota"`
}

type UserReqUpdate struct {
	Nama         string  `json:"nama" validate:""`
	Email        string  `json:"email" validate:""`
	NoTelp       string  `json:"no_telp" validate:""`
	TanggalLahir *string `json:"tanggal_Lahir" validate:""`
	Tentang      *string `json:"tentang" validate:""`
	Pekerjaan    *string `json:"pekerjaan" validate:""`
	KataSandi    string  `json:"kata_sandi" validate:""`
	ProvinsiID   string  `json:"id_provinsi" `
	KotaID       string  `json:"id_kota"`
}
