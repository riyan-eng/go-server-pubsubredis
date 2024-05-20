package dtorepository

import (
	"server/internal/datastruct"
	"server/internal/model"
)

type CreateExampleReq struct {
	ModelExample model.ModelExample
}

type ListExampleReq struct {
	Search string
	Limit  int
	Offset int
	Order  string
}

type ListExampleRes struct {
	Data      []datastruct.ListExample
	CountRows int
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
	ModelExample model.ModelExample
}

type PatchExampleReq struct {
	Item         model.Example
	ModelExample model.ModelExample
}

type ImportExampleReq struct {
	Items []model.Example
}
