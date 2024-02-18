package util

import (
	"strings"

	"github.com/lib/pq"
)

func PanicIfNeeded(err interface{}) {
	if err != nil {
		panic(err)
	}
}

func PanicBodyValidation(err error, listErr any) {
	if err != nil {
		PanicIfNeeded(BodyValidationError{
			Message:   MESSAGE_FAILED_VALIDATION,
			ListError: listErr,
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
		PanicIfNeeded(err)
	}
}
