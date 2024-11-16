package helpers

type MessageBody struct {
	EventType string `json:"eventType"`
	MediaType string `json:"mediaType"`
	SendTo    string `json:"sendTo"`
	Message   string `json:"message"`
}
