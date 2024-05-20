package util

import (
	"math"
)

type pageMeta struct {
	Page   int
	Limit  int
	Offset int
}

func PageMeta(page, limit int) pageMeta {
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 1
	}
	offset := limit * (page - 1)
	return pageMeta{
		Page:   page,
		Limit:  limit,
		Offset: offset,
	}
}

type paginationStruct struct{}

func NewPagination() *paginationStruct {
	return &paginationStruct{}
}

func (p *paginationStruct) GetCountPages(totalRows, limit int) (countPages int) {
	countPages = int(math.Ceil(float64(totalRows) / float64(limit)))
	return
}

func (p *paginationStruct) GetPageMeta(page, limit int) pageMeta {
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 1
	}
	return pageMeta{
		Page:   page,
		Limit:  limit,
		Offset: limit * (page - 1),
	}
}
