package common

type Pagination struct {
	Page    int
	PerPage int
	Offset  int
	Cursor  any
}

type PaginationMeta struct {
	Count      int  `json:"count"`
	Cursor     any  `json:"cursor"`
	Page       int  `json:"page"`
	PerPage    int  `json:"perPage"`
	TotalPages int  `json:"totalPages"`
	NextPage   int  `json:"nextPage"`
	PrevPage   int  `json:"prevPage"`
	LastPage   int  `json:"lastPage"`
	HasNext    bool `json:"hasNext"`
}

func GetPaginationMeta(count, page, perPage int) PaginationMeta {
	if page < 1 {
		page = 1
	}

	if perPage < 1 {
		perPage = 10
	}

	totalPages := (count + perPage - 1) / perPage

	if totalPages == 0 {
		totalPages = 1
	}

	return PaginationMeta{
		Count:      count,
		Page:       page,
		PerPage:    perPage,
		TotalPages: totalPages,
		NextPage:   min(page+1, totalPages),
		PrevPage:   max(page-1, 1),
		LastPage:   totalPages,
		HasNext:    page < totalPages,
	}
}

func (p *PaginationMeta) SetCursor[C ~int | ~string](cursor C) {
	if isEmpty(cursor) {
		return
	}

	p.Cursor = cursor
}
