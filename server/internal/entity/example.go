package entity

import (
	"server/internal/datastruct"

	"github.com/xuri/excelize/v2"
)

type CreateExampleReq struct {
	Nama   string
	Detail string
}

type ListExampleReq struct {
	Search string
	Limit  int
	Offset int
	Order  string
}

type ListExampleRes struct {
	Items []datastruct.Example
	Total int
}

type DetailExampleReq struct {
	UUID string
}

type DetailExampleRes struct {
	Item datastruct.Example
}

type DeleteExampleReq struct {
	UUID string
}

type PutExampleReq struct {
	UUID   string
	Nama   string
	Detail string
}

type PatchExampleReq struct {
	UUID   string
	Nama   string
	Detail string
}

type TemplateExampleRes struct {
	File *excelize.File
}

type ImportExampleReq struct {
	Items []ImportExampleItemReq
}

type ImportExampleItemReq struct {
	Nama   string
	Detail string
}
