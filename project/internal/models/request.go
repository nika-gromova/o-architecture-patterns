package models

type Request struct {
	Method string            `json:"method"`
	Header map[string]string `json:"header"`
	Body   []byte            `json:"body"`
}

const (
	LocaleVariable = "locale"
	TimeVariable   = "time"
)

type HeaderVariable struct {
	Name   string
	Header string
	Type   string
}
