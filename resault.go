package paginator

type Result struct {
	Total       int64 `json:"total" xml:"total"`
	Data        any   `json:"data" xml:"data"`
	PerPage     int64 `json:"per_page" xml:"perPage"`
	CurrentPage int64 `json:"current_page" xml:"currentPage"`
	LastPage    int64 `json:"last_page" xml:"lastPage"`
}
