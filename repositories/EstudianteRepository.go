package repositories

import (
	"encoding/json"
	"fmt"
	"github.com/Cristian-Infante/REGISTRO_ESTUDIANTES/interfaces"
	"github.com/Cristian-Infante/REGISTRO_ESTUDIANTES/models"
	"io/ioutil"
)

type EstudianteRepository struct {
	Archivo string
}

func NewEstudianteRepository(archivo string) interfaces.IEstudianteRepository {
	fmt.Println("Ruta del archivo JSON: ", archivo)
	return &EstudianteRepository{Archivo: archivo}
}

func (er *EstudianteRepository) GetAll() ([]models.Estudiante, error) {
	// Leer el archivo JSON
	datosJSON, err := ioutil.ReadFile(er.Archivo)
	if err != nil {
		// Imprimir el error y retornar
		fmt.Printf("Error al leer el archivo JSON: %v\n", err)
		return nil, err
	}

	// Intentar parsear el JSON
	var estudiantes []models.Estudiante
	err = json.Unmarshal(datosJSON, &estudiantes)
	if err != nil {
		// Imprimir el error si no se puede parsear
		fmt.Printf("Error al parsear el archivo JSON: %v\n", err)
		return nil, err
	}

	// Retornar los estudiantes leídos
	return estudiantes, nil
}
