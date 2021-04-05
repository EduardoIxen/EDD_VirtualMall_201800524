package Pedidos

import (
	"ProyectoEDD/MatrizDispersa"
	"fmt"
	"strconv"
)

type NodoLsPedido struct {
	Siguiente *NodoLsPedido
	Anterior *NodoLsPedido
	IdNodo int
	NumeroMes int
	NombreMes string
	MatrizDispersa *MatrizDispersa.MatrizDisp
}

type ListaDoblePedidos struct {
	Primero *NodoLsPedido
	contadorNodos int
}

func New_Nodo(numeroMes int) *NodoLsPedido {
	return &NodoLsPedido{nil, nil, 0,numeroMes,"nil", nil}
}

func New_ListaPedidos() *ListaDoblePedidos {
	return &ListaDoblePedidos{nil, 0}
}

func InsertarPedido(numeroMes int, listaDoble *ListaDoblePedidos)  {
	var nuevoNodo =  New_Nodo(numeroMes)
	nuevoNodo.IdNodo = listaDoble.contadorNodos
	if numeroMes == 1{
		nuevoNodo.NombreMes = "ENERO"
	}else if numeroMes == 2{
		nuevoNodo.NombreMes = "FEBRERO"
	}else if numeroMes == 3{
		nuevoNodo.NombreMes = "MARZO"
	}else if numeroMes == 4{
		nuevoNodo.NombreMes = "ABRIL"
	}else if numeroMes == 5{
		nuevoNodo.NombreMes = "MAYO"
	}else if numeroMes == 6{
		nuevoNodo.NombreMes = "JUNIO"
	}else if numeroMes == 7{
		nuevoNodo.NombreMes = "JULIO"
	}else if numeroMes == 8{
		nuevoNodo.NombreMes = "AGOSTO"
	}else if numeroMes == 9{
		nuevoNodo.NombreMes = "SEPTIEMBRE"
	}else if numeroMes == 10{
		nuevoNodo.NombreMes = "OCTUBRE"
	}else if numeroMes == 11{
		nuevoNodo.NombreMes = "NOVIEMBRE"
	}else if numeroMes == 12{
		nuevoNodo.NombreMes = "DICIEMBRE"
	}

	if listaDoble.Primero == nil{
		listaDoble.Primero = nuevoNodo
		listaDoble.contadorNodos += 1
	}else{
		aux := listaDoble.Primero
		for aux != nil{
			if aux.NumeroMes < numeroMes{
				if aux.Siguiente == nil{
					aux.Siguiente = nuevoNodo
					nuevoNodo.Anterior = aux
					listaDoble.contadorNodos += 1
					break
				}else {
					aux = aux.Siguiente
				}
			}else if aux.NumeroMes > numeroMes{
				if aux == listaDoble.Primero{
					listaDoble.Primero = nuevoNodo
					nuevoNodo.Siguiente = aux
					aux.Anterior = nuevoNodo
					listaDoble.contadorNodos += 1
					break
				}else {
					aux.Anterior.Siguiente = nuevoNodo
					nuevoNodo.Anterior = aux.Anterior
					aux.Anterior = nuevoNodo
					nuevoNodo.Siguiente = aux
					listaDoble.contadorNodos += 1
					break
				}
			}else if numeroMes == aux.NumeroMes{
				 fmt.Println("Los numeros son iguales")
				 break
			}
		}
	}
}

func ListarPedidoFecha(arbol *ListaDoblePedidos) {
	aux := arbol.Primero
	for aux != nil{
		fmt.Print(aux.NombreMes +"("+strconv.Itoa(aux.IdNodo)+")"+ " -> ")
		aux = aux.Siguiente
	}
}

func BuscarNodoPedido(mes int, lista *ListaDoblePedidos) *NodoLsPedido{
	aux := lista.Primero
	for aux != nil{
		if aux.NumeroMes == mes{
			return aux
		}
		aux = aux.Siguiente
	}
	fmt.Println("Buscar nodo pedido nil")
	return nil
}