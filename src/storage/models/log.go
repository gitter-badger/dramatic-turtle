package models

// Log struct
type Log struct {
	Task  string `json:"task"`
	Start int    `json:"start"`
	End   int    `json:"end"`
}
