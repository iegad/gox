package web

type BasicRequest struct {
	UserID     *int32      `json:"user_id,omitempty"`
	Token      *string     `json:"token,omitempty"`
	Idempotent *int64      `json:"idempotent,omitempty"`
	Data       interface{} `json:"data,omitempty"`
}
