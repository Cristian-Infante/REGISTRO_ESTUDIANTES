package models

type Estudiante struct {
	Index       int     `json:"index"`
	Nombre      string  `json:"nombre"`
	Apellido    string  `json:"apellido"`
	Edad        int     `json:"edad"`
	Genero      string  `json:"gender"`
	Email       string  `json:"email"`
	Telefono    string  `json:"phone"`
	Direccion   string  `json:"address"`
	AcercaDe    string  `json:"about"`
	Matriculado string  `json:"matriculado"`
	Cursos      []Curso `json:"cursos"`
}
