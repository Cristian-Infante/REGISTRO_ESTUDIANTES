package services

import (
	"github.com/Cristian-Infante/REGISTRO_ESTUDIANTES/interfaces"
	"github.com/Cristian-Infante/REGISTRO_ESTUDIANTES/models"
	"github.com/Cristian-Infante/REGISTRO_ESTUDIANTES/utils"
	"sort"
)

type EstudianteService struct {
	Repositorio interfaces.IEstudianteRepository
}

func NewEstudianteService(repo interfaces.IEstudianteRepository) *EstudianteService {
	return &EstudianteService{Repositorio: repo}
}

// Consulta 1: Estudiante con mejor promedio
func (es *EstudianteService) EstudianteMejorPromedio() (models.Estudiante, float64) {
	estudiantes, _ := es.Repositorio.GetAll()
	var mejorEstudiante models.Estudiante
	mejorPromedio := -1.0

	for _, est := range estudiantes {
		promedio := utils.CalcularPromedioEstudiante(est)
		if promedio > mejorPromedio {
			mejorPromedio = promedio
			mejorEstudiante = est
		}
	}

	return mejorEstudiante, mejorPromedio
}

// Consulta 2: Estudiante con peor promedio
func (es *EstudianteService) EstudiantePeorPromedio() (models.Estudiante, float64) {
	estudiantes, _ := es.Repositorio.GetAll()
	var peorEstudiante models.Estudiante
	peorPromedio := 5.1 // Asumiendo que la nota máxima es 5.0

	for _, est := range estudiantes {
		promedio := utils.CalcularPromedioEstudiante(est)
		if promedio < peorPromedio {
			peorPromedio = promedio
			peorEstudiante = est
		}
	}

	return peorEstudiante, peorPromedio
}

// Consulta 3: Top 10 estudiantes con mejores notas de cada curso
func (es *EstudianteService) Top10MejoresPorCurso() map[string][]models.Estudiante {
	estudiantes, _ := es.Repositorio.GetAll()
	cursos := utils.ObtenerCursos(estudiantes)
	topEstudiantes := make(map[string][]models.Estudiante)

	for _, curso := range cursos {
		estudiantesCurso := utils.ObtenerEstudiantesPorCurso(estudiantes, curso)

		// Ordenar por nota descendente (mejores estudiantes primero)
		sort.Slice(estudiantesCurso, func(i, j int) bool {
			return estudiantesCurso[i].Nota > estudiantesCurso[j].Nota
		})

		limite := 10
		if len(estudiantesCurso) < 10 {
			limite = len(estudiantesCurso)
		}

		var topEstudiantesCurso []models.Estudiante
		for i := 0; i < limite; i++ {
			topEstudiantesCurso = append(topEstudiantesCurso, estudiantesCurso[i].Estudiante)
		}

		topEstudiantes[curso] = topEstudiantesCurso
	}

	return topEstudiantes
}

// Consulta 4: Top 10 estudiantes con peores notas de cada curso
func (es *EstudianteService) Top10PeoresPorCurso() map[string][]models.Estudiante {
	estudiantes, _ := es.Repositorio.GetAll()
	cursos := utils.ObtenerCursos(estudiantes)
	topEstudiantes := make(map[string][]models.Estudiante)

	for _, curso := range cursos {
		estudiantesCurso := utils.ObtenerEstudiantesPorCurso(estudiantes, curso)

		// Ordenar por nota ascendente (peores estudiantes primero)
		sort.Slice(estudiantesCurso, func(i, j int) bool {
			return estudiantesCurso[i].Nota < estudiantesCurso[j].Nota
		})

		limite := 10
		if len(estudiantesCurso) < 10 {
			limite = len(estudiantesCurso)
		}

		var topEstudiantesCurso []models.Estudiante
		for i := 0; i < limite; i++ {
			topEstudiantesCurso = append(topEstudiantesCurso, estudiantesCurso[i].Estudiante)
		}

		topEstudiantes[curso] = topEstudiantesCurso
	}

	return topEstudiantes
}

// Consulta 5: Estudiante masculino de mayor edad
func (es *EstudianteService) EstudianteMasculinoMayorEdad() models.Estudiante {
	estudiantes, _ := es.Repositorio.GetAll()
	var estudianteMayorEdad models.Estudiante
	mayorEdad := -1

	for _, est := range estudiantes {
		if est.Genero == "male" && est.Edad > mayorEdad {
			mayorEdad = est.Edad
			estudianteMayorEdad = est
		}
	}

	return estudianteMayorEdad
}

// Consulta 6: Estudiante femenino de mayor edad
func (es *EstudianteService) EstudianteFemeninoMayorEdad() models.Estudiante {
	estudiantes, _ := es.Repositorio.GetAll()
	var estudianteMayorEdad models.Estudiante
	mayorEdad := -1

	for _, est := range estudiantes {
		if est.Genero == "female" && est.Edad > mayorEdad {
			mayorEdad = est.Edad
			estudianteMayorEdad = est
		}
	}

	return estudianteMayorEdad
}

// Consulta 7: Estadísticas por curso
func (es *EstudianteService) EstadisticasPorCurso() []utils.EstadisticasCurso {
	estudiantes, _ := es.Repositorio.GetAll()
	return utils.EstadisticasPorCurso(estudiantes)
}

// Consulta 8: Estudiantes matriculados en el año pasado (2022)
func (es *EstudianteService) EstudiantesMatriculadosEnAnio(anio int) []models.Estudiante {
	estudiantes, _ := es.Repositorio.GetAll()
	var estudiantesAnio []models.Estudiante

	for _, est := range estudiantes {
		// Convertir la fecha de string a time.Time
		fechaMatricula, err := utils.ParsearFecha(est.Matriculado)
		if err != nil {
			// Si no se puede parsear la fecha, continuar con el siguiente estudiante
			continue
		}

		// Comparar el año de matriculación
		if fechaMatricula.Year() == anio {
			estudiantesAnio = append(estudiantesAnio, est)
		}
	}

	return estudiantesAnio
}

// Consulta 9: Promedio de estudiantes por rango de edad
func (es *EstudianteService) PromedioPorRangoEdad() map[string]float64 {
	estudiantes, _ := es.Repositorio.GetAll()
	return utils.PromedioPorRangoEdad(estudiantes)
}

func (es *EstudianteService) CalcularPromedioDeEstudiante(est models.Estudiante) float64 {
	return utils.CalcularPromedioEstudiante(est)
}

func (es *EstudianteService) CalcularPromedioDeEstudianteEnCurso(est models.Estudiante, curso string) float64 {
	for _, c := range est.Cursos {
		if c.Nombre == curso {
			return c.Nota
		}
	}
	return 0.0
}
