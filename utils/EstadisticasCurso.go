package utils

import (
	"fmt"
	"github.com/Cristian-Infante/REGISTRO_ESTUDIANTES/models"
	"math"
	"time"
)

// ParsearFecha convierte una fecha en formato string a time.Time
func ParsearFecha(fechaStr string) (time.Time, error) {
	layout := "2006-01-02T15:04:05 -07:00" // Ajusta el layout al formato del JSON
	fecha, err := time.Parse(layout, fechaStr)
	if err != nil {
		fmt.Printf("Error al parsear la fecha: %v\n", err)
		return time.Time{}, err
	}
	return fecha, nil
}

// Calcular el promedio de notas de un estudiante
func CalcularPromedioEstudiante(estudiante models.Estudiante) float64 {
	sumaNotas := 0.0
	for _, curso := range estudiante.Cursos {
		sumaNotas += curso.Nota
	}
	return sumaNotas / float64(len(estudiante.Cursos))
}

// Obtener la lista de todos los cursos
func ObtenerCursos(estudiantes []models.Estudiante) []string {
	cursosMap := make(map[string]bool)
	for _, est := range estudiantes {
		for _, curso := range est.Cursos {
			cursosMap[curso.Nombre] = true
		}
	}

	var cursos []string
	for curso := range cursosMap {
		cursos = append(cursos, curso)
	}
	return cursos
}

// Estructura auxiliar para almacenar estudiantes y sus notas
type EstudianteNota struct {
	Estudiante models.Estudiante
	Nota       float64
}

// Obtener los estudiantes que toman un curso específico
func ObtenerEstudiantesPorCurso(estudiantes []models.Estudiante, nombreCurso string) []EstudianteNota {
	var estudiantesCurso []EstudianteNota

	for _, est := range estudiantes {
		for _, c := range est.Cursos {
			if c.Nombre == nombreCurso {
				estudiantesCurso = append(estudiantesCurso, EstudianteNota{Estudiante: est, Nota: c.Nota})
			}
		}
	}

	return estudiantesCurso
}

// Estructura para almacenar estadísticas de un curso
type EstadisticasCurso struct {
	Curso              string
	Promedio           float64
	Rango              float64
	Varianza           float64
	DesviacionEstandar float64
}

// Calcular estadísticas para cada curso
func EstadisticasPorCurso(estudiantes []models.Estudiante) []EstadisticasCurso {
	cursos := ObtenerCursos(estudiantes)
	var estadisticas []EstadisticasCurso

	for _, curso := range cursos {
		notas := []float64{}

		// Recolectar notas del curso
		for _, est := range estudiantes {
			for _, c := range est.Cursos {
				if c.Nombre == curso {
					notas = append(notas, c.Nota)
				}
			}
		}

		if len(notas) == 0 {
			continue
		}

		// Calcular promedio
		suma := 0.0
		for _, nota := range notas {
			suma += nota
		}
		promedio := suma / float64(len(notas))

		// Calcular rango
		minNota := notas[0]
		maxNota := notas[0]
		for _, nota := range notas {
			if nota < minNota {
				minNota = nota
			}
			if nota > maxNota {
				maxNota = nota
			}
		}
		rango := maxNota - minNota

		// Calcular varianza
		sumaVarianza := 0.0
		for _, nota := range notas {
			sumaVarianza += (nota - promedio) * (nota - promedio)
		}
		varianza := sumaVarianza / float64(len(notas))

		// Calcular desviación estándar
		desviacionEstandar := math.Sqrt(varianza)

		estadisticas = append(estadisticas, EstadisticasCurso{
			Curso:              curso,
			Promedio:           promedio,
			Rango:              rango,
			Varianza:           varianza,
			DesviacionEstandar: desviacionEstandar,
		})
	}

	return estadisticas
}

// Calcular promedio de estudiantes por rango de edad
func PromedioPorRangoEdad(estudiantes []models.Estudiante) map[string]float64 {
	rangos := []struct {
		Min int
		Max int
	}{
		{20, 29},
		{30, 39},
		{40, 49},
	}

	promedioPorRango := make(map[string]float64)

	for _, rango := range rangos {
		sumaPromedios := 0.0
		contador := 0
		for _, est := range estudiantes {
			if est.Edad >= rango.Min && est.Edad <= rango.Max {
				promedio := CalcularPromedioEstudiante(est)
				sumaPromedios += promedio
				contador++
			}
		}
		if contador > 0 {
			rangoStr := fmt.Sprintf("%d-%d", rango.Min, rango.Max)
			promedioPorRango[rangoStr] = sumaPromedios / float64(contador)
		}
	}

	return promedioPorRango
}
