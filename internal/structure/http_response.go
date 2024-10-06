package structure

type HttpResponse struct {
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}
