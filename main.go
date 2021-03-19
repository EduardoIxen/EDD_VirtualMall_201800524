package main

import (
	"ProyectoEDD/Arboles"
	"ProyectoEDD/Datos"
	"ProyectoEDD/Estructura"
	"ProyectoEDD/Grafo"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strconv"
)

type calificaciones []*Estructura.ListaDoble
var vector = calificaciones{}

type allData []*Datos.EstructuraGeneral
var listaDatos = allData{}

type allDep []string
var lsDep = allDep{}

type allIndex []string
var lsIndex = allIndex{}

type busqueda struct {
	Departamento string `json:"Departamento"`
	Nombre       string `json:"Nombre"`
	Calificacion int    `json:"Calificacion"`
}

type allInv []*Arboles.ContInventarios
var listaInventarios = allInv{}

type singleInv []*Arboles.Inventario
var lsSingleInv = singleInv{}

func cargartienda(w http.ResponseWriter, r *http.Request)  {
	listaDatos = allData{}
	vector = calificaciones{}
	lsDep = allDep{}
	lsIndex = allIndex{}
	var nuevaEntrada *Datos.EstructuraGeneral
	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil{
		fmt.Fprintf(w, "----Error al cargar la entrada----")
	}
	json.Unmarshal(reqBody, &nuevaEntrada)
	listaDatos = append(listaDatos, nuevaEntrada)

	indices := 0
	departamentos := 0

	for _, dat := range listaDatos{
		for _, indx := range dat.Datos{
			indices += 1
			lsIndex = append(lsIndex, indx.Indice)
		}
	}
	fmt.Printf("Numerod de Indices %v \n", indices)
	for _, ind := range lsIndex{
		fmt.Println(ind)
	}
	for _, dat := range listaDatos{
		for _, indice := range dat.Datos{
			for _, depa := range indice.Departamentos {
				departamentos += 1
				lsDep = append(lsDep, depa.Nombre)
			}
			//lsDep = append(lsDep, "/////////////////////////////////////////////")  //OJOOOOOO ACA
			break
		}
	}
	fmt.Printf("Numero de departamentos %v \n",departamentos)
	for _, dep := range lsDep{
		fmt.Println(dep)
	}
	fmt.Println("/////////////////////")
	for _, dat := range listaDatos{
		contador := 0
		for _, auxDepa := range lsDep{
			for _, contenido := range dat.Datos {
				for _, depa := range contenido.Departamentos{
					if depa.Nombre == auxDepa{
						fmt.Println("Departamento= ", depa.Nombre)
						var lista1 = Estructura.NewListaDoble("1", depa.Nombre, contenido.Indice)
						var lista2 = Estructura.NewListaDoble("2", depa.Nombre, contenido.Indice)
						var lista3 = Estructura.NewListaDoble("3", depa.Nombre, contenido.Indice)
						var lista4 = Estructura.NewListaDoble("4", depa.Nombre, contenido.Indice)
						var lista5 = Estructura.NewListaDoble("5", depa.Nombre, contenido.Indice)
						for _, tienda := range depa.Tiendas{
							contador += 1
							fmt.Println("Contadooooooorrrrrr", contador)
							fmt.Println("\t ", tienda.Nombre)
							if tienda.Calificacion == 1 {
								var t = Datos.NewTienda(tienda.Nombre, tienda.Descripcion, tienda.Contacto, tienda.Calificacion, tienda.Logo)
								Estructura.Insertar(t, lista1,contador)
							} else if tienda.Calificacion == 2 {
								var t = Datos.NewTienda(tienda.Nombre, tienda.Descripcion, tienda.Contacto, tienda.Calificacion, tienda.Logo)
								Estructura.Insertar(t, lista2,contador)
							} else if tienda.Calificacion == 3 {
								var t = Datos.NewTienda(tienda.Nombre, tienda.Descripcion, tienda.Contacto, tienda.Calificacion, tienda.Logo)
								Estructura.Insertar(t, lista3,contador)
							} else if tienda.Calificacion == 4 {
								var t = Datos.NewTienda(tienda.Nombre, tienda.Descripcion, tienda.Contacto, tienda.Calificacion, tienda.Logo)
								Estructura.Insertar(t, lista4,contador)
							} else if tienda.Calificacion == 5 {
								var t = Datos.NewTienda(tienda.Nombre, tienda.Descripcion, tienda.Contacto, tienda.Calificacion, tienda.Logo)
								Estructura.Insertar(t, lista5,contador)
							}
							fmt.Println("Lista1----------")
							Estructura.Imprimir(lista1)
							fmt.Println("Lista2---------------")
							Estructura.Imprimir(lista2)
							fmt.Println("Lista3------------------")
							Estructura.Imprimir(lista3)
							fmt.Println("Lista4---------------")
							Estructura.Imprimir(lista4)
							fmt.Println("Lista5------------------")
							Estructura.Imprimir(lista5)
						}
						vector = append(vector, lista1)
						vector = append(vector, lista2)
						vector = append(vector, lista3)
						vector = append(vector, lista4)
						vector = append(vector, lista5)
					}
				}
			}
		}
		fmt.Println("contador:", contador)
	}
	fmt.Println("///////////// Comienzo de vectores ////////////////")
	for _, vec := range vector{
		fmt.Println("vec..")
		Estructura.Imprimir(vec)
	}
	fmt.Println("Tamaño del vector",len(vector))
	//generarGrafo()

	//w.Header().Set("Content-Type", "application/json")
	//w.WriteHeader(http.StatusCreated)
	//json.NewEncoder(w).Encode(nuevaEntrada)
}

