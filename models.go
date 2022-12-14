package main

import "time"

type Usuario struct {
	Id             int
	Ingreso_social int
	Correo         string
	Nombre         string
	Apellido       string
	Rut            int
	Contrasena     string
	Fecha_creacion time.Time
	DV             string
	Estado         bool
}

type Dispositivo struct {
	Id     int
	Tipo   string
	Marca  string
	Precio int
}

type UsuarioDispositivo struct {
	Id             int
	Id_usuario     int
	Id_dispositivo int
}
