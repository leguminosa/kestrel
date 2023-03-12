package entity

import "time"

type (
	GearSetOption struct {
		ID         int       `json:"id" db:"id"`
		Name       string    `json:"name" db:"name"`
		SetCount   int       `json:"set_count" db:"set_count"`
		Status     int       `json:"status" db:"status"`
		CreateTime time.Time `json:"create_time" db:"create_time"`
		UpdateTime time.Time `json:"update_time,omitempty" db:"update_time"`
	}
	GearSetOptionFilter struct {
		Datatable Datatable
		ID        int
		Name      string
		SetCount  int
		StatusRaw string
		Status    int
	}
	GearSetOptionResult struct {
		List []GearSetOption `json:"list"`
		BasicPagination
	}
)

func (filter *GearSetOptionFilter) Validate() {
	filter.Datatable.Validate()
}

func (filter *GearSetOptionFilter) NameFilter() string {
	return filter.Datatable.NameFilter(filter.Name)
}

func (filter *GearSetOptionFilter) UseStatusFilter() bool {
	return filter.StatusRaw != ""
}
