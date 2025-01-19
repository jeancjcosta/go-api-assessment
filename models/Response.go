package models

type Response struct {
	Value   int    `json:"value"`
	Index   int    `json:"index"`
	Message string `json:"message,omitempty"`
}
