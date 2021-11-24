package entity

type HttpResponse struct {
	Code uint16      `json:"code"`
	Data interface{} `json:"data"`
}
