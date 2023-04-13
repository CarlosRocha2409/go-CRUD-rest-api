package responses

type PaginationResponse struct {
	Limit int64 `json:"limit"`
	Page  int64 `json:"page"`
	GenericResponse
}
