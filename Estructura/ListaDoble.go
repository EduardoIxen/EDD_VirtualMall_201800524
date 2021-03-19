package Estructura

import (
	"ProyectoEDD/Datos"
	"fmt"
)

type Nodo struct {
	Siguiente *Nodo
	Anterior *Nodo
	Tienda *Datos.Tienda
	IdNodo int
}

type ListaDoble struct {
	Primero *Nodo
	Ultimo *Nodo
	Identificador  string
	Contador int
	Departamento string
	Indice string
}

func NewNodo(tienda *Datos.Tienda, idNodo int) *Nodo  {
	return &Nodo{nil,nil,tienda, idNodo}
}

func NewListaDoble(identificador string, departamento string, indice string) *ListaDoble {
	return &ListaDoble{nil,nil, identificador, 0 , departamento, indice}
}

func Insertar(tienda *Datos.Tienda, listaDoble *ListaDoble, idNodo int)  {
	var nuevoNodo = NewNodo(tienda, idNodo)

	if listaDoble.Primero == nil{ //lista vacía inserto en el primero
		listaDoble.Primero = nuevoNodo
		listaDoble.Ultimo = nuevoNodo
		listaDoble.Contador += 1
		fmt.Println("Agregado en cabeza")
		//Imprimir(listaDoble)
	}else { //para una lista no vacia
		aux := listaDoble.Primero

		for aux != nil{
			if aux.Tienda.Nombre > tienda.Nombre {
				if aux.Siguiente == nil{ //llego al final de la lista, si o si insertar nodo
					aux.Siguiente = nuevoNodo
					nuevoNodo.Anterior = aux
					listaDoble.Contador += 1
					fmt.Println("Agregado al final")
					//Imprimir(listaDoble)
					break
				}else {
					aux = aux.Siguiente
				}
			}else if aux.Tienda.Nombre < tienda.Nombre{
				if aux == listaDoble.Primero{
					listaDoble.Primero = nuevoNodo
					nuevoNodo.Siguiente = aux
					aux.Anterior = nuevoNodo
					listaDoble.Contador += 1
					fmt.Println("Agregado en cabeza no primero")
					break
					//Imprimir(listaDoble)
				}else{
					aux.Anterior.Siguiente = nuevoNodo
					nuevoNodo.Anterior = aux.Anterior
					aux.Anterior = nuevoNodo
					nuevoNodo.Siguiente = aux
					listaDoble.Contador += 1
					fmt.Println("entre nodos")
					//Imprimir(listaDoble)
					break
				}
			}
		}
	}
}
func Imprimir(listaR *ListaDoble){
	aux := listaR.Primero
	for aux != nil{
		fmt.Printf("Nombre %v ->", aux.Tienda.Nombre)
		aux = aux.Siguiente
	}
	fmt.Println("\n")
}