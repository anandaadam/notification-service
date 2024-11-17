package helpers

type MessageBody struct {
	EventType string `json:"eventType"`
	MediaType string `json:"mediaType"`
	Subject   string `json:"subject"`
	SendTo    string `json:"sendTo"`
	Message   string `json:"message"`
}
