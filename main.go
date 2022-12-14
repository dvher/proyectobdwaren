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

	fmt.Println(
		"Bienvenido al sistema de gestión de dispositivos, qué desea hacer?",
	)

MAIN_CYCLE:
	for {
		PrintMenu(
			"",
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
			break MAIN_CYCLE
		default:
			fmt.Println("Opción inválida")
		}

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
		"Asignar dispositivo a usuario",
	)

	input := GetInput[int]("Ingrese una opción")

	switch input {
	case 1:
		AgregarUsuario()
	case 2:
		AddDispositivo()
	case 3:
		AddUsuarioDispositivo()
	default:
		fmt.Println("Opción inválida")
	}
}

func Modificar() {

	PrintMenu(
		"Qué desea modificar?",
		"Nombre usuario",
		"Apellido usuario",
		"Correo usuario",
		"Contraseña usuario",
	)

	input := GetInput[int]("Ingrese una opción")

	switch input {
	case 1:
		ModificarNombre()
	case 2:
		ModificarApellido()
	case 3:
		ModificarCorreo()
	case 4:
		ModificarContrasena()
	default:
		fmt.Println("Opción inválida")
	}
}

func Eliminar() {
	PrintMenu(
		"Qué desea eliminar?",
		"Usuario",
		"Dispositivo",
		"Quitar dispositivo a un usuario",
	)

	input := GetInput[int]("Ingrese una opción")

	switch input {
	case 1:
		EliminarUsuario()
	case 2:
		EliminarDispositivo()
	case 3:
		QuitarDispositivoUsuario()
	default:
		fmt.Println("Opción inválida")
	}
}

func Listar() {
	PrintMenu(
		"Qué desea listar?",
		"Usuarios",
		"Dispositivos",
		"Usuarios y sus dispositivos",
		"Dispositivos de un usuario",
	)

	input := GetInput[int]("Ingrese una opción")

	switch input {
	case 1:
		ListarUsuarios()
	case 2:
		ListarDispositivos()
	case 3:
		ListarUsuariosDispositivos()
	case 4:
		ListarDispositivosUsuario()
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

func AddUsuarioDispositivo() {
	fmt.Println("Ingrese los datos del usuario")
	id_usuario := GetInput[int]("ID del usuario")
	id_dispositivo := GetInput[int]("ID del dispositivo")

	ud := UsuarioDispositivo{
		Id_usuario:     id_usuario,
		Id_dispositivo: id_dispositivo,
	}

	err := AgregarUsuarioDispositivo(&ud)

	if err == nil {
		fmt.Println("Relación creada correctamente")
	}
}

func ModificarContrasena() {

	id := GetInput[int]("Ingrese el ID del usuario")

	passwd := GetInput[string]("Ingrese la nueva contraseña")

	CambiarContrasena(id, passwd)

	fmt.Println("Contraseña cambiada correctamente")

}

func ModificarNombre() {

	id := GetInput[int]("Ingrese el ID del usuario")

	nombre := GetInput[string]("Ingrese el nuevo nombre")

	CambiarNombre(id, nombre)

	fmt.Println("Nombre cambiado correctamente")

}

func ModificarApellido() {

	id := GetInput[int]("Ingrese el ID del usuario")

	apellido := GetInput[string]("Ingrese el nuevo apellido")

	CambiarApellido(id, apellido)

	fmt.Println("Apellido cambiado correctamente")

}

func ModificarCorreo() {

	id := GetInput[int]("Ingrese el ID del usuario")

	correo := GetInput[string]("Ingrese el nuevo correo")

	CambiarCorreo(id, correo)

	fmt.Println("Correo cambiado correctamente")

}

func ListarDispositivosUsuario() {

	id_usuario := GetInput[int]("ID del usuario")

	ListarDispositivosDeUsuario(id_usuario)

}

func EliminarUsuario() {

	id_usuario := GetInput[int]("ID del usuario")

	err := DeleteUser(id_usuario)

	if err == nil {
		fmt.Println("Usuario eliminado correctamente")
	}
}

func EliminarDispositivo() {

	id_dispositivo := GetInput[int]("ID del dispositivo")

	err := DeleteDispositivo(id_dispositivo)

	if err == nil {
		fmt.Println("Dispositivo eliminado correctamente")
	}
}

func QuitarDispositivoUsuario() {

	id_usuario := GetInput[int]("ID del usuario")
	id_dispositivo := GetInput[int]("ID del dispositivo")

	err := DeleteUsuarioDispositivo(id_usuario, id_dispositivo)

	if err == nil {
		fmt.Println("Relación eliminada correctamente")
	}
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
	fmt.Println()
	return input
}
