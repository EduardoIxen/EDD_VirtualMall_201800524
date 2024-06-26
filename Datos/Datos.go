package Datos

import "ProyectoEDD/Arboles"

type EstructuraGeneral struct {
	Datos []Datos `json:"Datos"`
}
type Tienda struct {
	Nombre       string `json:"Nombre"`
	Descripcion  string `json:"Descripcion"`
	Contacto     string `json:"Contacto"`
	Calificacion int    `json:"Calificacion"`
	Logo 		 string `json:"Logo"`
	ArbolProd	 *Arboles.Arbol
}
type Departamento struct {
	Nombre  string    `json:"Nombre"`
	Tiendas []Tienda `json:"Tiendas"`
}
type Datos struct {
	Indice        string          `json:"Indice"`
	Departamentos []Departamento `json:"Departamentos"`
}

func NewTienda(nombre string, descripcion string, contacto string, calificacion int, logo string) *Tienda {
	return &Tienda{nombre, descripcion, contacto, calificacion, logo, nil}
}
