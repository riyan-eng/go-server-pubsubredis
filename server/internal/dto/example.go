package dto

type CreateExample struct {
	Nama   string `json:"nama" valid:"required"`
	Detail string `json:"detail"`
}

type PutExample struct {
	Nama   string `json:"nama" valid:"required"`
	Detail string `json:"detail"`
}

type PatchExample struct {
	Nama   string `json:"nama" valid:"required"`
	Detail any    `json:"detail"`
}

type ImportExample struct {
	Nama   string `xlsx:"column(Nama)"`
	Detail string `xlsx:"column(Detail)"`
}
