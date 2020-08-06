package viewmodel

// Category presentation model
type Category struct {
	ID          string `json:"id" mapper:"UUID"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Active      bool   `json:"active"`
}
