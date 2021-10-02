package model

type Result struct {
	Data  interface{}
	Error error
}

type Response struct {
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}
