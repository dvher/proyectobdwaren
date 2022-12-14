package main

import (
	"fmt"
	"log"
)

func main() {
	err := InitDB()

	if err != nil {
		return
	}

	PrintMenu(
		"Bienvenido al sistema de gestión de dispositivos, qué desea hacer?",
		"Agregar",
		"Modificar",
		"Eliminar",
		"Listar",
		"Salir",
	)

	input := GetInput[int]("Ingrese una opción")

	switch input {
	case 1:
		Agregar()
	case 2:
		Modificar()
	case 3:
		Eliminar()
	case 4:
		Listar()
	case 5:
		fmt.Println("Saliendo...")
	default:
		fmt.Println("Opción inválida")
	}

	err = CloseDB()

	if err != nil {
		log.Println(err)
	}
}

func Agregar() {
	PrintMenu(
		"Qué desea agregar?",
		"Usuario",
		"Dispositivo",
	)

	input := GetInput[int]("Ingrese una opción")

	switch input {
	case 1:
		AgregarUsuario()
	case 2:
		AddDispositivo()
	default:
		fmt.Println("Opción inválida")
	}
}

func Modificar() {
	PrintMenu(
		"Qué desea modificar?",
		"Usuario",
		"Dispositivo",
	)

	input := GetInput[int]("Ingrese una opción")

	switch input {
	case 1:
		ModificarUsuario()
	case 2:
		ModificarDispositivo()
	default:
		fmt.Println("Opción inválida")
	}
}

func Eliminar() {
	PrintMenu(
		"Qué desea eliminar?",
		"Usuario",
		"Dispositivo",
	)

	input := GetInput[int]("Ingrese una opción")

	switch input {
	case 1:
		EliminarUsuario()
	case 2:
		EliminarDispositivo()
	default:
		fmt.Println("Opción inválida")
	}
}

func Listar() {
	PrintMenu(
		"Qué desea listar?",
		"Usuarios",
		"Dispositivos",
	)

	input := GetInput[int]("Ingrese una opción")

	switch input {
	case 1:
		ListarUsuarios()
	case 2:
		ListarDispositivos()
	default:
		fmt.Println("Opción inválida")
	}
}

func AgregarUsuario() {

	var (
		rut int
		dv  string
	)

	fmt.Println("Ingrese los datos del usuario")
	ingreso_social := GetInput[int]("Ingreso social")
	correo := GetInput[string]("Correo")
	nombre := GetInput[string]("Nombre")
	apellido := GetInput[string]("Apellido")
	fmt.Print("Rut: ")
	fmt.Scanf("%d-%s", &rut, &dv)
	if dv == "" {
		dv = GetInput[string]("Dígito verificador")
	}
	contrasena := GetInput[string]("Contraseña")

	user := Usuario{
		Ingreso_social: ingreso_social,
		Correo:         correo,
		Nombre:         nombre,
		Apellido:       apellido,
		Rut:            rut,
		Contrasena:     contrasena,
		DV:             dv,
		Estado:         false,
	}

	err := CrearUsuario(&user)

	if err == nil {
		fmt.Println("Usuario creado correctamente")
	}

}

func AddDispositivo() {
	fmt.Println("Ingrese los datos del dispositivo")
	tipo := GetInput[string]("Tipo")
	marca := GetInput[string]("Marca")
	precio := GetInput[int]("Precio")

	disp := Dispositivo{
		Tipo:   tipo,
		Marca:  marca,
		Precio: precio,
	}

	err := AgregarDispositivo(&disp)

	if err == nil {
		fmt.Println("Dispositivo creado correctamente")
	}
}

func ModificarUsuario() {

	var (
		rut int
		dv  string
	)

	fmt.Print("Ingrese el rut del usuario: ")
	fmt.Scanf("%d-%s", &rut, &dv)

}

func ModificarDispositivo() {
}

func EliminarUsuario() {
}

func EliminarDispositivo() {
}

func ListarUsuarios() {
}

func ListarDispositivos() {
}

func PrintMenu(title string, options ...string) {
	fmt.Println(title)
	for i, option := range options {
		fmt.Printf("%d. %s\n", i+1, option)
	}
	fmt.Println()
}

func GetInput[T any](input_message string) T {
	var input T
	fmt.Printf("%s: ", input_message)
	fmt.Scanln(&input)
	return input
}
