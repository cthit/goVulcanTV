package common

type Response struct {
	Success bool `json:"success"`
	Data interface{} `json:"data"`
	Error string `json:"error"`
}
