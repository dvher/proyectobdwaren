package main

import (
	"crypto/sha512"
	"database/sql"
	"encoding/base64"
	"io"
	"log"
	"time"

	_ "github.com/lib/pq"
)

var db *sql.DB

func InitDB() error {
	connStr := "postgresql://proyecto_user:proyecto_pass@localhost/proyectodb?sslmode=disable"

	var err error

	db, err = sql.Open("postgres", connStr)

	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func CloseDB() error {
	return db.Close()
}

func CrearUsuario(user *Usuario) error {
	pass := sha512.New()
	io.WriteString(pass, user.Contrasena)
	pwd := base64.StdEncoding.EncodeToString(pass.Sum(nil))
	_, err := db.Exec("INSERT INTO Usuario(ingreso_social, correo, nombre, apellido, rut, contrasena, fecha_creacion, dv, estado) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9);", user.Ingreso_social, user.Correo, user.Nombre, user.Apellido, user.Rut, pwd, time.Now(), user.DV, user.Estado)

	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func AgregarDispositivo(disp *Dispositivo) error {
	_, err := db.Exec("INSERT INTO Dispositivo(tipo, marca, precio) VALUES ($1, $2, $3);", disp.Tipo, disp.Marca, disp.Precio)

	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func AgregarUsuarioDispositivo(ud *UsuarioDispositivo) error {
	_, err := db.Exec("INSERT INTO UsuarioDispositivo(id_usuario, id_dispositivo) VALUES($1, $2);", ud.Id_usuario, ud.Id_dispositivo)

	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func BorrarUsuario(id int) error {

	_, err := db.Exec("DELETE FROM UsuarioDispositivo WHERE id_usuario = $1;", id)

	if err != nil {
		log.Println(err)
		return err
	}

	_, err = db.Exec("DELETE FROM Usuario WHERE id = $1;", id)

	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func CantidadDispositivosUsuario(id int) (int, error) {

	var cantidad int

	stmt, err := db.Prepare("SELECT cantidad_dispositivos_usuario($1);")

	if err != nil {
		log.Println(err)
		return 0, err
	}

	rows, err := stmt.Query(id)

	if err != nil {
		log.Println(err)
		return 0, err
	}

	rows.Next()
	err = rows.Scan(&cantidad)

	if err != nil {
		log.Println(err)
		return 0, err
	}

	err = rows.Close()

	if err != nil {
		log.Println(err)
		return 0, err
	}

	return cantidad, nil
}

func DispositivosLibres() (disps []Dispositivo, err error) {

	var (
		d Dispositivo
	)

	stmt, err := db.Prepare("SELECT * FROM dispositivos_libres();")

	if err != nil {
		log.Println(err)
		return
	}

	rows, err := stmt.Query()

	if err != nil {
		log.Println(err)
		return
	}

	for rows.Next() {
		err = rows.Scan(&d.Id, &d.Tipo, &d.Marca, &d.Precio)

		if err != nil {
			return
		}
		disps = append(disps, d)
	}

	return
}
