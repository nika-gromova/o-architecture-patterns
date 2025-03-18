package models

type Request struct {
	Method string            `json:"method"`
	Header map[string]string `json:"header"`
	Body   []byte            `json:"body"`
}
