package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	initDB()
	crearUsuario(&Usuario{
		Ingreso_social: 1,
		Correo:         "a@a.cl",
		Nombre:         "Diego",
		Apellido:       "Vergara",
		Rut:            20383812,
		Contrasena:     "Dvher2510%.",
		Fecha_creacion: time.Now(),
		DV:             "3",
		Estado:         true,
	})
	fmt.Println("Hola mundo!")
	err := closeDB()

	if err != nil {
		log.Println(err)
	}
}
