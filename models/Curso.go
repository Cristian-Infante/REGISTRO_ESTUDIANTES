package models

type Curso struct {
	ID     int     `json:"id"`
	Nombre string  `json:"curso"`
	Nota   float64 `json:"nota"`
}
