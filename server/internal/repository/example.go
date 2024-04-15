package repository

import (
	"database/sql"
	"fmt"

	dtorepository "server/internal/dto_repository"
	"server/internal/model"
	"server/util"

	"github.com/jmoiron/sqlx"
	"github.com/valyala/fasthttp"
	"gorm.io/gorm"
)

type ExampleQuery interface {
	List(dtorepository.ListExampleReq)
	Create(dtorepository.CreateExampleReq)
	Delete(dtorepository.DeleteExampleReq)
	Detail(*fasthttp.RequestCtx, dtorepository.DetailExampleReq)
	Put(dtorepository.PutExampleReq)
	Patch(*fasthttp.RequestCtx, dtorepository.PatchExampleReq)
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

func (t *exampleQuery) List(req dtorepository.ListExampleReq) {
	// query := fmt.Sprintf(`select id, uuid, nama, detail, created_at, updated_at, count(*) over() as total_rows from example 
	// 	where lower(nama) like lower('%%%v%%') order by nama %v limit %v offset %v`,
	// 	req.Search, req.Order, req.Limit, req.Offset)

	err := t.sqlxDB.Select(req.Data, "select id, uuid, nama, detail, created_at, updated_at, count(*) over() as total_rows from example where lower(nama) like lower(%%%$1%%)", req.Search)
	util.PanicSql(err)
	fmt.Println(req.Data)
}

func (t *exampleQuery) Create(req dtorepository.CreateExampleReq) {
	// err := t.gormDB.Create(&req.Item).Error
	// util.PanicSql(err)

	example := ModelExample{
		UUID:   req.Item.UUID,
		Name:   req.Item.Nama,
		Detail: req.Item.Detail,
	}
	_, err := t.sqlxDB.NamedExec("insert into example (uuid, nama, detail) values (:uuid, :nama, :detail)", example)
	util.PanicSql(err)
}

func (t *exampleQuery) Delete(req dtorepository.DeleteExampleReq) {
	err := t.gormDB.Where("uuid = ?", req.UUID).Delete(&model.Example{}).Error
	util.PanicSql(err)
}

func (t *exampleQuery) Detail(ctx *fasthttp.RequestCtx, req dtorepository.DetailExampleReq) {
	err := t.sqlxDB.Get(req.Data, "select uuid, nama, detail from example where uuid=$1", req.UUID)
	util.PanicSql(err)
}

func (t *exampleQuery) Put(req dtorepository.PutExampleReq) {
	err := t.gormDB.Model(&model.Example{}).Select("nama", "detail").Where("uuid = ?", req.Item.UUID).Updates(req.Item).Error
	util.PanicSql(err)
}

func (t *exampleQuery) Patch(ctx *fasthttp.RequestCtx, req dtorepository.PatchExampleReq) {
	// err := t.gormDB.Model(&model.Example{}).Where("uuid = ?", req.Item.UUID).Updates(req.Item).Error
	// util.PanicSql(err)
	example := ModelExample{
		UUID:   req.Item.UUID,
		Name:   req.Item.Nama,
		Detail: req.Item.Detail,
	}
	lala, err := t.sqlxDB.NamedExecContext(ctx, "update example set nama=:nama, detail=coalesce(:detail, detail), updated_at=now() where uuid=:uuid", example)
	fmt.Println(lala)
	util.PanicSql(err)
}

func (t *exampleQuery) Import(req dtorepository.ImportExampleReq) {
	err := t.gormDB.Create(&req.Items).Error
	util.PanicSql(err)
}
