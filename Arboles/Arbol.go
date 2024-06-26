package Arboles

import "fmt"

var increase = false

type Arbol struct {
	Raiz *NodoArbol
}

//constructor
func Nuevo_Arbol() *Arbol {
	return &Arbol{nil}
}

func Insert(arbol *Arbol, producto *Producto) bool {
	if Find(arbol,*producto) == nil{
		fmt.Println("No esta ", producto.Codigo)
		increase = false
		return Insertar2(&arbol.Raiz, producto, &increase)
	}else {
		fmt.Println("Ya esta dentro ", producto.Codigo)
		Find(arbol, *producto).Cantidad = Find(arbol, *producto).Cantidad + producto.Cantidad
		return false
	}
}

func Insertar2(raiz_local **NodoArbol, producto *Producto, increase *bool) bool {
	if *raiz_local == nil{
		*raiz_local = New_NodoArbol(producto)
		*increase = true
		//fmt.Println("inserto en la raiz")
		return true
	}else{
		if producto.Codigo < (*raiz_local).Producto.Codigo {
			return_value := Insertar2(&(*raiz_local).Left, producto, increase)

			if *increase { //si bubo una insercion
				switch (*raiz_local).Balance {
				case BALANCED:
					//LEFT-HEAVY
					(*raiz_local).Balance = LEFT_HEAVY
				case RIGHT_HEAVY:
					//RIGHT_HEAVY
					(*raiz_local).Balance = BALANCED
					*increase = false
					break
				case LEFT_HEAVY:
					//CRITICO PORQUE TIENE LEFT-LEFT -> BALANCE DE -2
					Rebalance_Left(&*raiz_local)
					*increase = false
					break
				}
			}
			return return_value
		}else if (producto.Codigo > (*raiz_local).Producto.Codigo) {
			return_value :=  Insertar2(&(*raiz_local).Right,producto,increase) //RETORNA UN BOOL
			if *increase {                                                     //si hubo una insercion
				switch (*raiz_local).Balance {
				case BALANCED:
					// RIGHT-HEAVY
					(*raiz_local).Balance = RIGHT_HEAVY
				case LEFT_HEAVY:
					//LEFT-HEAVY
					(*raiz_local).Balance = BALANCED
					*increase = false
					break
				case RIGHT_HEAVY:
					//ES CRITICO PORQUE TIENE UN RIGHT-RIGHT POR LO TANTO LA RAIZ ES -2
					Rebalance_right(&*raiz_local)
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

func Rebalance_Left(raiz_local **NodoArbol)  {
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

func Rebalance_right(local_root **NodoArbol){
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

func Rotate_Right(local_root **NodoArbol){
	tmp := (*local_root).Left
	(*local_root).Left = tmp.Right
	tmp.Right = *local_root
	*local_root = tmp
}

func Rotate_Left(local_root **NodoArbol){
	tmp := (*local_root).Right
	(*local_root).Right = tmp.Left
	tmp.Left = *local_root
	*local_root = tmp
}

func Find( tree *Arbol, prod Producto) *Producto{
	return Find_2(tree.Raiz, prod)
}

func Find_2(local_root *NodoArbol, prod Producto) *Producto{
	if local_root == nil {
		return nil
	}
	if prod.Codigo < local_root.Producto.Codigo {
		return Find_2(local_root.Left,prod)
	} else if prod.Codigo > local_root.Producto.Codigo {
		return Find_2(local_root.Right,prod)
	} else {
		return local_root.Producto
	}
}