func cargarInventario(W http.ResponseWriter, r *http.Request){
	listaInventarios = allInv{}
	lsSingleInv = singleInv{}
	var nuevaEntrada *Arboles.ContInventarios
	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil{
		fmt.Fprintf(W, "----Error al cargar inventarios----")
	}

	json.Unmarshal(reqBody, &nuevaEntrada)
	listaInventarios = append(listaInventarios, nuevaEntrada)

	for _, dat := range listaInventarios {
		//fmt.Println("dat ",dat.Inventarios)
		for _, inv := range dat.Inventarios{
			fmt.Println("Departamento ",inv.Departamento)
			fmt.Println("Calificacion ",inv.Calificacion)
			fmt.Println("Tienda ",inv.Tienda)
			arbolInv := Arboles.Nuevo_Arbol()
			for _, prod := range inv.Productos{
				var pr = Arboles.NewProduct(prod.Nombre, prod.Codigo, prod.Descripcion, prod.Precio, prod.Cantidad, prod.Imagen)
				Arboles.Insert(arbolInv, pr)
				//fmt.Println(prod)
			}
			aux := &*vector[obtenerPosicion(inv.Departamento, inv.Calificacion, inv.Tienda)].Primero
			for aux != nil{
				if aux.Tienda.Nombre == inv.Tienda{
					aux.Tienda.ArbolProd = &*arbolInv
					//aux.Tienda.ArbolProd = *arbolInv
					fmt.Println("Arbol insertado")
				}
				aux = aux.Siguiente
			}
			/*deppr := vector[obtenerPosicion(inv.Departamento, inv.Calificacion, inv.Tienda)].Departamento
			indcpr := vector[obtenerPosicion(inv.Departamento, inv.Calificacion, inv.Tienda)].Indice
			calipr := vector[obtenerPosicion(inv.Departamento, inv.Calificacion, inv.Tienda)].Identificador*/

		}
	}
	Grafo.Generar_Grafo(vector[1].Primero.Tienda.ArbolProd)

}
func obtenerPosicion(departamento string, calificacion int, tienda string) int{
	//--------------ALGORITMO DE BUSQUEDA EN VECTORES COLUMN MAJOR-----------------
	var j = 0
	var i = 0
	var tamFilas = 0
	var k = 0
	for indice, dat := range lsDep  {
		if dat == departamento{
			j = indice
		}
	}
	for indice2, dat2 := range lsIndex{
		if dat2 == string(tienda[0]){
			i = indice2
		}
	}
	tamFilas = len(lsIndex)
	k = calificacion - 1
	return (j*tamFilas+i)*5+k
}
func generarGrafo(W http.ResponseWriter, r *http.Request)  {
	err3 := os.Remove("grafo.dot")
	if err3 != nil {
		fmt.Printf("Error eliminando archivo: %v\n", err3)
	} else {
		fmt.Println("Eliminado correctamente")
	}
	path := "grafo.dot"
	acum := "digraph G {\n node [shape=record]; \n "
	nodo := ""
	nodoLista := ""
	enlace := ""
	enlaceLista := ""
	for i, vect := range vector{
		if i == 0{
			nodo += "vec" +""+ "[height=1 width="+strconv.Itoa(len(vector)*3)+" label=\"<f"+strconv.Itoa(i+1)+"> "+strconv.Itoa(i+1)+" |"
			if vect.Primero != nil{
				aux := vect.Primero
				for aux != nil {
					nodoLista += strconv.Itoa(aux.IdNodo) + "[label=\"Nombre: "+aux.Tienda.Nombre+" \\n"
					nodoLista += " Calificacion: "+strconv.Itoa(aux.Tienda.Calificacion)+"\"];\n"
					if aux.Siguiente != nil{
						enlaceLista += strconv.Itoa(aux.IdNodo)+ "->" + strconv.Itoa(aux.Siguiente.IdNodo)+";\n"
						enlaceLista += strconv.Itoa(aux.Siguiente.IdNodo)+ "->" + strconv.Itoa(aux.IdNodo)+";\n"
					}
					if aux == vect.Primero{
						enlaceLista += "vec:f"+strconv.Itoa(i+1)+" -> " + strconv.Itoa(aux.IdNodo)+";\n"
					}
					aux = aux.Siguiente
				}
			}
		}else if i > 0 && i < len(vector)-1{
			nodo += "<f"+strconv.Itoa(i+1)+"> "+strconv.Itoa(i+1)+" |"
			if vect.Primero != nil{
				aux := vect.Primero
				for aux != nil {
					nodoLista += strconv.Itoa(aux.IdNodo) + "[label=\"Nombre: "+aux.Tienda.Nombre+" \\n"
					nodoLista += " Calificacion: "+strconv.Itoa(aux.Tienda.Calificacion)+"\"];\n"
					if aux.Siguiente != nil{
						enlaceLista += strconv.Itoa(aux.IdNodo)+ "->" + strconv.Itoa(aux.Siguiente.IdNodo)+";\n"
						enlaceLista += strconv.Itoa(aux.Siguiente.IdNodo)+ "->" + strconv.Itoa(aux.IdNodo)+";\n"
					}
					if aux == vect.Primero{
						enlaceLista += "vec:f"+strconv.Itoa(i+1)+" -> " + strconv.Itoa(aux.IdNodo)+";\n"
					}
					aux = aux.Siguiente
				}
			}
		}else if i == len(vector)-1{
			nodo += "<f"+strconv.Itoa(i+1)+"> "+strconv.Itoa(i+1)+"\"];\n"
			if vect.Primero != nil{
				aux := vect.Primero
				for aux != nil {
					nodoLista += strconv.Itoa(aux.IdNodo) + "[label=\"Nombre: "+aux.Tienda.Nombre+" \\n "
					nodoLista += " Calificacion: "+strconv.Itoa(aux.Tienda.Calificacion)+"\"];"
					if aux.Siguiente != nil{
						enlaceLista += strconv.Itoa(aux.IdNodo)+ "->" + strconv.Itoa(aux.Siguiente.IdNodo)+";\n"
						enlaceLista += strconv.Itoa(aux.Siguiente.IdNodo)+ "->" + strconv.Itoa(aux.IdNodo)+";\n"
					}
					if aux == vect.Primero{
						enlaceLista += "vec:f"+strconv.Itoa(i+1)+" -> " + strconv.Itoa(aux.IdNodo)+";\n"
					}
					aux = aux.Siguiente
				}
			}
		}
	}
	acum += nodo + enlace + nodoLista + enlaceLista + "\n}\n"

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

	_, err = file.WriteString(acum)
	if existeError(err) {
		return
	}

	err = file.Sync()
	if existeError(err) {
		return
	}

	fmt.Println("Archivo actualizado existosamente.")

	path2, _ := exec.LookPath("dot")
	cmd, _ := exec.Command(path2, "-Tpng","grafo.dot").Output()
	mode := int(0777)
	ioutil.WriteFile("grafo.png",cmd, os.FileMode(mode))
}

