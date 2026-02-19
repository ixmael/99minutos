package domain

type PaginationRequest struct {
	Page  int
	Limit int
}

type PaginationResult[T any] struct {
	Data []T `json:"data"`
}
