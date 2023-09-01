package models

type ResponseMessage struct {
	Message string `json:"message"`
}

type AuthHeader struct {
	ClientId string `header:"client_id"`
}
