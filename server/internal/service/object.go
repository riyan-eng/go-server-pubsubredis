package service

import (
	"database/sql"

	dtorepository "server/internal/dto_repository"
	"server/internal/entity"
	"server/internal/model"
	"server/internal/repository"
	"server/util"

	"github.com/blockloop/scan/v2"
)

type ObjectService interface {
	List(entity.ListObjectReq) entity.ListObjectRes
	Create(entity.CreateObjectReq)
	Delete(entity.DeleteObjectReq)
	Detail(entity.DetailObjectReq) entity.DetailObjectRes
	Put(entity.PutObjectReq)
	Patch(entity.PatchObjectReq)
}

type objectService struct {
	dao repository.DAO
}

func NewObjectService(dao repository.DAO) ObjectService {
	return &objectService{
		dao: dao,
	}
}

func (t *objectService) List(req entity.ListObjectReq) (res entity.ListObjectRes) {
	sqlrows := t.dao.NewObjectQuery().List(dtorepository.ListObjectReq{
		Search: req.Search,
		Limit:  req.Limit,
		Offset: req.Offset,
		Order:  req.Order,
	})
	err := scan.Rows(&res.Items, sqlrows)
	util.PanicIfNeeded(err)

	if len(res.Items) > 0 {
		res.Total = res.Items[0].TotalRows
	}
	return
}

func (t *objectService) Create(req entity.CreateObjectReq) {
	item := model.Object{
		UUID:     req.UUID,
		Bucket:   sql.NullString{String: req.Bukcet, Valid: util.IsValid(req.Bukcet)},
		Nama:     sql.NullString{String: req.Nama, Valid: util.IsValid(req.Nama)},
		Size:     sql.NullInt64{Int64: req.Size, Valid: util.IsValid(int(req.Size))},
		MimeType: sql.NullString{String: req.MimeType, Valid: util.IsValid(req.MimeType)},
		Url:      sql.NullString{String: req.Url, Valid: util.IsValid(req.Url)},
		Path:     sql.NullString{String: req.Path, Valid: util.IsValid(req.Path)},
	}
	t.dao.NewObjectQuery().Create(dtorepository.CreateObjectReq{
		Item: item,
	})
}

func (t *objectService) Delete(req entity.DeleteObjectReq) {
	t.dao.NewObjectQuery().Delete(dtorepository.DeleteObjectReq{
		ID: req.ID,
	})
}

func (t *objectService) Detail(req entity.DetailObjectReq) (res entity.DetailObjectRes) {
	sqlrows := t.dao.NewObjectQuery().Detail(dtorepository.DetailObjectReq{
		ID: req.ID,
	})
	err := scan.Row(&res.Item, sqlrows)
	util.PanicIfNeeded(err)
	res.Item.SizeString = util.NewFile().GetFileSizeString(res.Item.Size)
	return
}

func (t *objectService) Put(req entity.PutObjectReq) {
	item := model.Object{
		ID: req.ID,
		// Nama: req.Nama,
	}
	t.dao.NewObjectQuery().Put(dtorepository.PutObjectReq{
		Item: item,
	})
}

func (t *objectService) Patch(req entity.PatchObjectReq) {
	item := model.Object{
		ID: req.ID,
		// Nama: req.Nama,
	}
	t.dao.NewObjectQuery().Patch(dtorepository.PatchObjectReq{
		Item: item,
	})
}
