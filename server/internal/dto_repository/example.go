package dtorepository

import "server/internal/model"

type CreateExampleReq struct {
	Item model.Example
}

type ListExampleReq struct {
	Search string
	Limit  int
	Offset int
	Order  string
}

type DetailExampleReq struct {
	ID int
}

type DeleteExampleReq struct {
	ID int
}

type PutExampleReq struct {
	Item model.Example
}

type PatchExampleReq struct {
	Item model.Example
}
