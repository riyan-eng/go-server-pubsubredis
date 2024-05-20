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

func (q *queryStruct) CheckExistingData(object_key, table, uuid string) (ruuid string) {
	query := fmt.Sprintf(`
		select t.uuid from %v t where t.uuid::text = '%v' limit 1
	`, table, uuid)
	sqlrows, err := infrastructure.SqlDB.Query(query)
	if err != nil {
		PanicIfNeeded(err)
	}
	err = scan.Row(&ruuid, sqlrows)
	if err != nil {
		// PanicIfNeeded(NoData{
		// 	Message: fmt.Sprintf("%v with id: %v not found.", object_key, uuid),
		// })
		PanicIfNeeded(CustomBadRequest{
			Messages: infrastructure.Localize("NOT_FOUND"),
		})
	}
	return
}
