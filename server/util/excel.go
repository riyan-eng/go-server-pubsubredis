package util

import (
	"io"
	"mime/multipart"

	"github.com/szyhf/go-excel"
)

func ReadImportExcel[T any](file *multipart.FileHeader) (data T) {
	multipartFile, err := file.Open()
	PanicIfNeeded(err)
	byteFile, err := io.ReadAll(multipartFile)
	PanicIfNeeded(err)

	conn := excel.NewConnecter()
	if err := conn.OpenBinary(byteFile); err != nil {
		PanicIfNeeded(err)
	}
	defer conn.Close()

	rd, err := conn.NewReader("Data Import")
	if err != nil {
		PanicIfNeeded(err)
	}
	defer rd.Close()

	err = rd.ReadAll(&data)
	if err != nil {
		PanicIfNeeded(err)
	}
	return
}
