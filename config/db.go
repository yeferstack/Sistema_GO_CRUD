package config

import (
	"database/sql" // para conexiones con sql
	"fmt"
	"log" // imprimir y manejar logs
	"os"  // para leer variables de entorno

	_ "github.com/lib/pq" // driver postgreSQL
)

var DB *sql.DB // instancia global de la base de datos

// getEnv obtiene una variable de entorno o retorna un valor por defecto
func getEnv(key, defaultVal string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultVal
}

// ConnectDB establece conexion con postgreSQL usando variables de entorno
func ConnectDB() {
	// variables de entorno
	host     := getEnv("INSCRIPCION_CRUD_PGHOST", "localhost")
	port     := getEnv("INSCRIPCION_CRUD_PGPORT", "5432")
	user     := getEnv("INSCRIPCION_CRUD_PGUSER", "postgres")
	password := getEnv("INSCRIPCION_CRUD_PGPASS", "0219251007")
	dbname   := getEnv("INSCRIPCION_CRUD_PGDB", "xchango_db")
	schema   := getEnv("INSCRIPCION_CRUD_PGSCHEMA", "Sistema")

	// CADENA DE CONEXION
	psqlInfo := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s search_path=%s sslmode=disable",
		host, port, user, password, dbname, schema,
	)

	// abrir conexion db
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal("error al conectar:", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("no se puede conectar:", err)
	}

	fmt.Println("conexion a base de datos exitosa")
	fmt.Println("Conectado a la db:", dbname, "y esquema:", schema)
	DB = db // ASIGNAR A CONEXION GLOBAL
}
