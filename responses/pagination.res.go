package responses

type PaginationResponse struct {
	Limit int `json:"limit"`
	Page  int `json:"page"`
	GenericResponse
}
