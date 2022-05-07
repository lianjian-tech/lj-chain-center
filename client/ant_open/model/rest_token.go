package model

type RestToken struct {
	Success bool        `json:"success"`
	Code    string      `json:"code"`
	Data    interface{} `json:"data"`
}