func existeError(err error) bool{
	if err != nil{
		fmt.Println(err.Error())
	}
	return (err != nil)
}

func busquedaEspecifica(w http.ResponseWriter, r *http.Request){
	var busqueda busqueda
	consulta, err := ioutil.ReadAll(r.Body)
	if err != nil{
		fmt.Fprintf(w, "Entrada invalida, intente nuevamente")
	}
	json.Unmarshal(consulta, &busqueda)
	for _, vect := range vector{
		if vect.Departamento == busqueda.Departamento && vect.Identificador == strconv.Itoa(busqueda.Calificacion){
			if vect.Primero != nil{
				aux := vect.Primero
				for aux != nil{
					if aux.Tienda.Nombre == busqueda.Nombre{
						w.Header().Set("Content-Type", "application/json")
						json.NewEncoder(w).Encode(aux.Tienda)
						break
					}
					aux = aux.Siguiente
				}
				if aux == nil{
					fmt.Fprintf(w, "------------TIENDA NO ENCONTRADA------------")
				}
			}
		}
	}
}

func busquedaDePosicion(w http.ResponseWriter, r *http.Request)  {
	variables := mux.Vars(r)
	idPosicion, err := strconv.Atoi(variables["id"])
	if err != nil{
		fmt.Fprintf(w, "ERROR// Id inválido")
	}
	if idPosicion > len(vector){
		fmt.Fprintf(w, "-----------EL VECTOR NO TIENE TANTAS POSICIONES-----------")
	}else {
		for i, vect := range vector{
			if i == idPosicion - 1{
				if vect.Primero != nil{
					aux := vect.Primero
					for aux != nil{
						w.Header().Set("Content-Type", "application/json")
						json.NewEncoder(w).Encode(aux.Tienda)
						aux = aux.Siguiente
					}
				}else {
					fmt.Fprintf(w, "-------------NO SE ENCUENTRAN TIENDAS EN ESTA POSICION----------")
				}
			}
		}
	}
}

