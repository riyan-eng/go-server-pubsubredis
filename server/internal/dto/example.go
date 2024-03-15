package dto

type CreateExample struct {
	Nama   string `json:"nama" valid:"required;min:8;in:ADMIN,STAFF"`
	Detail string `json:"detail" valid:"date:yyyy-mm-dd"`
}

type PutExample struct {
	Nama   string `json:"nama" valid:"required"`
	Detail string `json:"detail"`
}

type PatchExample struct {
	Nama   string `json:"nama" valid:"required"`
	Detail string `json:"detail"`
}

type ImportExample struct {
	Nama   string `xlsx:"column(Nama)"`
	Detail string `xlsx:"column(Detail)"`
}
