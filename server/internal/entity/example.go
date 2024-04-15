package entity

import (
	"server/internal/datastruct"

	"github.com/xuri/excelize/v2"
)

type CreateExampleReq struct {
	Nama   string
	Detail string
}

type CreateExampleRes struct {
	Item datastruct.Example
}

type ListExampleReq struct {
	Search string
	Limit  int
	Offset int
	Order  string
}

type ListExampleRes struct {
	// Data []datastruct.Example
	DataData  []datastruct.ListExample
	Total int
}

type DetailExampleReq struct {
	UUID string
}

type DetailExampleRes struct {
	Data datastruct.DetailExample
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
	Detail any
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
