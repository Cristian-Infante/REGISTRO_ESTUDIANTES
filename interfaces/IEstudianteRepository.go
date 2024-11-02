package interfaces

import "github.com/Cristian-Infante/REGISTRO_ESTUDIANTES/models"

type IEstudianteRepository interface {
	GetAll() ([]models.Estudiante, error)
}