func eliminar(w http.ResponseWriter, r *http.Request){
	rIndice := 0
	rDepa := 0
	rCalif := 0
	rPosi := 0
	var busqueda busqueda
	consulta, err := ioutil.ReadAll(r.Body)
	if err != nil{
		fmt.Fprintf(w, "--------------ENTRADA INVALIDA, INTENTE NUEVAMETE--------------")
	}
	json.Unmarshal(consulta, &busqueda)
	for _, vect := range vector{
		contadorPosicion := 0
		if vect.Departamento == busqueda.Departamento && vect.Identificador == strconv.Itoa(busqueda.Calificacion){
			if vect.Primero != nil{
				aux := vect.Primero
				for aux != nil{
					contadorPosicion += 1
					for i, indice := range lsIndex{
						if indice == vect.Indice{
							rIndice = i + 1
						}
					}
					for i, dep := range lsDep{
						if dep == busqueda.Departamento{
							rDepa = i + 1
						}
					}
					if aux.Tienda.Nombre == busqueda.Nombre{
						rCalif = aux.Tienda.Calificacion
						rPosi = contadorPosicion
						if aux.Siguiente == nil{
							if aux == vect.Primero{
								vect.Primero = nil
								respuesta := strconv.Itoa(rIndice) +", "+ strconv.Itoa(rDepa) +", "+strconv.Itoa(rCalif)+", "+strconv.Itoa(rPosi)
								fmt.Fprintf(w, "Tienda elimincada carrectamente ( %v )", respuesta)
								break
							}else {
								aux.Anterior.Siguiente = nil
								respuesta := strconv.Itoa(rIndice) +", "+ strconv.Itoa(rDepa) +", "+strconv.Itoa(rCalif)+", "+strconv.Itoa(rPosi)
								fmt.Fprintf(w, "Tienda elimincada carrectamente ( %v )", respuesta)
								break
							}
						}else if aux == vect.Primero{
							aux.Siguiente.Anterior = nil
							vect.Primero = aux.Siguiente
							respuesta := strconv.Itoa(rIndice) +", "+ strconv.Itoa(rDepa) +", "+strconv.Itoa(rCalif)+", "+strconv.Itoa(rPosi)
							fmt.Fprintf(w, "Tienda elimincada carrectamente ( %v )", respuesta)
							break
						}else {
							aux.Anterior.Siguiente = aux.Siguiente
							aux.Siguiente.Anterior = aux.Anterior
							respuesta := strconv.Itoa(rIndice) +", "+ strconv.Itoa(rDepa) +", "+strconv.Itoa(rCalif)+", "+strconv.Itoa(rPosi)
							fmt.Fprintf(w, "Tienda elimincada carrectamente ( %v )", respuesta)
							break
						}
					}
					aux = aux.Siguiente
				}
			}
		}
	}
}

