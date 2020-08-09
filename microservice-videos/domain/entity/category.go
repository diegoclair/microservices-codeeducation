package entity

// Category entity data
type Category struct {
	ID          int64  `json:"id"`
	UUID        string `json:"uuid"`
	Name        string `json:"name" validate:"required,max=300"`
	Description string `json:"description" validate:"max=8000"`
	Active      bool   `json:"active"`
}
