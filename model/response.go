package model

type Response struct {
	Status  int    `json:"status"`
	Data    any    `json:"data"`
	Total   int64  `json:"total"`
	Message string `json:"message"`
}

type LoginResponse struct {
	Status  int    `json:"status"`
	Data    string `json:"data"`
	ID      uint   `json:"id"`
	Message string `json:"message"`
	Token   string `json:"token"`
}
