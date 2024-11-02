package controllers

import (
	"fmt"
	"github.com/Cristian-Infante/REGISTRO_ESTUDIANTES/models"
	"github.com/Cristian-Infante/REGISTRO_ESTUDIANTES/services"
	"html/template"
	"net/http"
	"strconv"
)

type EstudianteController struct {
	Servicio *services.EstudianteService
}

func NewEstudianteController(servicio *services.EstudianteService) *EstudianteController {
	return &EstudianteController{Servicio: servicio}
}

// Ruta principal con el menú
func (ec *EstudianteController) MostrarMenu(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html", "templates/layout.html"))
	err := tmpl.ExecuteTemplate(w, "layout", nil)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
}

// Mostrar el mejor estudiante por promedio
func (ec *EstudianteController) MostrarMejorEstudiante(w http.ResponseWriter, r *http.Request) {
	mejorEstudiante, promedio := ec.Servicio.EstudianteMejorPromedio()
	data := struct {
		Nombre   string
		Apellido string
		Promedio float64
	}{
		Nombre:   mejorEstudiante.Nombre,
		Apellido: mejorEstudiante.Apellido,
		Promedio: promedio,
	}
	tmpl := template.Must(template.ParseFiles("templates/mejor_estudiante.html", "templates/layout.html"))
	err := tmpl.ExecuteTemplate(w, "layout", data)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
}

// Mostrar el peor estudiante por promedio
func (ec *EstudianteController) MostrarPeorEstudiante(w http.ResponseWriter, r *http.Request) {
	peorEstudiante, promedio := ec.Servicio.EstudiantePeorPromedio()
	data := struct {
		Nombre   string
		Apellido string
		Promedio float64
	}{
		Nombre:   peorEstudiante.Nombre,
		Apellido: peorEstudiante.Apellido,
		Promedio: promedio,
	}
	tmpl := template.Must(template.ParseFiles("templates/peor_estudiante.html", "templates/layout.html"))
	err := tmpl.ExecuteTemplate(w, "layout", data)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
}

// Mostrar los mejores estudiantes por curso
func (ec *EstudianteController) MostrarMejoresEstudiantes(w http.ResponseWriter, r *http.Request) {
	mejoresEstudiantes := ec.Servicio.Top10MejoresPorCurso()

	// Crear una estructura que incluya el título y los estudiantes con sus promedios
	data := struct {
		Cursos map[string][]struct {
			Nombre   string
			Apellido string
			Promedio float64
		}
	}{
		Cursos: make(map[string][]struct {
			Nombre   string
			Apellido string
			Promedio float64
		}),
	}

	// Añadir los estudiantes con su promedio a la estructura
	for curso, estudiantes := range mejoresEstudiantes {
		for _, est := range estudiantes {
			promedio := ec.Servicio.CalcularPromedioDeEstudianteEnCurso(est, curso)
			data.Cursos[curso] = append(data.Cursos[curso], struct {
				Nombre   string
				Apellido string
				Promedio float64
			}{
				Nombre:   est.Nombre,
				Apellido: est.Apellido,
				Promedio: promedio,
			})
		}
	}

	tmpl := template.Must(template.ParseFiles("templates/mejores_estudiantes.html", "templates/layout.html"))
	err := tmpl.ExecuteTemplate(w, "layout", data)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
}

// Mostrar los peores estudiantes por curso
func (ec *EstudianteController) MostrarPeoresEstudiantes(w http.ResponseWriter, r *http.Request) {
	peoresEstudiantes := ec.Servicio.Top10PeoresPorCurso()

	// Crear una estructura que incluya el título y los estudiantes con sus promedios
	data := struct {
		Cursos map[string][]struct {
			Nombre   string
			Apellido string
			Promedio float64
		}
	}{
		Cursos: make(map[string][]struct {
			Nombre   string
			Apellido string
			Promedio float64
		}),
	}

	// Añadir los estudiantes con su promedio a la estructura
	for curso, estudiantes := range peoresEstudiantes {
		for _, est := range estudiantes {
			promedio := ec.Servicio.CalcularPromedioDeEstudianteEnCurso(est, curso)
			data.Cursos[curso] = append(data.Cursos[curso], struct {
				Nombre   string
				Apellido string
				Promedio float64
			}{
				Nombre:   est.Nombre,
				Apellido: est.Apellido,
				Promedio: promedio,
			})
		}
	}

	tmpl := template.Must(template.ParseFiles("templates/peores_estudiantes.html", "templates/layout.html"))
	err := tmpl.ExecuteTemplate(w, "layout", data)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
}

// Mostrar el estudiante masculino de mayor edad
func (ec *EstudianteController) MostrarEstudianteMasculinoMayorEdad(w http.ResponseWriter, r *http.Request) {
	estudiante := ec.Servicio.EstudianteMasculinoMayorEdad()
	tmpl := template.Must(template.ParseFiles("templates/mayor_masculino.html", "templates/layout.html"))
	err := tmpl.ExecuteTemplate(w, "layout", estudiante)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
}

// Mostrar la estudiante femenina de mayor edad
func (ec *EstudianteController) MostrarEstudianteFemeninoMayorEdad(w http.ResponseWriter, r *http.Request) {
	estudiante := ec.Servicio.EstudianteFemeninoMayorEdad()
	tmpl := template.Must(template.ParseFiles("templates/mayor_femenino.html", "templates/layout.html"))
	err := tmpl.ExecuteTemplate(w, "layout", estudiante)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
}

// Estadísticas por curso
func (ec *EstudianteController) MostrarEstadisticasPorCurso(w http.ResponseWriter, r *http.Request) {
	estadisticas := ec.Servicio.EstadisticasPorCurso()
	tmpl := template.Must(template.ParseFiles("templates/estadisticas_curso.html", "templates/layout.html"))
	err := tmpl.ExecuteTemplate(w, "layout", estadisticas)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
}

// Estudiantes matriculados en un año específico
func (ec *EstudianteController) MostrarMatriculadosEnAnio(w http.ResponseWriter, r *http.Request) {
	anio, _ := strconv.Atoi(r.URL.Query().Get("anio"))
	estudiantes := ec.Servicio.EstudiantesMatriculadosEnAnio(anio)

	// Crear una estructura que incluya el año y los estudiantes
	data := struct {
		Anio        int
		Estudiantes []models.Estudiante
	}{
		Anio:        anio,
		Estudiantes: estudiantes,
	}

	tmpl := template.Must(template.ParseFiles("templates/matriculados_anio.html", "templates/layout.html"))
	err := tmpl.ExecuteTemplate(w, "layout", data)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
}

// Promedio por rango de edad
func (ec *EstudianteController) MostrarPromedioPorRangoEdad(w http.ResponseWriter, r *http.Request) {
	promedioPorRango := ec.Servicio.PromedioPorRangoEdad()
	tmpl := template.Must(template.ParseFiles("templates/rango_edad.html", "templates/layout.html"))
	err := tmpl.ExecuteTemplate(w, "layout", promedioPorRango)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
}
