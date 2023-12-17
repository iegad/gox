package web

type BasicResponse struct {
	Code  int32       `json:"code"`
	Error string      `json:"error,omitempty"`
	Data  interface{} `json:"data,omitempty"`
}
