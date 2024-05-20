package util

import (
	"github.com/lib/pq"
)

type isValidStruct struct{}

func NewIsValid() *isValidStruct {
	return &isValidStruct{}
}

func (i *isValidStruct) String(arg string) bool {
	return arg != ""
}

func (i *isValidStruct) Int(arg int) bool {
	return arg != 0
}

func (i *isValidStruct) Float64(arg float64) bool {
	return arg != 0
}

func (i *isValidStruct) Float32(arg float32) bool {
	return arg != 0
}

func (i *isValidStruct) Int64(num int64) bool {
	return num != 0
}

func (i *isValidStruct) Int32(num int32) bool {
	return num != 0
}

func (i *isValidStruct) Any(input interface{}) bool {
	return input != nil
}

func (i *isValidStruct) ErrUniqViol(err error) bool {
	if err != nil {
		pqErr, ok := err.(*pq.Error)
		if ok {
			if pqErr.Code.Name() == "unique_violation" {
				return true
			}
		}
	}
	return false
}
