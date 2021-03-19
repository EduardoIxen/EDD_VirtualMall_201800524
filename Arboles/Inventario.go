package Arboles

type ContInventarios struct {
	Inventarios []Inventario `json:"Inventarios"`
}
type Producto struct {
	Nombre      string `json:"Nombre"`
	Codigo      int    `json:"Codigo"`
	Descripcion string `json:"Descripcion"`
	Precio      float64 `json:"Precio"`
	Cantidad    int    `json:"Cantidad"`
	Imagen      string `json:"Imagen"`
}

func NewProduct(nombre string, codigo int, descripcion string, precio float64, cantidad int, imagen string) *Producto {
	return &Producto{nombre, codigo, descripcion, precio, cantidad, imagen}
}

type Inventario struct {
	Tienda       string     `json:"Tienda"`
	Departamento string     `json:"Departamento"`
	Calificacion int        `json:"Calificacion"`
	Productos    []Producto `json:"Productos"`
}