package models

type EventTrigger struct {
	Event     TriggerEvent `json:"event"`
	CreatedAt string       `json:"created_at"`
	ID        string       `json:"id"`
	Trigger   Trigger      `json:"trigger"`
	Table     Table        `json:"table"`
}
type SessionVariables struct {
}
type Old struct {
}
type New struct {
}
type Data struct {
	Old Old `json:"old"`
	New New `json:"new"`
}
type TriggerEvent struct {
	SessionVariables SessionVariables `json:"session_variables"`
	Op               string           `json:"op"`
	Data             Data             `json:"data"`
}
type Trigger struct {
	Name string `json:"name"`
}
type Table struct {
	Schema string `json:"schema"`
	Name   string `json:"name"`
}
