package dtos

type ResponsePaging struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Total   interface{}
	Page    interface{}
	Limit   interface{}
}
