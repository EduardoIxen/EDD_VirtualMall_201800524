package Pedidos

type NodoPedido struct {
	Anio int
	Left *NodoPedido
	Right *NodoPedido
	Balance int
	ListaPedidos *ListaDoblePedidos
}

const (
	LEFT_HEAVY = -1
	BALANCED = 0
	RIGHT_HEAVY = +1
)

func New_NodoPedido(anio int) *NodoPedido{
	return &NodoPedido{anio, nil,nil,0, nil}
}
func New_NodoPedido2(anio int, lef_child *NodoPedido, right_child *NodoPedido)	*NodoPedido  {
	return &NodoPedido{anio,lef_child, right_child, 0, nil}
}
