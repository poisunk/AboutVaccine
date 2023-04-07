package handler

type RespBody struct {
	Code    int64       `json:"status"`
	Message string      `json:"msg"`
	Data    interface{} `json:"data,omitempty"`
}
