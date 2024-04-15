package util

import (
	"database/sql"
	"server/infrastructure"
	"strings"

	"github.com/lib/pq"
)

func PanicIfNeeded(err interface{}) {
	if err != nil {
		panic(err)
	}
}

func PanicBodyValidation(errors any, err error) {
	if err != nil {
		PanicIfNeeded(CustomBadRequest{
			Errors:   errors,
			Messages: infrastructure.Localize("FAILED_VALIDATION"),
		})
	}
}

func PanicSql(err error) {
	if err != nil {
		pqErr, ok := err.(*pq.Error)
		if ok {
			if pqErr.Code.Name() == "unique_violation" {
				temp1 := strings.Split(pqErr.Detail, "=")
				temp2 := strings.ReplaceAll(temp1[1], "(", "")
				temp2 = strings.ReplaceAll(temp2, ")", "")
				PanicIfNeeded(Duplicate{
					Message: temp2,
				})
			}
		}
		if err == sql.ErrNoRows {
			PanicIfNeeded(NoData{
				Message: err.Error(),
			})
		}
		PanicIfNeeded(err)
	}
}
