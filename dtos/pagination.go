package dtos

type Pagination struct {
	CurrentPage       int `json:"current_page"`
	TotalPage         int `json:"total_page"`
	PageElementsCount int `json:"page_elements_count"`
}
