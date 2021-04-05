package Pedidos

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
)

func Generar_Grafo(arbol *ArbolPedido){
	estructuraDot := "digraph G{\nnode [shape=circle];\n"
	acum := ""

	if arbol.Raiz != nil{
		Recorrer_Arbol(&arbol.Raiz,&acum)
	}

	estructuraDot += acum + "\n}\n"

	path := "grafoPedidos.dot"
	//SE ESCRIBE EL ARCHIVO .DOT
	var _, err = os.Stat(path)
	if os.IsNotExist(err){
		var file, err = os.Create(path)
		if existeError(err) {
			return
		}
		defer file.Close()
		fmt.Println("Se ha creado un archivo")
	}

	var file, err2 = os.OpenFile(path, os.O_RDWR, 0644)
	if existeError(err2) {
		return
	}
	defer file.Close()

	//SE ESCRIBE EN ARCHIVO
	_, err = file.WriteString(estructuraDot)
	if existeError(err) {
		return
	}

	// Salva los cambios
	err = file.Sync()
	if existeError(err) {
		return
	}

	fmt.Println("Archivo actualizado existosamente.")

	//PARTE EN DONDE GENERO EL GRAFO
	path2, _ := exec.LookPath("dot")
	cmd, _ := exec.Command(path2, "-Tpng","grafoPedidos.dot").Output()
	mode := int(0777)
	ioutil.WriteFile("grafoPedidos.png",cmd, os.FileMode(mode))

}

func Recorrer_Arbol(actual **NodoPedido, acum *string){

	if *actual != nil{

		//SE OBTIENE INFORMACION DEL NODO ACTUAL
		*acum += "\"" + fmt.Sprint(&(*actual)) + "\"[label=\"" + strconv.Itoa((*actual).Anio) +"\"];\n"
		//VIAJAMOS A LA SUBRAMA IZQ
		if (*actual).Left != nil{
			*acum += "\"" +  fmt.Sprint(&(*actual)) + "\" -> \"" + fmt.Sprint(&(*actual).Left) + "\";\n"
		}
		//VIAJAMOS A LA SUBRAMA DER
		if (*actual).Right != nil{
			*acum += "\"" +  fmt.Sprint(&(*actual)) + "\" -> \"" + fmt.Sprint(&(*actual).Right) + "\";\n"
		}

		Recorrer_Arbol(&(*actual).Left,acum)
		Recorrer_Arbol(&(*actual).Right, acum)

	}
}

func existeError(err error) bool{
	if err != nil{
		fmt.Println(err.Error())
	}
	return (err != nil)
}
