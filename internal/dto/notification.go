package dto

type EventNotification struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Direction string `json:"direction"`
}

func NewEventNotification(firstName string, lastName string, direction string) *EventNotification {
	return &EventNotification{FirstName: firstName, LastName: lastName, Direction: direction}
}
