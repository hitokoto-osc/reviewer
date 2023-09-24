package model

type Page[T any] struct {
	Total      int `json:"total" dc:"总数"`
	Page       int `json:"page" dc:"页码"`
	PageSize   int `json:"page_size" dc:"每页数量"`
	Collection []T `json:"collection" dc:"数据"`
}
