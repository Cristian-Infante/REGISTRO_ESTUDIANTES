package routes

import (
	"fmt"
	"github.com/Cristian-Infante/REGISTRO_ESTUDIANTES/controllers"
	"github.com/Cristian-Infante/REGISTRO_ESTUDIANTES/repositories"
	"github.com/Cristian-Infante/REGISTRO_ESTUDIANTES/services"
	"net/http"
)

func LoadRoutes() {
	// Crear repositorio, servicio y controlador
	repo := repositories.NewEstudianteRepository("data/generated.json")
	fmt.Print(repo)
	servicio := services.NewEstudianteService(repo)
	controller := controllers.NewEstudianteController(servicio)

	// Rutas
	http.HandleFunc("/", controller.MostrarMenu)
	http.HandleFunc("/mejor_estudiante", controller.MostrarMejorEstudiante)
	http.HandleFunc("/peor_estudiante", controller.MostrarPeorEstudiante)
	http.HandleFunc("/mejores", controller.MostrarMejoresEstudiantes)
	http.HandleFunc("/peores", controller.MostrarPeoresEstudiantes)
	http.HandleFunc("/mayor_masculino", controller.MostrarEstudianteMasculinoMayorEdad)
	http.HandleFunc("/mayor_femenino", controller.MostrarEstudianteFemeninoMayorEdad)
	http.HandleFunc("/estadisticas_curso", controller.MostrarEstadisticasPorCurso)
	http.HandleFunc("/matriculados_anio", controller.MostrarMatriculadosEnAnio)
	http.HandleFunc("/rango_edad", controller.MostrarPromedioPorRangoEdad)
}
