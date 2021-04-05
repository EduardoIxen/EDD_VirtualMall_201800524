package Pedidos

import "fmt"

var increase = false

type ArbolPedido struct {
	Raiz *NodoPedido
}

//constructor
func Nuevo_ArbolPedido() *ArbolPedido {
	return &ArbolPedido{nil}
}

func Insertar(arbol *ArbolPedido, anio int) bool {
	if Find(arbol, anio) == nil{
		increase = false
		return Insertar2(&arbol.Raiz, anio, &increase)
	}else {
		fmt.Println("El año ya existe")
		return false
	}
}

func Insertar2(local_root **NodoPedido, anio int, increase *bool) bool {
	if *local_root == nil{
		*local_root = New_NodoPedido(anio)
		*increase = true
		//fmt.Println("inserto en la raiz")
		return true
	}else{
		if anio < (*local_root).Anio {
			return_value := Insertar2(&(*local_root).Left, anio, increase)

			if *increase { //si bubo una insercion
				switch (*local_root).Balance {
				case BALANCED:
					//LEFT-HEAVY
					(*local_root).Balance = LEFT_HEAVY
				case RIGHT_HEAVY:
					//RIGHT_HEAVY
					(*local_root).Balance = BALANCED
					*increase = false
					break
				case LEFT_HEAVY:
					//CRITICO PORQUE TIENE LEFT-LEFT -> BALANCE DE -2
					Rebalance_Left(&*local_root)
					*increase = false
					break
				}
			}
			return return_value
		}else if (anio > (*local_root).Anio) {
			return_value :=  Insertar2(&(*local_root).Right,anio,increase) //RETORNA UN BOOL
			if *increase {                                                     //si hubo una insercion
				switch (*local_root).Balance {
				case BALANCED:
					// RIGHT-HEAVY
					(*local_root).Balance = RIGHT_HEAVY
				case LEFT_HEAVY:
					//LEFT-HEAVY
					(*local_root).Balance = BALANCED
					*increase = false
					break
				case RIGHT_HEAVY:
					//ES CRITICO PORQUE TIENE UN RIGHT-RIGHT POR LO TANTO LA RAIZ ES -2
					Rebalance_right(&*local_root)
					*increase = false
					break
				}
			}
			return return_value
		} else {
			return false
		}
	}
}

func Find(arbol *ArbolPedido, anio int) *NodoPedido {
	return Find2(arbol.Raiz, anio)
}

func Find2(local_root *NodoPedido, anio int) *NodoPedido {
	if local_root == nil{
		return nil
	}
	if anio < local_root.Anio {
		return Find2(local_root.Left, anio)
	}else if anio > local_root.Anio {
		return Find2(local_root.Right, anio)
	}else {
		return local_root
	}
}

func Rebalance_Left(raiz_local **NodoPedido)  {
	left_child := (*raiz_local).Left
	if left_child.Balance == RIGHT_HEAVY {
		//OBTENER REFERENCIA DE LEFT-RIGHT CHILD
		left_right_child := left_child.Right
		//se ajustan los nuevos balances despues de realizar la rotacion
		if left_right_child.Balance == LEFT_HEAVY {
			left_child.Balance = BALANCED
			left_right_child.Balance = BALANCED
			(*raiz_local).Balance = RIGHT_HEAVY
		}else if left_right_child.Balance == BALANCED {
			left_child.Balance = BALANCED
			left_right_child.Balance = BALANCED
			(*raiz_local).Balance = BALANCED
		}else {
			left_child.Balance = LEFT_HEAVY
			left_right_child.Balance = BALANCED
			(*raiz_local).Balance = BALANCED
		}
		//REALIZO LEFT ROTATION
		Rotate_Left(&(*raiz_local).Left)
	}else{ //CASO LEFT-LEFT
		left_child.Balance = BALANCED
		(*raiz_local).Balance = BALANCED
	}
	//REALIZO RIGHT ROTATION
	Rotate_Right(&*raiz_local)
}

func Rebalance_right(local_root **NodoPedido){
	right_child := (*local_root).Right
	if right_child.Balance == LEFT_HEAVY { // CASO RIGH-LEFT
		//OBTENGO UNA REFERENCIA DE LEFT-RIGHT CHILD
		right_left_child := right_child.Left
		//se ajustan los nuevos valances despues de haber realizado la rotacion
		if right_left_child.Balance == RIGHT_HEAVY {
			right_child.Balance = BALANCED
			right_left_child.Balance = BALANCED
			(*local_root).Balance = LEFT_HEAVY
		} else if right_left_child.Balance == BALANCED {
			right_child.Balance = BALANCED
			right_left_child.Balance = BALANCED
			(*local_root).Balance = BALANCED
		} else {
			right_child.Balance = RIGHT_HEAVY
			right_left_child.Balance = BALANCED
			(*local_root).Balance = BALANCED
		}
		// REALIZO RIGH ROTATION
		Rotate_Right(&(*local_root).Right)
	} else { // CASO RIGHT-RIGHT
		right_child.Balance = BALANCED
		(*local_root).Balance = BALANCED
	}
	//REALIZO RIGHT ROTATION
	Rotate_Left(&*local_root)
}

func Rotate_Left(local_root **NodoPedido){
	tmp := (*local_root).Right
	(*local_root).Right = tmp.Left
	tmp.Left = *local_root
	*local_root = tmp
}

func Rotate_Right(local_root **NodoPedido){
	tmp := (*local_root).Left
	(*local_root).Left = tmp.Right
	tmp.Right = *local_root
	*local_root = tmp
}