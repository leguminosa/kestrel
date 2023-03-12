package entity

type (
	GearSetOption struct {
		ID       int    `json:"id" db:"id"`
		Name     string `json:"name" db:"name"`
		SetCount int    `json:"set_count" db:"set_count"`
	}
)
