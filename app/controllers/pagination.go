package controllers

import (
	"github.com/morkid/paginate"
)

func NewPagination() *paginate.Pagination {
	return paginate.New(&paginate.Config{
		DefaultSize:          5,
		Operator:             "AND",
		ErrorEnabled:         true,
		FieldSelectorEnabled: true,
		CustomParamEnabled:   true,
		SizeParams:           []string{"per_page"},
		FieldsParams:         []string{"select_by"},
		FilterParams:         []string{"filter_by"},
		SortParams:           []string{"sort_by"},
	})
}
