package pagination

import (
	"math"

	api "gitlab.ubrato.ru/ubrato/core/api/gen"
	"golang.org/x/exp/constraints"
)

const (
	Page    = 0
	PerPage = 100
)

type Pagination struct {
	Found   uint64
	Page    uint64
	Pages   uint64
	PerPage uint64
}

func New(page, perPage, found uint64) Pagination {
	return Pagination{
		Found:   found,
		Page:    page,
		Pages:   CalculatePages(found, perPage),
		PerPage: perPage,
	}
}

// CalculatePages вычисляет количество страниц для пагинации
func CalculatePages[T1, T2 constraints.Float | constraints.Integer](total T1, perPage T2) uint64 {
	if perPage <= 0 {
		return 0
	}

	return uint64(math.Ceil(float64(total) / float64(perPage)))
}

func ConvertPaginationToAPI(p Pagination) api.Pagination {
	return api.Pagination{
		Found:   int(p.Found),
		Page:    int(p.Page),
		Pages:   int(p.Pages),
		PerPage: int(p.PerPage),
	}
}
