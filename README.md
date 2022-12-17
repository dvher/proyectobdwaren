Para el correcto funcionamiento de la aplicación es necesario crear una base de datos llamada *proyectodb* en PostgreSQL y su usuario correspondiente, para esto se debe ejecutar el siguiente comando en la consola de PostgreSQL:

```sql
CREATE DATABASE proyectodb;
CREATE USER proyecto_user WITH PASSWORD 'proyecto_pass';
GRANT ALL PRIVILEGES ON DATABASE proyectodb TO proyecto_user;
```

Luego se deben ejecutar los comandos del archivo db.sql para crear las tablas y sus relaciones.