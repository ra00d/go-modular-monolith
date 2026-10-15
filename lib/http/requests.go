package http

type PaginationParams struct {
	Page    int `json:"page,omitempty"    from:"query" d:"1"`
	PerPage int `json:"perPage,omitempty" from:"query" d:"20"`
}
type SearchParams struct {
	Query string `json:"query" from:"query"`
}

type SortParams struct {
	Field string `json:"field,omitempty" from:"query"`
	Dir   string `json:"dir,omitempty"   from:"query" v:"in:desc,asc" d:"asc"`
}

type RequestBaseParams struct {
	PaginationParams
	SearchParams
	Sort SortParams `json:"sort"`
}
