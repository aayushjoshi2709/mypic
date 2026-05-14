package common

type ErrorResponseDto struct {
	Error string `json:"error"`
}

type PaginatedResponseDto[T any] struct {
	Data       []T   `json:"data"`
	Page       int64 `json:"page"`
	Limit      int64 `json:"limit"`
	TotalPages int64 `json:"totalPages"`
}

func (p *PaginatedResponseDto[T]) Init(data []T, page, limit, TotalPages int64) {
	p.Data = data
	p.Page = page
	p.Limit = limit
	p.TotalPages = TotalPages
}
