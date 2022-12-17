package main

import (
	"crypto/sha512"
	"database/sql"
	"encoding/base64"
	"errors"
	"fmt"
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
	_, err := db.Exec(
		"INSERT INTO Usuario(ingreso_social, correo, nombre, apellido, rut, contrasena, fecha_creacion, dv, estado) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9);",
		user.Ingreso_social,
		user.Correo,
		user.Nombre,
		user.Apellido,
		user.Rut,
		pwd,
		time.Now(),
		user.DV,
		user.Estado,
	)

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
	_, err := db.Exec(
		"INSERT INTO UsuarioDispositivo(id_usuario, id_dispositivo) VALUES($1, $2);",
		ud.Id_usuario,
		ud.Id_dispositivo,
	)

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

func ListarUsuarios() (err error) {

	var (
		u    Usuario
		usrs []Usuario
	)

	stmt, err := db.Prepare("SELECT * FROM Usuario;")

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
		err = rows.Scan(
			&u.Id,
			&u.Ingreso_social,
			&u.Correo,
			&u.Nombre,
			&u.Apellido,
			&u.Rut,
			&u.Contrasena,
			&u.Fecha_creacion,
			&u.DV,
			&u.Estado,
		)

		if err != nil {
			return
		}
		usrs = append(usrs, u)
	}

	for _, v := range usrs {
		fmt.Println(v.String())
	}

	return
}

func ListarDispositivos() (err error) {

	var (
		d     Dispositivo
		disps []Dispositivo
	)

	stmt, err := db.Prepare("SELECT * FROM Dispositivo;")

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

	for _, v := range disps {
		fmt.Println(v.String())
	}

	return
}

func ListarDispositivosDeUsuario(id int) (err error) {

	var (
		nombre   string
		apellido string
		rut      int
		dv       string
		d        Dispositivo
		disps    []Dispositivo
	)

	stmt, err := db.Prepare("SELECT nombre, apellido, rut, dv FROM usuario WHERE id = $1;")

	if err != nil {
		log.Println(err)
		return
	}

	rows, err := stmt.Query(id)

	if err != nil {
		log.Println(err)
		return
	}

	if !rows.Next() {
		return errors.New("Usuario no encontrado")
	}

	err = rows.Scan(&nombre, &apellido, &rut, &dv)

	if err != nil {
		return
	}

	fmt.Printf("Usuario: %s %s de RUT %d-%s tiene:\n\n", nombre, apellido, rut, dv)

	stmt, err = db.Prepare("SELECT id_dispositivo FROM UsuarioDispositivo WHERE id_usuario = $1;")

	if err != nil {
		log.Println(err)
		return
	}

	rows, err = stmt.Query(id)

	if err != nil {
		log.Println(err)
		return
	}

	for rows.Next() {
		err = rows.Scan(&d.Id)

		if err != nil {
			log.Println(err)
			return
		}

		stmt, err = db.Prepare("SELECT * FROM Dispositivo WHERE id = $1;")

		if err != nil {
			log.Println(err)
			return
		}

		disp_rows, err := stmt.Query(d.Id)

		if err != nil {
			log.Println(err)
			return err
		}

		if !disp_rows.Next() {
			return errors.New("Dispositivo no encontrado")
		}

		err = disp_rows.Scan(&d.Id, &d.Tipo, &d.Marca, &d.Precio)

		if err != nil {
			log.Println(err)
			return err
		}

		disps = append(disps, d)

	}

	for _, v := range disps {
		fmt.Println(v.String())
	}

	return

}

func ListarUsuariosDispositivos() (err error) {

	var (
		id int
	)

	stmt, err := db.Prepare(
		"SELECT id FROM Usuario WHERE (SELECT COUNT(*) FROM UsuarioDispositivo WHERE id_usuario = Usuario.id) > 0;",
	)

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
		err = rows.Scan(&id)

		if err != nil {
			log.Println(err)
			return
		}

		err = ListarDispositivosDeUsuario(id)

		if err != nil {
			log.Println(err)
			return
		}
	}

	return
}

func DeleteUser(id int) error {

	_, err := db.Exec("DELETE FROM Usuario WHERE id = $1;", id)

	if err != nil {
		log.Println(err)
		return err
	}

	_, err = db.Exec("DELETE FROM UsuarioDispositivo WHERE id_usuario = $1;", id)

	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func DeleteDispositivo(id int) error {

	_, err := db.Exec("DELETE FROM Dispositivo WHERE id = $1;", id)

	if err != nil {
		log.Println(err)
		return err
	}

	_, err = db.Exec("DELETE FROM UsuarioDispositivo WHERE id_dispositivo = $1;", id)

	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func DeleteUsuarioDispositivo(id_usr, id_dev int) error {

	_, err := db.Exec("DELETE FROM UsuarioDispositivo WHERE id_usuario = $1 AND id_dispositivo = $2;", id_usr, id_dev)

	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
