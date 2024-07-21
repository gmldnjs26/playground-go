package types

type ApiResponse struct {
	Result      int64  `json:"result"`
	Description string `json:"description"`
}

func NewApiResponse(result int64, description string) *ApiResponse {
	return &ApiResponse{
		Result:      result,
		Description: description,
	}
}
