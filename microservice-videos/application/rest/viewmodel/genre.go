package viewmodel

// Genre presentation model
type Genre struct {
	ID     string `json:"id" mapper:"UUID"`
	Name   string `json:"name"`
	Active bool   `json:"active"`
}
