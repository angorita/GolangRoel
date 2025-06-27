package models

import (
	"fmt"
	"web2/db"
)

type Material struct {
	Descripcion string
	Precio      float64
	Cantidad    int
	Fecha       string
	Dolar       float64
	Id          int
	Bhabilitado bool
}
type Materiales []Material

const MaterialSchema = `
CREATE TABLE material (
  descripcion varchar(100) DEFAULT NULL,
  precio decimal(15,2) NOT NULL,
  cantidad int(11) DEFAULT NULL,
  fecha date DEFAULT NULL,
  dolar decimal(6,2) NOT NULL,
  id bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  bhabilitado tinyint(1) DEFAULT NULL,
  PRIMARY KEY (id),
  UNIQUE KEY id (id)
) 
`

// *********** id solo en busqueda y lista **********************//
// construir material no incluir el id van 6 campos
func NewMaterial(descripcion string, precio float64, cantidad int, fecha string, dolar int, bhabilitado bool) *Material {
	material := &Material{Descripcion: descripcion, Precio: precio, Cantidad: cantidad, Fecha: fecha, Dolar: float64(dolar), Bhabilitado: bhabilitado}
	return material
}

// crear material e insertar no cincluir el id van 6 campos
func CreateMaterial(descripcion string, precio float64, cantidad int, fecha string, dolar int, bhabilitado bool) *Material {
	material := NewMaterial(descripcion, precio, cantidad, fecha, dolar, bhabilitado)
	material.Save()
	return material
}

// insertar registro no incluir el id van 6 campos
func (material *Material) Insert() {
	sql := `insert material descripcion=?,precio=?,cantidad=?,fecha=?,dolar=?,bhabilitado=?`
	db.Exec(sql, material.Descripcion, material.Precio, material.Cantidad, material.Fecha, material.Dolar, material.Bhabilitado)
}

//listar todos los registros

func ListarMateriales() Materiales {
	sql := `select id,descripcion,
precio,cantidad,fecha,dolar,
bhabilitado from material where id <10`
	materiales := Materiales{}
	rows, _ := db.Query(sql)
	for rows.Next() {
		material := Material{}
		rows.Scan(&material.Id, &material.Descripcion, &material.Precio, &material.Cantidad, &material.Fecha, &material.Dolar, &material.Bhabilitado)

		materiales = append(materiales, material)
	}
	return materiales
}

//obtener un registro por id

func GetMateriales(id int) *Material {
	material := NewMaterial("", 0, 0, "", 0, true)
	sql := `select id,descripcion,
precio,cantidad,fecha,dolar,
bhabilitado from material where id=?`
	rows, _ := db.Query(sql, id)
	for rows.Next() {
		rows.Scan(&material.Id, &material.Descripcion, &material.Precio, &material.Cantidad, &material.Fecha, &material.Dolar, &material.Bhabilitado)
	}
	return material
}

// actualizar registro no incluir el id van 6 campos
func (material *Material) update() {
	sql := `update material set descripcion=?,precio=?,cantidad=?,fecha=?,dolar=?,bhabilitado=?`
	db.Exec(sql, material.Descripcion, material.Precio, material.Cantidad, material.Fecha, material.Dolar, material.Bhabilitado)

}

// guardar o editar materiales
func (material *Material) Save() {
	if material.Id == 0 {
		material.Insert()
	} else {
		material.update()
	}
}

// eliminar registro
func (material *Material) Delete() {
	sql := `delete from material where id =?`
	db.Exec(sql, material.Id)
}
func TruncateTable(tablename string) {
	sql := fmt.Sprintf("truncate %s", tablename)
	db.Exec(sql)
}
