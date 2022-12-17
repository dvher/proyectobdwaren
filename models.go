package main

import (
	"fmt"
	"time"
)

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

func (u *Usuario) String() string {
	return fmt.Sprintf(
		"%s %s:\nID: %d\nRUT: %d-%s\nIngreso Social: %d\nCorreo: %s\nFecha Creación: %v\nEstado: %t\n",
		u.Nombre,
		u.Apellido,
		u.Id,
		u.Rut,
		u.DV,
		u.Ingreso_social,
		u.Correo,
		u.Fecha_creacion.Format("2006-01-02"),
		u.Estado,
	)
}

func (d *Dispositivo) String() string {
	return fmt.Sprintf("%s %s\nID: %d\nPrecio: $%d\n", d.Tipo, d.Marca, d.Id, d.Precio)
}
