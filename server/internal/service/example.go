package service

import (
	"database/sql"

	dtorepository "server/internal/dto_repository"
	dtoservice "server/internal/dto_service"
	"server/internal/model"
	"server/internal/repository"
	"server/pkg/util"

	"github.com/blockloop/scan/v2"
	"github.com/google/uuid"
	"github.com/xuri/excelize/v2"
)

type ExampleService interface {
	List(dtoservice.ListExampleReq) dtoservice.ListExampleRes
	Create(dtoservice.CreateExampleReq)
	Delete(dtoservice.DeleteExampleReq)
	Detail(dtoservice.DetailExampleReq) dtoservice.DetailExampleRes
	Put(dtoservice.PutExampleReq)
	Patch(dtoservice.PatchExampleReq)
	Template() dtoservice.TemplateExampleRes
	Import(dtoservice.ImportExampleReq)
}

type exampleService struct {
	dao repository.DAO
}

func NewExampleService(dao repository.DAO) ExampleService {
	return &exampleService{
		dao: dao,
	}
}

func (t *exampleService) List(req dtoservice.ListExampleReq) (res dtoservice.ListExampleRes) {
	sqlrows := t.dao.NewExampleQuery().List(dtorepository.ListExampleReq{
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

func (t *exampleService) Create(req dtoservice.CreateExampleReq) {
	item := model.Example{
		UUID:   uuid.NewString(),
		Nama:   sql.NullString{String: req.Nama, Valid: util.ValidIsNotBlankString(req.Nama)},
		Detail: sql.NullString{String: req.Detail, Valid: util.ValidIsNotBlankString(req.Detail)},
	}
	t.dao.NewExampleQuery().Create(dtorepository.CreateExampleReq{
		Item: item,
	})
}

func (t *exampleService) Delete(req dtoservice.DeleteExampleReq) {
	t.dao.NewExampleQuery().Delete(dtorepository.DeleteExampleReq{
		UUID: req.UUID,
	})
}

func (t *exampleService) Detail(req dtoservice.DetailExampleReq) (res dtoservice.DetailExampleRes) {
	sqlrows := t.dao.NewExampleQuery().Detail(dtorepository.DetailExampleReq{
		UUID: req.UUID,
	})
	err := scan.Row(&res.Item, sqlrows)
	util.PanicIfNeeded(err)
	return
}

func (t *exampleService) Put(req dtoservice.PutExampleReq) {
	item := model.Example{
		UUID: req.UUID,
		Nama: sql.NullString{String: req.Nama, Valid: util.ValidIsNotBlankString(req.Nama)},
	}
	t.dao.NewExampleQuery().Put(dtorepository.PutExampleReq{
		Item: item,
	})
}

func (t *exampleService) Patch(req dtoservice.PatchExampleReq) {
	item := model.Example{
		UUID: req.UUID,
		Nama: sql.NullString{String: req.Nama, Valid: util.ValidIsNotBlankString(req.Nama)},
	}
	t.dao.NewExampleQuery().Patch(dtorepository.PatchExampleReq{
		Item: item,
	})
}

func (t *exampleService) Template() (res dtoservice.TemplateExampleRes) {
	f, err := excelize.OpenFile("./media/excel/Template Example.xlsx")
	if err != nil {
		util.PanicIfNeeded(err)
		return
	}
	defer func() {
		// Close the spreadsheet.
		if err := f.Close(); err != nil {
			util.PanicIfNeeded(err)
		}
	}()

	res.File = f
	return
}

func (t *exampleService) Import(req dtoservice.ImportExampleReq) {
	var items []model.Example
	for _, i := range req.Items {
		items = append(items, model.Example{
			UUID:   uuid.NewString(),
			Nama:   sql.NullString{String: i.Nama, Valid: util.ValidIsNotBlankString(i.Nama)},
			Detail: sql.NullString{String: i.Detail, Valid: util.ValidIsNotBlankString(i.Detail)},
		})
	}

	t.dao.NewExampleQuery().Import(dtorepository.ImportExampleReq{
		Items: items,
	})
}
