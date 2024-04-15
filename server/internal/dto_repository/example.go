package dtorepository

import (
	"server/internal/datastruct"
	"server/internal/model"
)

type CreateExampleReq struct {
	Item model.Example
}

type ListExampleReq struct {
	Search string
	Limit  int
	Offset int
	Order  string
	Data   *[]datastruct.ListExample
}

type DetailExampleReq struct {
	UUID string
	Data *datastruct.DetailExample
}

type DeleteExampleReq struct {
	UUID string
}

type PutExampleReq struct {
	Item model.Example
}

type PatchExampleReq struct {
	Item model.Example
}

type ImportExampleReq struct {
	Items []model.Example
}
