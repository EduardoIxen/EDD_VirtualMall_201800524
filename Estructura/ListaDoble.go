package Estructura

import (
	"ProyectoEDD/Datos"
	"fmt"
)

type Nodo struct {
	siguiente *Nodo
	anterior *Nodo

	tienda *Datos.Tienda
}

type ListaDoble struct {
	primero *Nodo
	ultimo *Nodo
	idNodo int
}

func NewNodo(tienda *Datos.Tienda) *Nodo  {
	return &Nodo{nil,nil,tienda}
}

func NewListaDoble() *ListaDoble {
	return &ListaDoble{nil,nil, 0}
}

func Insertar(tienda *Datos.Tienda, listaDoble *ListaDoble)  {
	var nuevoNodo = NewNodo(tienda)

	if listaDoble.primero == nil{ //lista vacía inserto en el primero
		listaDoble.primero = nuevoNodo
		listaDoble.ultimo = nuevoNodo
		listaDoble.idNodo += 1
		Imprimir(listaDoble)
	}else { //para una lista no vacia
		aux := listaDoble.primero

		for aux != nil{
			if aux.tienda.Nombre > tienda.Nombre {
				if aux.siguiente == nil{ //llego al final de la lista, si o si insertar nodo
					aux.siguiente = nuevoNodo
					nuevoNodo.anterior = aux
					listaDoble.idNodo += 1
					Imprimir(listaDoble)
					break
				}else {
					aux = aux.siguiente
				}
			}else if aux.tienda.Nombre < tienda.Nombre{
				if aux == listaDoble.primero{
					listaDoble.primero = nuevoNodo
					nuevoNodo.siguiente = aux
					aux.anterior = nuevoNodo
					listaDoble.idNodo += 1
					Imprimir(listaDoble)
				}else{
					aux.anterior.siguiente = nuevoNodo
					aux.anterior = nuevoNodo
					nuevoNodo.siguiente = aux
					nuevoNodo.anterior = aux.anterior
					listaDoble.idNodo += 1
					Imprimir(listaDoble)
					break
				}
			}
		}
	}
}
func Imprimir(lista *ListaDoble){
	fmt.Println("LLego a print")
	aux :=lista.primero
	for aux != nil{
		fmt.Printf("Nombre %v ->", aux.tienda.Nombre)
		aux = aux.siguiente
	}
	fmt.Println("\n")
}