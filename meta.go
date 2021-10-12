package posmoni

// Meta represents the response status code and message after called APIs.
type Meta struct {
	Code        int    `json:"code"`
	Message     string `json:"message"`
	CurrentPage int    `json:"current_page"`
	NextPage    int    `json:"next_page"`
	PrevPage    int    `json:"prev_page"`
	TotalPages  int    `json:"total_pages"`
	TotalCount  int    `json:"total_count"`
}
