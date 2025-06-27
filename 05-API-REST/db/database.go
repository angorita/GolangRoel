package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

//username:password@tcp(localhost:3036)/database
const url = "oscar:emi@tcp(localhost:3306)/loft"

// guarda la conexion y creo la variable db que es global y trae a *sql.DB
var db *sql.DB

// realiza la conexion
func Connect() {
	conection, err := sql.Open("mysql", url)
	if err != nil {
		panic(err)
	}
	fmt.Println("Conexion exitosa")
	db = conection

}

// cierra la conexion
func Close() {
	db.Close()
}

// verificar conexion
func Ping() {
	if err := db.Ping(); err != nil {
		panic(err)
	}
}

// crea una tabla material, con su nombre, chequea si existe..y la crea o no
func CreateTable(schema, name string) {
	if !ExistsTable(name) {
		_, err := Exec(schema)
		if err != nil {
			fmt.Println(err)
		}
	}

}

// reiniciar el registro de una tabla
func TruncateTable(tablename string) {
	sql := fmt.Sprintf("Truncate %s", tablename)
	Exec(sql)
}

// **********************************************************************
// polimorfismo de Exec
// **********************************************************************
func Exec(query string, args ...interface{}) (sql.Result, error) {
	Connect()
	result, err := db.Exec(query, args...)
	Close()
	if err != nil {
		fmt.Println(err)
	}
	return result, err
}

// polimorfismo de Query
func Query(query string, args ...interface{}) (*sql.Rows, error) {
	Connect()
	rows, err := db.Query(query, args...)
	if err != nil {
		fmt.Println(err)
	}
	return rows, err
}

// **********************************************************************

// verifica si existe una tabla,con fmt.Sprintf devuelve un string
func ExistsTable(tablename string) bool {
	sql := fmt.Sprintf("Show tables like '%s'", tablename)
	rows, err := Query(sql)
	if err != nil {
		fmt.Println("Error", err)
	}
	return rows.Next() //devuelve true si la puede recorrer false si no ...
}
