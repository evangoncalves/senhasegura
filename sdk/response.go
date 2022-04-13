package sdk

type IResponse interface {
	Unmarshal(msg []byte) error
}

type Response struct {
	Status    int    `json:"status,omitempty"`
	Message   string `json:"message,omitempty"`
	Error     bool   `json:"error,omitempty"`
	ErrorCode int    `json:"error_code,omitempty"`
}
