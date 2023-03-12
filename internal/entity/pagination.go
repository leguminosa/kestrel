package entity

type (
	BasicPagination struct {
		Page  int   `json:"page"`
		Row   int   `json:"row"`
		Count int64 `json:"count"`
	}
)
