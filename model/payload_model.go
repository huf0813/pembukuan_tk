package model

type Payload struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (p *Payload) NewPayload(status, message string, data interface{}) *Payload {
	return &Payload{
		Status:  status,
		Message: message,
		Data:    data,
	}
}
