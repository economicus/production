package models

type QueryOption struct {
	Page    int    `form:"page" json:"page"`
	PerPage int    `form:"per_page" json:"per_page"`
	Order   string `form:"order" json:"order"`
	Query   string `form:"query" json:"query"`
}

func NewQueryOption() *QueryOption {
	return &QueryOption{
		Page:    0,
		PerPage: 10,
		Order:   "",
		Query:   "",
	}
}
