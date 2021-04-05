package MatrizDispersa

import "fmt"

//-------------Nodo y lista para encabezado columnas--------------
type NodoColumna struct {
	Dia int
	Siguiente *NodoColumna
	Anterior *NodoColumna
}
type ListaColumnas struct {
	Primero *NodoColumna
}

func New_NodoCol(dia int) *NodoColumna {
	return &NodoColumna{dia, nil, nil}
}

func New_ListaCol() *ListaColumnas {
	return &ListaColumnas{nil}
}

//------------------Nodo y lista encabezado filas---------------

type NodoFila struct {
	Departamento string
	Siguiente *NodoFila
	Anterior *NodoFila
}
type ListaFilas struct {
	Primero *NodoFila
	Ultimo *NodoFila
}

func New_NodoFila(departamento string) *NodoFila {
	return &NodoFila{departamento,nil,nil}
}

func New_ListaFila() *ListaFilas {
	return &ListaFilas{nil, nil}
}

//////////////////METODOS PARA INSERTAR EN LAS LISTAS//////////////////77

func InsertarColumna(dia int, listaDoble *ListaColumnas)  {
	var nuevoNodo = New_NodoCol(dia)
	if listaDoble.Primero == nil{
		listaDoble.Primero = nuevoNodo
	}else{
		aux := listaDoble.Primero
		for aux != nil{
			if aux.Dia < dia{
				if aux.Siguiente == nil{
					aux.Siguiente = nuevoNodo
					nuevoNodo.Anterior = aux
					break
				}else {
					aux = aux.Siguiente
				}
			}else if aux.Dia > dia{
				if aux == listaDoble.Primero{
					listaDoble.Primero = nuevoNodo
					nuevoNodo.Siguiente = aux
					aux.Anterior = nuevoNodo
					break
				}else {
					aux.Anterior.Siguiente = nuevoNodo
					nuevoNodo.Anterior = aux.Anterior
					aux.Anterior = nuevoNodo
					nuevoNodo.Siguiente = aux
					break
				}
			}else if aux.Dia == dia{
				fmt.Println("El dia ya existe en la lista")
				break
			}
		}
	}
}

func BuscarColumna(dia int, lista *ListaColumnas ) *NodoColumna {
	if lista.Primero != nil{
		encontrado := false
		aux := lista.Primero
		for aux != nil{
			if aux.Dia == dia{
				encontrado = true
				return &*aux
			}
			aux = aux.Siguiente
		}
		if !encontrado{
			return 	nil
		}else{
			return &*aux
		}
	}else {
		return nil
	}
}

func Insertar_Fila(departamento string, listaDoble *ListaFilas)  {
	var nuevoNodo = New_NodoFila(departamento)
	if listaDoble.Primero == nil{
		listaDoble.Primero = nuevoNodo
		listaDoble.Ultimo = listaDoble.Primero
	}else {
		aux := listaDoble.Primero
		var esta = false
		for aux != nil{
			if aux.Departamento == departamento{
				esta = true
			}
			aux = aux.Siguiente
		}
		if !esta{
			listaDoble.Ultimo.Siguiente = nuevoNodo
			listaDoble.Ultimo = nuevoNodo
		}

	}
}
func BuscarFila(departamento string, filas *ListaFilas) *NodoFila{
	if filas.Primero != nil{
		aux := filas.Primero
		encontrado := false
		for aux != nil{
			if aux.Departamento == departamento{
				encontrado = true
				return &*aux
			}
			aux = aux.Siguiente
		}
		if !encontrado{
			return nil
		}else {
			return &*aux
		}
	}else {
		return nil
	}
}