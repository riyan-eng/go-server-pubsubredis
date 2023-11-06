package dtorepository

import "server/internal/model"

type CreateObjectReq struct {
	Item model.Object
}

type ListObjectReq struct {
	Search string
	Limit  int
	Offset int
	Order  string
}

type DetailObjectReq struct {
	ID int
}

type DeleteObjectReq struct {
	ID int
}

type PutObjectReq struct {
	Item model.Object
}

type PatchObjectReq struct {
	Item model.Object
}
