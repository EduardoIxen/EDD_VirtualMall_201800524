package Arboles

type NodoArbol struct {
	Producto *Producto
	Left *NodoArbol
	Right *NodoArbol
	Balance int
}

const (
	LEFT_HEAVY = -1
	BALANCED = 0
	RIGHT_HEAVY = +1
)

func New_NodoArbol(producto *Producto) *NodoArbol {
	return &NodoArbol{producto,nil,nil,0}
}
func New_NodoArbol_2(producto *Producto, left_child *NodoArbol, right_child *NodoArbol) *NodoArbol {
	return &NodoArbol{producto, left_child, right_child, 0}
}