func guardar(w http.ResponseWriter, r *http.Request)  {
	type tienda []Datos.Tienda //definir tipo de dato de arreglo de tiendas
	//var t1 Datos.Tienda
	//t1.Nombre = "Nombre1"
	//t1.Calificacion = 3
	//t1.Contacto = "1217389721"
	//t1.Descripcion = "Descripcion de prueba"

	//tiendas = append(tiendas, t1)
	//tiendas = append(tiendas, t1)

	type depa []Datos.Departamento  //definir tipo de dato arreglo de departamentos
	//var dep Datos.Departamento
	//dep.Nombre = "Depa prueb1"
	//dep.Tiendas = tiendas

	//departamentos = append(departamentos, dep)
	//departamentos = append(departamentos, dep)

	type dato []Datos.Datos
	var datos = dato{} //arreglo de datos guardados (Inice, departamentos[])

	//var d Datos.Datos
	//d.Indice = "Prueba1"
	//d.Departamentos = departamentos
	//datos = append(datos, d)
	//datos = append(datos, d)

	for _, indice := range lsIndex{
		var nuevoDato Datos.Datos // (indice y departamentos[])
		var departamentos = depa{}
		nuevoDato.Indice = indice
		for _, depa := range lsDep{
			var tiendas = tienda{} //nuevo listado de tiendas para cada depa
			var nuevoDepa Datos.Departamento //(nombre y Tiendas[])
			for _, vect := range vector{
				if vect.Indice == indice && vect.Departamento == depa{
					if vect.Primero != nil{
						aux := vect.Primero
						for aux != nil{
							var nuevaTienda Datos.Tienda
							nuevaTienda.Nombre = aux.Tienda.Nombre
							nuevaTienda.Descripcion = aux.Tienda.Descripcion
							nuevaTienda.Contacto = aux.Tienda.Contacto
							nuevaTienda.Calificacion = aux.Tienda.Calificacion
							tiendas = append(tiendas, nuevaTienda)  //guarar tienda en listado de tiendas
							aux = aux.Siguiente
						}
					}
					nuevoDepa.Nombre = depa
				}
				nuevoDepa.Tiendas = tiendas
			}
			departamentos = append(departamentos, nuevoDepa) //listado de departamentos para cada indice
			nuevoDato.Departamentos = departamentos
		}
		datos = append(datos, nuevoDato)
	}

	ult := Datos.EstructuraGeneral{
		Datos: datos,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ult)


}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", indexRoute)
	router.HandleFunc("/cargartienda", cargartienda).Methods("POST")
	router.HandleFunc("/getArreglo", generarGrafo).Methods("GET")
	router.HandleFunc("/TiendaEspecifica", busquedaEspecifica).Methods("POST")
	router.HandleFunc("/id/{id}", busquedaDePosicion).Methods("GET")
	router.HandleFunc("/Eliminar", eliminar).Methods("DELETE")
	router.HandleFunc("/guardar", guardar).Methods("GET")
	router.HandleFunc("/cargarInventarios", cargarInventario).Methods("POST")

	log.Fatal(http.ListenAndServe(":3000", router))

}

func indexRoute(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "Welcome to my apisss22")
}
