package models

type GenericEmail struct {
	From            string   `json:"from"`
	To              string   `json:"to"`
	Subject         string   `json:"subject"`
	Body            any      `json:"body"`
	BodyTemplate    string   `json:"bodyTemplate"`
	MessageTemplate string   `json:"messageTemplate"`
	Attachments     []string `json:"attachments"`
}
