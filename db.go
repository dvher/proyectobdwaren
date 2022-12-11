package main

import (
	"crypto/sha1"
	"database/sql"
	"fmt"
	"io"
	"log"

	_ "github.com/lib/pq"
)

var db *sql.DB

func initDB() {
	connStr := "postgresql://proyecto_user:proyecto_pass@localhost/proyectodb?sslmode=disable"

	var err error

	db, err = sql.Open("postgres", connStr)

	if err != nil {
		log.Println(err)
		return
	}
}

func closeDB() error {
	return db.Close()
}

func crearUsuario(user *Usuario) {
	pass := sha1.New()
	io.WriteString(pass, user.Contrasena)
	res, err := db.Exec("INSERT INTO Usuario(ingreso_social, correo, nombre, apellido, rut, contrasena, fecha_creacion, dv, estado) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9);", user.Ingreso_social, user.Correo, user.Nombre, user.Apellido, user.Rut, pass, user.Fecha_creacion, user.DV, user.Estado)

	if err != nil {
		log.Println(err)
	}

	user_id, err := res.LastInsertId()

	if err != nil {
		log.Println(err)
	}

	fmt.Println(user_id)
}

func agregarDispositivo(disp *Dispositivo) {

}
