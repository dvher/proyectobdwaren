CREATE TABLE IF NOT EXISTS Usuario (
    id SERIAL,
    ingreso_social INT NOT NULL,
    correo VARCHAR(20) NOT NULL UNIQUE,
    nombre VARCHAR(20) NOT NULL,
    apellido VARCHAR(20) NOT NULL,
    rut INT NOT NULL UNIQUE,
    contrasena VARCHAR(88) NOT NULL,
    fecha_creacion DATE NOT NULL,
    DV VARCHAR(1) NOT NULL UNIQUE,
    estado BOOLEAN NOT NULL,
    PRIMARY KEY(id)
);

CREATE TABLE IF NOT EXISTS Dispositivo (
    id SERIAL,
    tipo VARCHAR(15) NOT NULL,
    marca VARCHAR(15) NOT NULL,
    precio INT NOT NULL,
    PRIMARY KEY(id)
);

CREATE TABLE IF NOT EXISTS UsuarioDispositivo (
    id SERIAL PRIMARY KEY,
    id_usuario INT NOT NULL,
    id_dispositivo INT NOT NULL,
    FOREIGN KEY(id_usuario) REFERENCES Usuario(id),
    FOREIGN KEY(id_dispositivo) REFERENCES Dispositivo(id)
);

CREATE OR REPLACE FUNCTION cantidad_dispositivos_usuario(id_user INT) RETURNS INT AS $$
DECLARE
    cantidad INT;
BEGIN
    SELECT COUNT(*) INTO cantidad FROM UsuarioDispositivo WHERE UsuarioDispositivo.id_usuario = id_user;
    RETURN cantidad;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION dispositivos_libres() RETURNS TABLE(id INT, tipo VARCHAR(15), marca VARCHAR(15), precio INT) AS $$
BEGIN
    RETURN QUERY SELECT * FROM Dispositivo WHERE Dispositivo.id NOT IN (SELECT id_dispositivo FROM UsuarioDispositivo);
END;
$$ LANGUAGE plpgsql;
