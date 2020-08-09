package entity

// Genre entity data
type Genre struct {
	ID     int64  `json:"id"`
	UUID   string `json:"uuid"`
	Name   string `json:"name" validate:"required,max=300"`
	Active bool   `json:"active"`
}
