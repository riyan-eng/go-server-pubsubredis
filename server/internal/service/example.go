package service

import (
	"database/sql"
	"strings"

	"server/infrastructure"
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
	List(*fasthttp.RequestCtx, entity.ListExampleReq) entity.ListExampleRes
	Create(*fasthttp.RequestCtx, entity.CreateExampleReq) entity.DetailExampleRes
	Detail(*fasthttp.RequestCtx, entity.DetailExampleReq) entity.DetailExampleRes
	Put(*fasthttp.RequestCtx, entity.PutExampleReq) entity.DetailExampleRes
	Patch(*fasthttp.RequestCtx, entity.PatchExampleReq) entity.DetailExampleRes
	Delete(*fasthttp.RequestCtx, entity.DeleteExampleReq)
	Template(*fasthttp.RequestCtx) entity.TemplateExampleRes
	Import(*fasthttp.RequestCtx, entity.ImportExampleReq)
	Pdf(*fasthttp.RequestCtx) (pdf *wkhtmltopdf.PDFGenerator)
}

type exampleService struct {
	dao repository.DAO
}

func NewExampleService(dao repository.DAO) ExampleService {
	return &exampleService{
		dao: dao,
	}
}

func (t *exampleService) List(ctx *fasthttp.RequestCtx, req entity.ListExampleReq) (res entity.ListExampleRes) {
	repo, err := t.dao.NewExampleQuery().List(ctx, dtorepository.ListExampleReq{
		Search: req.Search,
		Limit:  req.Limit,
		Offset: req.Offset,
		Order:  req.Order,
	})
	if err != nil {
		util.PanicIfNeeded(util.CustomBadRequest{
			StatusCodes: 500,
		})
	}

	res.Data = repo.Data
	res.CountRows = repo.CountRows
	return
}

func (t *exampleService) Create(ctx *fasthttp.RequestCtx, req entity.CreateExampleReq) (res entity.DetailExampleRes) {
	modelExample := model.ModelExample{
		UUID:   req.UUID,
		Name:   sql.NullString{String: req.Nama, Valid: util.NewIsValid().String(req.Nama)},
		Detail: sql.NullString{String: req.Detail, Valid: util.NewIsValid().String(req.Detail)},
	}
	err := t.dao.NewExampleQuery().Create(ctx, dtorepository.CreateExampleReq{
		ModelExample: modelExample,
	})
	if util.NewIsValid().ErrUniqViol(err) {
		util.PanicIfNeeded(util.CustomBadRequest{
			StatusCodes: 409,
		})
	}

	if err != nil {
		util.PanicIfNeeded(util.CustomBadRequest{
			Errors:      err.Error(),
			Messages:    infrastructure.Localize("FAILED_CREATE"),
			StatusCodes: 500,
		})
	}

	res.Data = t.Detail(ctx, entity.DetailExampleReq{
		UUID: req.UUID,
	}).Data
	return
}

func (t *exampleService) Delete(ctx *fasthttp.RequestCtx, req entity.DeleteExampleReq) {
	err:=t.dao.NewExampleQuery().Delete(ctx, dtorepository.DeleteExampleReq{
		UUID: req.UUID,
	})
	if err != nil {
		util.PanicIfNeeded(util.CustomBadRequest{
			Errors:      err.Error(),
			Messages:    infrastructure.Localize("FAILED_UPDATE"),
			StatusCodes: 500,
		})
	}
}

func (t *exampleService) Detail(ctx *fasthttp.RequestCtx, req entity.DetailExampleReq) (res entity.DetailExampleRes) {
	repo, err := t.dao.NewExampleQuery().Detail(ctx, dtorepository.DetailExampleReq{
		UUID: req.UUID,
	})
	if err != nil && err == sql.ErrNoRows {
		util.PanicIfNeeded(util.CustomBadRequest{
			Messages: infrastructure.Localize("NOT_FOUND"),
		})
	} else if err != nil {
		util.PanicIfNeeded(util.CustomBadRequest{
			Errors:      err.Error(),
			StatusCodes: 500,
		})
	}

	res.Data = repo.Data
	return
}

func (t *exampleService) Put(ctx *fasthttp.RequestCtx, req entity.PutExampleReq) (res entity.DetailExampleRes) {
	modelExample := model.ModelExample{
		UUID:   req.UUID,
		Name:   sql.NullString{String: req.Nama, Valid: util.NewIsValid().String(req.Nama)},
		Detail: sql.NullString{String: req.Detail, Valid: util.NewIsValid().String(req.Detail)},
	}
	err := t.dao.NewExampleQuery().Put(ctx, dtorepository.PutExampleReq{
		ModelExample: modelExample,
	})

	if util.NewIsValid().ErrUniqViol(err) {
		util.PanicIfNeeded(util.CustomBadRequest{
			StatusCodes: 409,
		})
	}

	if err != nil {
		util.PanicIfNeeded(util.CustomBadRequest{
			Errors:      err.Error(),
			Messages:    infrastructure.Localize("FAILED_UPDATE"),
			StatusCodes: 500,
		})
	}

	res.Data = t.Detail(ctx, entity.DetailExampleReq{
		UUID: req.UUID,
	}).Data
	return
}

func (t *exampleService) Patch(ctx *fasthttp.RequestCtx, req entity.PatchExampleReq) (res entity.DetailExampleRes) {
	modelExample := model.ModelExample{
		UUID:   req.UUID,
		Name:   sql.NullString{String: req.Nama, Valid: util.NewIsValid().String(req.Nama)},
		Detail: sql.NullString{String: util.Convert().AnyToString(req.Detail), Valid: util.NewIsValid().Any(req.Detail)},
	}
	err := t.dao.NewExampleQuery().Patch(ctx, dtorepository.PatchExampleReq{
		ModelExample: modelExample,
	})
	if util.NewIsValid().ErrUniqViol(err) {
		util.PanicIfNeeded(util.CustomBadRequest{
			StatusCodes: 409,
		})
	}

	if err != nil {
		util.PanicIfNeeded(util.CustomBadRequest{
			Errors:      err.Error(),
			Messages:    infrastructure.Localize("FAILED_UPDATE"),
			StatusCodes: 500,
		})
	}

	res.Data = t.Detail(ctx, entity.DetailExampleReq{
		UUID: req.UUID,
	}).Data
	return
}

func (t *exampleService) Template(ctx *fasthttp.RequestCtx) (res entity.TemplateExampleRes) {
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

func (t *exampleService) Import(ctx *fasthttp.RequestCtx, req entity.ImportExampleReq) {
	var items []model.Example
	for _, i := range req.Items {
		items = append(items, model.Example{
			UUID:   uuid.NewString(),
			Nama:   sql.NullString{String: i.Nama, Valid: util.NewIsValid().String(i.Nama)},
			Detail: sql.NullString{String: i.Detail, Valid: util.NewIsValid().String(i.Detail)},
		})
	}

	t.dao.NewExampleQuery().Import(dtorepository.ImportExampleReq{
		Items: items,
	})
}

func (t *exampleService) Pdf(ctx *fasthttp.RequestCtx) (pdf *wkhtmltopdf.PDFGenerator) {

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
