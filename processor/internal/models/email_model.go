package models

type GenericEmail struct {
	To              string   `json:"to"`
	Subject         string   `json:"subject"`
	Body            any      `json:"body"`
	BodyTemplate    string   `json:"bodyTemplate"`
	MessageTemplate string   `json:"messageTemplate"`
	Attachments     []string `json:"attachments"`
}

type TransactionEmail struct {
	UserName string `json:"userName"`
	TransactionResume
}
