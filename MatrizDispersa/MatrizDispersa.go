package MatrizDispersa

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

type NodoPrincipal struct {
	headLsColumn *ListaColumnas
	headLsRow *ListaFilas
}

type NodoContenido struct {
	Fila *NodoFila
	Columna *NodoColumna
	ColaPedidos *ColaPedidos
	Siguiente *NodoContenido
}

type MatrizDisp struct {
	Principal *NodoPrincipal
	PrimerContenido *NodoContenido
	UltimoContenido *NodoContenido
}

func New_NodoPrincipal(dia int, depa string) *NodoPrincipal {
	nuevaCol := New_ListaCol()
	InsertarColumna(dia, &(*nuevaCol))
	nuevaFila := New_ListaFila()
	Insertar_Fila(depa, &(*nuevaFila))
	return &NodoPrincipal{&*nuevaCol, &*nuevaFila}
}

func New_NodoContenido(depa string, dia int, cola *ColaPedidos, matriz *MatrizDisp) *NodoContenido {
	Insertar_Fila(depa, matriz.Principal.headLsRow)
	InsertarColumna(dia, matriz.Principal.headLsColumn)
	fila := BuscarFila(depa, matriz.Principal.headLsRow)
	columna := BuscarColumna(dia, matriz.Principal.headLsColumn)
	return &NodoContenido{&(*fila), &(*columna), &(*cola), nil}
}

func New_MatrizDis(dia int, depa string) *MatrizDisp {
	nuevoPrincipal := New_NodoPrincipal(dia, depa)
	return &MatrizDisp{nuevoPrincipal, nil, nil}
}

func New_Matriz2() *MatrizDisp {
	return &MatrizDisp{nil,nil,nil}
}
func Insertar_NodoPrincipal(dia int, depa string) *MatrizDisp {
	nuevoPrincipal := New_NodoPrincipal(dia, depa)
	return &MatrizDisp{nuevoPrincipal, nil, nil}
}

func InsertarEnMatriz(dia int, departamento string, colaPed *ColaPedidos, matriz *MatrizDisp)  {
	if matriz.PrimerContenido == nil{
		matriz.PrimerContenido = New_NodoContenido(departamento, dia, &(*colaPed) ,&(*matriz))
		matriz.UltimoContenido = matriz.PrimerContenido
	}else {
		matriz.UltimoContenido.Siguiente = New_NodoContenido(departamento, dia, &(*colaPed), &(*matriz))
		matriz.UltimoContenido = matriz.UltimoContenido.Siguiente
	}
}

func Find_NodoContenido(fila string, columna int, matriz *MatrizDisp) *NodoContenido {
	aux := matriz.PrimerContenido
	for aux != nil{
		if aux.Columna.Dia == columna && aux.Fila.Departamento == fila{
			fmt.Println("Encontro nodo de cola")
			return &(*aux)
		}
		aux = aux.Siguiente
	}
	fmt.Println("No esta la cola buscada")
	return &NodoContenido{nil,nil,nil,nil}
}

func ListarFilas(matriz *MatrizDisp){
	aux := matriz.Principal.headLsRow.Primero
	for aux != nil{
		fmt.Print(aux.Departamento + " -> ")
		aux = aux.Siguiente
	}
}
func ListarColumnas(matriz *MatrizDisp)  {
	aux := matriz.Principal.headLsColumn.Primero
	for aux != nil{
		dia := strconv.Itoa(aux.Dia)
		fmt.Print(dia + " -> ")
		aux = aux.Siguiente
	}
}

func GraficarMatriz(matriz *MatrizDisp){
	columnas := matriz.Principal.headLsColumn.Primero
	filas := matriz.Principal.headLsRow.Primero
	contenido := matriz.PrimerContenido
	acum1 := ""
	acum2 := ""
	acum3 := ""
	enlacesFila := ""
	enlacesColumna := ""

	for columnas != nil{
		acum1 += strconv.Itoa(columnas.Dia) + "[label=\""+strconv.Itoa(columnas.Dia)+"\" width=2 style=filled, fillcolor=bisque1, group=1];\n"
		columnas = columnas.Siguiente
	}
	columnas = matriz.Principal.headLsColumn.Primero
	for columnas != nil{
		if columnas.Siguiente != nil{
			enlacesColumna += strconv.Itoa(columnas.Dia) + " -> "
		}else{
			enlacesColumna += strconv.Itoa(columnas.Dia) + ";\n"
		}
		columnas = columnas.Siguiente
	}
	for filas != nil{
		acum2 += "\""+filas.Departamento+"\"" + "[label=\""+filas.Departamento+"\" style=filled, fillcolor=lightskyblue, group=\""+filas.Departamento+"\"];\n"
		filas = filas.Siguiente
	}
	filas = matriz.Principal.headLsRow.Primero
	for filas != nil{
		if filas.Siguiente != nil{
			enlacesFila += "\""+filas.Departamento+"\"" + " -> "
		}else{
			enlacesFila += "\""+filas.Departamento+"\";\n"
		}

		filas = filas.Siguiente
	}
	ranksame := "{rank=same; mt; "+enlacesColumna+"}\n"
	ranksame2 := strings.ReplaceAll(ranksame, "->", ";")
	estructuraDot := "digraph G{\nnode [shape=box];\n"
	estructuraDot += "mt[label=\"Principal\"];\n"

	//auxRank := ""
	for contenido != nil{
		acum3 += "\""+fmt.Sprint(&(*contenido))+"\""+"[label=\"Tamaño"+strconv.Itoa(contenido.ColaPedidos.Tamanio)+" Dia:"+ strconv.Itoa(contenido.Columna.Dia)+" \", group=\""+contenido.Fila.Departamento+"\"];\n"
		//acum3 += "\""+contenido.Fila.Departamento+"\"" + " -> " + "\""+fmt.Sprint(&(*contenido))+"\"\n"
		acum3 += "\""+strconv.Itoa(contenido.Columna.Dia)+"\"" + " -> " + "\""+fmt.Sprint(&(*contenido))+"\"\n"
		//auxRank += "\""+contenido.Fila.Departamento+"\";" + "\""+fmt.Sprint(&(*contenido))+"\";\n"
		contenido = contenido.Siguiente
	}

	contenido = matriz.PrimerContenido

	filas = matriz.Principal.headLsRow.Primero
	ranksameCont := ""
	for filas != nil {
		ranksameCont += "{rank=same;"
		for contenido != nil{
			if contenido.Fila.Departamento == filas.Departamento{
				ranksameCont += "\""+contenido.Fila.Departamento+"\";" + "\""+fmt.Sprint(&(*contenido))+"\";\n"
			}
			contenido = contenido.Siguiente
		}
		ranksameCont += "}"
		contenido = matriz.PrimerContenido
		filas = filas.Siguiente
	}
	//ranksameCont := "{rank=same;"+auxRank+"}"

	estructuraDot += acum1
	estructuraDot += "mt -> " + enlacesColumna
	estructuraDot += acum2
	estructuraDot += "mt ->" + enlacesFila
	estructuraDot += ranksame2
	estructuraDot += acum3
	estructuraDot += ranksameCont
	estructuraDot += "\n}\n"
	path := "matriz.dot"
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
	cmd, _ := exec.Command(path2, "-Tpng","matriz.dot").Output()
	mode := int(0777)
	ioutil.WriteFile("matriz.png",cmd, os.FileMode(mode))

}

func existeError(err error) bool{
	if err != nil{
		fmt.Println(err.Error())
	}
	return (err != nil)
}