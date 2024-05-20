package repository

import (
	"database/sql"
	"fmt"

	"server/infrastructure"
	dtorepository "server/internal/dto_repository"
	"server/util"

	"github.com/blockloop/scan/v2"
	"github.com/jmoiron/sqlx"
	"github.com/valyala/fasthttp"
	"gorm.io/gorm"
)

type ExampleQuery interface {
	List(*fasthttp.RequestCtx, dtorepository.ListExampleReq) (dtorepository.ListExampleRes, error)
	Create(*fasthttp.RequestCtx, dtorepository.CreateExampleReq) error
	Delete(*fasthttp.RequestCtx, dtorepository.DeleteExampleReq) error
	Detail(*fasthttp.RequestCtx, dtorepository.DetailExampleReq) (dtorepository.DetailExampleRes, error)
	Put(*fasthttp.RequestCtx, dtorepository.PutExampleReq) error
	Patch(*fasthttp.RequestCtx, dtorepository.PatchExampleReq) error
	Import(dtorepository.ImportExampleReq)
}

type exampleQuery struct {
	sqlDB  *sql.DB
	gormDB *gorm.DB
	sqlxDB *sqlx.DB
}

type ModelExample struct {
	UUID   string         `db:"uuid"`
	Name   sql.NullString `db:"nama"`
	Detail sql.NullString `db:"detail"`
}

type DataExample struct {
	UUID   string `db:"uuid"`
	Name   string `db:"nama"`
	Detail any    `db:"detail"`
}

func (t *exampleQuery) List(ctx *fasthttp.RequestCtx, req dtorepository.ListExampleReq) (res dtorepository.ListExampleRes, err error) {
	query := fmt.Sprintf(`
	select id, uuid, nama, detail, created_at, updated_at, count(uuid) over() as total_rows from example 
	where lower(nama) like lower('%%%v%%') order by nama %v limit %v offset %v
	`, req.Search, req.Order, req.Limit, req.Offset)
	sqlRows, err := t.sqlDB.QueryContext(ctx, query)
	if err != nil {
		return
	}
	err = scan.Rows(&res.Data, sqlRows)
	if err != nil {
		return
	}

	if len(res.Data) > 0 {
		res.CountRows = res.Data[0].TotalRows
	}
	return
}

func (t *exampleQuery) Create(ctx *fasthttp.RequestCtx, req dtorepository.CreateExampleReq) (err error) {
	sqlRslt, err := t.sqlxDB.NamedExecContext(ctx, "insert into example (uuid, nama, detail) values (:uuid, :nama, :detail)", req.ModelExample)
	if err != nil {
		return
	}
	rowsAffected, err := sqlRslt.RowsAffected()
	if err != nil {
		return
	}
	if rowsAffected == 0 {
		util.PanicIfNeeded(util.CustomBadRequest{
			Messages:    infrastructure.Localize("FAILED_INSERT_NO_DATA"),
			StatusCodes: 400,
		})
	}
	return
}

func (t *exampleQuery) Delete(ctx *fasthttp.RequestCtx, req dtorepository.DeleteExampleReq) (err error) {
	sqlRslt, err := t.sqlxDB.ExecContext(ctx, "delete from example where uuid = $1", req.UUID)
	if err != nil {
		return
	}
	rowsAffected, err := sqlRslt.RowsAffected()
	if err != nil {
		return
	}
	if rowsAffected == 0 {
		util.PanicIfNeeded(util.CustomBadRequest{
			Messages:    infrastructure.Localize("FAILED_DELETE_NO_DATA"),
			StatusCodes: 400,
		})
	}
	return
}

func (t *exampleQuery) Detail(ctx *fasthttp.RequestCtx, req dtorepository.DetailExampleReq) (res dtorepository.DetailExampleRes, err error) {
	query := fmt.Sprintf(`
	select uuid, nama, detail from example where uuid::text='%v'
	`, req.UUID)
	sqlRows, err := t.sqlDB.QueryContext(ctx, query)
	if err != nil {
		return
	}
	err = scan.Row(&res.Data, sqlRows)
	if err != nil {
		return
	}

	return
}

func (t *exampleQuery) Put(ctx *fasthttp.RequestCtx, req dtorepository.PutExampleReq) (err error) {
	sqlRslt, err := t.sqlxDB.NamedExecContext(ctx, "update example set nama=:nama, detail=:detail, updated_at=now() where uuid=:uuid", req.ModelExample)
	if err != nil {
		return
	}
	rowsAffected, err := sqlRslt.RowsAffected()
	if err != nil {
		return
	}
	if rowsAffected == 0 {
		util.PanicIfNeeded(util.CustomBadRequest{
			Messages:    infrastructure.Localize("FAILED_UPDATE_NO_DATA"),
			StatusCodes: 400,
		})
	}
	return
}

func (t *exampleQuery) Patch(ctx *fasthttp.RequestCtx, req dtorepository.PatchExampleReq) (err error) {
	sqlRslt, err := t.sqlxDB.NamedExecContext(ctx, "update example set nama=:nama, detail=coalesce(:detail, detail), updated_at=now() where uuid=:uuid", req.ModelExample)
	fmt.Println("err: ", err)
	if err != nil {
		return
	}
	rowsAffected, err := sqlRslt.RowsAffected()
	if err != nil {
		return
	}
	if rowsAffected == 0 {
		util.PanicIfNeeded(util.CustomBadRequest{
			Messages:    infrastructure.Localize("FAILED_UPDATE_NO_DATA"),
			StatusCodes: 400,
		})
	}
	return
}

func (t *exampleQuery) Import(req dtorepository.ImportExampleReq) {
	err := t.gormDB.Create(&req.Items).Error
	util.PanicSql(err)
}
