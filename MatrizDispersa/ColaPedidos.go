package MatrizDispersa

type NodoProductos struct {
	CodigoProducto int
	Siguiente *NodoProductos
}

type ListaProdutos struct {
	Primero *NodoProductos
	Ultimo *NodoProductos
	tamanio int
}

type NodoCola struct {
	ListaProd *ListaProdutos
	Siguiente *NodoCola
}

type ColaPedidos struct {
	Primero *NodoCola
	Ultimo *NodoCola
	Tamanio int
}

func New_NodoProducto(codigo int) *NodoProductos {
	return &NodoProductos{codigo, nil}
}
func New_ListaProductos() *ListaProdutos {
	return &ListaProdutos{nil,nil, 0}
}

func New_NodoCola(lista *ListaProdutos) *NodoCola {
	return &NodoCola{lista, nil}
}

func New_ColaPedido() *ColaPedidos {
	return &ColaPedidos{nil, nil,0}
}

func InsertarProductos(codigo int, lista *ListaProdutos)  {
	nuevoNodo := New_NodoProducto(codigo)
	if lista.Primero == nil{
		lista.Primero = nuevoNodo
		lista.Ultimo = nuevoNodo
		lista.tamanio += 1
	}else{
		lista.Ultimo.Siguiente = nuevoNodo
		lista.Ultimo = nuevoNodo
		lista.tamanio += 1
	}
}

func Push(cola *ColaPedidos, lista *ListaProdutos)  {
	nuevoNodo := New_NodoCola(lista)
	if cola.Primero == nil{
		cola.Primero = nuevoNodo
		cola.Ultimo = cola.Primero
		cola.Tamanio += 1
	}else {
		cola.Ultimo.Siguiente = nuevoNodo
		cola.Ultimo = nuevoNodo
		cola.Tamanio += 1
	}
}
