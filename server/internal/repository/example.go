package repository

import (
	"database/sql"
	"fmt"

	dtorepository "server/internal/dto_repository"
	"server/internal/model"
	"server/pkg/util"

	"gorm.io/gorm"
)

type ExampleQuery interface {
	List(dtorepository.ListExampleReq) *sql.Rows
	Create(dtorepository.CreateExampleReq)
	Delete(dtorepository.DeleteExampleReq)
	Detail(dtorepository.DetailExampleReq) *sql.Rows
	Put(dtorepository.PutExampleReq)
	Patch(dtorepository.PatchExampleReq)
	Import(dtorepository.ImportExampleReq)
}

type exampleQuery struct {
	sqlDB  *sql.DB
	gormDB *gorm.DB
}

func (t *exampleQuery) List(req dtorepository.ListExampleReq) *sql.Rows {
	query := fmt.Sprintf(`select id, uuid, nama, detail, created_at, updated_at, count(*) over() as total_rows from example 
		where lower(nama) like lower('%%%v%%') order by nama %v limit %v offset %v`,
		req.Search, req.Order, req.Limit, req.Offset)
	rows, err := t.sqlDB.Query(query)
	util.PanicSql(err)
	return rows
}

func (t *exampleQuery) Create(req dtorepository.CreateExampleReq) {
	err := t.gormDB.Create(&req.Item).Error
	util.PanicSql(err)
}

func (t *exampleQuery) Delete(req dtorepository.DeleteExampleReq) {
	err := t.gormDB.Where("uuid = ?", req.UUID).Delete(&model.Example{}).Error
	util.PanicSql(err)
}

func (t *exampleQuery) Detail(req dtorepository.DetailExampleReq) *sql.Rows {
	query := fmt.Sprintf(`
		select id, uuid, nama, detail, created_at, updated_at from example where uuid = '%v'
	`, req.UUID)
	rows, err := t.sqlDB.Query(query)
	util.PanicSql(err)
	return rows
}

func (t *exampleQuery) Put(req dtorepository.PutExampleReq) {
	err := t.gormDB.Model(&model.Example{}).Select("nama", "detail").Where("uuid = ?", req.Item.UUID).Updates(req.Item).Error
	util.PanicSql(err)
}

func (t *exampleQuery) Patch(req dtorepository.PatchExampleReq) {
	err := t.gormDB.Model(&model.Example{}).Where("uuid = ?", req.Item.UUID).Updates(req.Item).Error
	util.PanicSql(err)
}

func (t *exampleQuery) Import(req dtorepository.ImportExampleReq) {
	err := t.gormDB.Create(&req.Items).Error
	util.PanicSql(err)
}
