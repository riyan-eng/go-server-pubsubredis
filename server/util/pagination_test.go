package util

import (
	"testing"
)

func TestGetTotalPages(t *testing.T) {
	p := NewPagination()
	totalRows := 10
	limit := 5
	expectedTotalPages := 2

	totalPages := p.GetCountPages(totalRows, limit)

	if totalPages != expectedTotalPages {
		t.Errorf("GetTotalPages(%d, %d) = %d; want %d", totalRows, limit, totalPages, expectedTotalPages)
	}
}

func TestGetPageMeta(t *testing.T) {
	p := NewPagination()
	page := 2
	limit := 5
	expectedPageMeta := pageMeta{Page: 2, Limit: 5, Offset: 5}

	pageMeta := p.GetPageMeta(page, limit)

	if pageMeta != expectedPageMeta {
		t.Errorf("GetPageMeta(%d, %d) = %v; want %v", page, limit, pageMeta, expectedPageMeta)
	}
}
