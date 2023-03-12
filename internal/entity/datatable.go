package entity

import "strings"

type (
	Datatable struct {
		Sort       DatatableSort
		Pagination DatatablePagination
	}
)

func (filter *Datatable) Validate() {
	filter.Sort.Validate()
	filter.Pagination.Validate()
}

func (filter *Datatable) IsPaginated() bool {
	return !filter.Pagination.IsDisabled()
}

func (filter *Datatable) NameFilter(name string) string {
	return "%" + strings.ToLower(name) + "%"
}

type (
	DatatableSort struct {
		Field     string
		Direction string
	}
)

func (sort *DatatableSort) Validate() {
	if !strings.EqualFold(sort.Direction, "desc") {
		sort.Direction = "asc"
	}
}

type (
	DatatablePagination struct {
		Disable bool
		Page    int
		Limit   int
		Offset  int
	}
)

func (pagination *DatatablePagination) Validate() {
	if pagination.IsDisabled() {
		return
	}
	if !(pagination.Limit > 0) {
		pagination.Limit = 10
	}
	pagination.SetOffset()
}

func (pagination *DatatablePagination) SetOffset() {
	if pagination.IsDisabled() {
		return
	}
	if pagination.Page > 0 {
		pagination.Offset = (pagination.Page - 1) * pagination.Limit
	}
}

func (pagination *DatatablePagination) IsDisabled() bool {
	return pagination.Disable
}
