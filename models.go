package main

import "time"

type Usuario struct {
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
	Tipo_dispositivo string
	Marca            string
	Precio           int
}

type UsuarioDispositivo struct {
	id_usuario     int
	id_dispositivo int
}
