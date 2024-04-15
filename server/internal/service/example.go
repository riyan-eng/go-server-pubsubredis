package service

import (
	"database/sql"
	"strings"

	dtorepository "server/internal/dto_repository"
	"server/internal/entity"
	"server/internal/model"
	"server/internal/repository"
	"server/util"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"github.com/google/uuid"
	"github.com/valyala/fasthttp"
	"github.com/xuri/excelize/v2"
)

type ExampleService interface {
	List(entity.ListExampleReq) entity.ListExampleRes
	Create(*fasthttp.RequestCtx, entity.CreateExampleReq) entity.CreateExampleRes
	Delete(entity.DeleteExampleReq)
	Detail(*fasthttp.RequestCtx, entity.DetailExampleReq) entity.DetailExampleRes
	Put(entity.PutExampleReq)
	Patch(*fasthttp.RequestCtx, entity.PatchExampleReq) entity.DetailExampleRes
	Template() entity.TemplateExampleRes
	Import(entity.ImportExampleReq)
	Pdf() (pdf *wkhtmltopdf.PDFGenerator)
}

type exampleService struct {
	dao repository.DAO
}

func NewExampleService(dao repository.DAO) ExampleService {
	return &exampleService{
		dao: dao,
	}
}

func (t *exampleService) List(req entity.ListExampleReq) (res entity.ListExampleRes) {
	t.dao.NewExampleQuery().List(dtorepository.ListExampleReq{
		Search: req.Search,
		Limit:  req.Limit,
		Offset: req.Offset,
		Order:  req.Order,
		Data:   &res.DataData,
	})
	// err := scan.Rows(&res.DataData, sqlrows)
	// util.PanicIfNeeded(err)
	
	if len(res.DataData) > 0 {
		res.Total = res.DataData[0].TotalRows
	}
	return
}

func (t *exampleService) Create(ctx *fasthttp.RequestCtx, req entity.CreateExampleReq) (res entity.CreateExampleRes) {
	newUUID := uuid.NewString()
	item := model.Example{
		UUID:   newUUID,
		Nama:   sql.NullString{String: req.Nama, Valid: util.IsValid(req.Nama)},
		Detail: sql.NullString{String: req.Detail, Valid: util.IsValid(req.Detail)},
	}
	t.dao.NewExampleQuery().Create(dtorepository.CreateExampleReq{
		Item: item,
	})

	// detail := t.Detail(ctx, entity.DetailExampleReq{UUID: newUUID})
	// res.Item = detail.Data
	return
}

func (t *exampleService) Delete(req entity.DeleteExampleReq) {
	t.dao.NewExampleQuery().Delete(dtorepository.DeleteExampleReq{
		UUID: req.UUID,
	})
}

func (t *exampleService) Detail(ctx *fasthttp.RequestCtx, req entity.DetailExampleReq) (res entity.DetailExampleRes) {
	t.dao.NewExampleQuery().Detail(ctx, dtorepository.DetailExampleReq{
		UUID: req.UUID,
		Data: &res.Data,
	})
	return
}

func (t *exampleService) Put(req entity.PutExampleReq) {
	item := model.Example{
		UUID: req.UUID,
		Nama: sql.NullString{String: req.Nama, Valid: util.IsValid(req.Nama)},
	}
	t.dao.NewExampleQuery().Put(dtorepository.PutExampleReq{
		Item: item,
	})
}

func (t *exampleService) Patch(ctx *fasthttp.RequestCtx, req entity.PatchExampleReq) (res entity.DetailExampleRes) {
	item := model.Example{
		UUID:   req.UUID,
		Nama:   sql.NullString{String: req.Nama, Valid: util.IsValid(req.Nama)},
		Detail: sql.NullString{String: util.Convert().AnyToString(req.Detail), Valid: util.IsValid(req.Detail)},
	}
	t.dao.NewExampleQuery().Patch(ctx, dtorepository.PatchExampleReq{
		Item: item,
	})

	res.Data = t.Detail(ctx, entity.DetailExampleReq{
		UUID: req.UUID,
	}).Data
	return
}

func (t *exampleService) Template() (res entity.TemplateExampleRes) {
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

func (t *exampleService) Import(req entity.ImportExampleReq) {
	var items []model.Example
	for _, i := range req.Items {
		items = append(items, model.Example{
			UUID:   uuid.NewString(),
			Nama:   sql.NullString{String: i.Nama, Valid: util.IsValid(i.Nama)},
			Detail: sql.NullString{String: i.Detail, Valid: util.IsValid(i.Detail)},
		})
	}

	t.dao.NewExampleQuery().Import(dtorepository.ImportExampleReq{
		Items: items,
	})
}

func (t *exampleService) Pdf() (pdf *wkhtmltopdf.PDFGenerator) {

	pdf, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		util.PanicIfNeeded(err)
	}
	pdf.Orientation.Set(wkhtmltopdf.OrientationPortrait)
	pdf.PageSize.Set(wkhtmltopdf.PageSizeA4)
	pdf.MarginTop.Set(20)
	pdf.MarginBottom.Set(20)
	pdf.MarginLeft.Set(20)
	pdf.MarginRight.Set(20)

	template := util.NewTemplate().PDFExample()

	pdf.AddPage(wkhtmltopdf.NewPageReader(strings.NewReader(template)))
	err = pdf.Create()
	if err != nil {
		util.PanicIfNeeded(err)
	}
	return
}
