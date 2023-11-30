package util

import (
	"fmt"

	"server/infrastructure"

	"github.com/blockloop/scan/v2"
)

type queryStruct struct{}

func NewQuery() *queryStruct {
	return &queryStruct{}
}

func (q *queryStruct) GetIDByUUID(table, uuid string) (id int) {
	if uuid == "" {
		return
	}
	query := fmt.Sprintf(`
		select t.id from %v t where t.uuid::text = '%v'
	`, table, uuid)
	sqlrows, err := infrastructure.SqlDB.Query(query)
	if err != nil {
		PanicIfNeeded(err)
	}
	err = scan.Row(&id, sqlrows)
	if err != nil {
		PanicIfNeeded(BadRequest{
			Message: fmt.Sprintf("%v tidak ditemukan.", table),
		})
	}
	return
}

func (q *queryStruct) CheckExistingData(object_key, table, uuid string) (ruuid string) {
	query := fmt.Sprintf(`
		select t.uuid from %v t where t.uuid::text = '%v' limit 1
	`, table, uuid)
	sqlrows, err := infrastructure.SqlDB.Query(query)
	if err != nil {
		PanicIfNeeded(err)
	}
	err = scan.Row(&uuid, sqlrows)
	if err != nil {
		PanicIfNeeded(BadRequest{
			Message: fmt.Sprintf("%v dengan id: %v tidak ditemukan.", object_key, uuid),
		})
	}
	return
}
