package web

type basicResponse struct {
	Code  int32       `json:"code"`
	Error string      `json:"error,omitempty"`
	Data  interface{} `json:"data,omitempty"`
}

func NewResponse(code int32, err string, data interface{}) *basicResponse {
	return &basicResponse{
		Code:  code,
		Error: err,
		Data:  data,
	}
}
