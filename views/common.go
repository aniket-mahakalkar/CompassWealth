package views

type Response struct {
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

type ListResponse struct {
	Total   int    `json:"total"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}
