package main

import (
	"fmt"
	"sort"
)

type producto struct {
	nombre   string
	cantidad int
	precio   int
}
type listaProductos []producto

var lProductos listaProductos

const existenciaMinima int = 10

func (l *listaProductos) agregarProducto(nombre string, cantidad int, precio int) {

	for i := 0; i < len(*l); i++ { // for _, prod := range *l {
		if (*l)[i].nombre == nombre {
			(*l)[i].cantidad += cantidad
			if (*l)[i].precio != precio {
				(*l)[i].precio = precio
			}
			return
		}
	}
	*l = append(*l, producto{nombre: nombre, cantidad: cantidad, precio: precio})
}

func (l *listaProductos) agregarProductos(productos ...producto) {
	for _, prod := range productos {
		l.agregarProducto(prod.nombre, prod.cantidad, prod.precio)
	}
}

func (l *listaProductos) buscarProducto(nombre string) (*producto, int) { //el retorno es el índice del producto encontrado y -1 si no existe
	var error = 0
	var i int
	for i = 0; i < len(*l); i++ {
		if (*l)[i].nombre == nombre {
			return &(*l)[i], error
		}
	}
	//Si no se encuentra
	error = -1
	return nil, error
}

func (l *listaProductos) eliminarProducto(lista *listaProductos, nombre string) {
	var indiceAEliminar = -1
	for i, prod := range *lista {
		if prod.nombre == nombre {
			indiceAEliminar = i
			break
		}
	}

	if indiceAEliminar != -1 {
		if indiceAEliminar == len(*lista)-1 { // Si el elemento a eliminar es el último
			*lista = (*lista)[:indiceAEliminar] // Eliminar simplemente ajustando la longitud
		} else {
			*l = append((*lista)[:indiceAEliminar], (*lista)[indiceAEliminar+1:]...)
		}
	}
}

func (l *listaProductos) venderProducto(nombre string, cantidad int) {
	var prod, error = (*l).buscarProducto(nombre)
	if error != -1 {
		// Validacion
		if cantidad > prod.cantidad {
			fmt.Printf("No hay cantidad suficiente de '%s' para vender.\n", nombre)
			return
		}
		prod.cantidad -= cantidad

		if prod.cantidad <= 0 {
			l.eliminarProducto(l, nombre)
			fmt.Println("Producto vendido y agotado")
			return
		} else {
			fmt.Println("Producto vendido")
			return
		}
	}
	fmt.Println("No existe el producto")
}

func llenarDatos() {
	lProductos.agregarProducto("arroz", 15, 2500)
	lProductos.agregarProducto("frijoles", 4, 2000)
	lProductos.agregarProducto("leche", 8, 1200)
	lProductos.agregarProducto("café", 12, 4500)

}

func (l *listaProductos) listarProductosMínimos() listaProductos {
	productosMinimos := make(listaProductos, 0)
	for _, prod := range *l {
		if prod.cantidad < existenciaMinima {
			productosMinimos = append(productosMinimos, prod)
		}
	}
	//fmt.Println("Productos minimos:")
	//fmt.Println(newSlice)
	return productosMinimos
}

func (l *listaProductos) aumentarInventarioDeMinimos(listaMinimos listaProductos) {
	for _, prod := range listaMinimos {
		diferencia := existenciaMinima - prod.cantidad
		if diferencia > 0 {
			l.agregarProducto(prod.nombre, diferencia, prod.precio)
		}
	}
}

func (l *listaProductos) ordenarPorNombre() {
	sort.SliceStable(*l, func(i, j int) bool {
		return (*l)[i].nombre < (*l)[j].nombre
	})
}

func main() {
	llenarDatos()
	fmt.Println(lProductos)

	//venda productos
	lProductos.venderProducto("arroz", 15)
	//lProductos.venderProducto("frijoles", 4)
	fmt.Println(lProductos)

	//lProductos.agregarProducto("arroz", 20, 2500)
	//fmt.Println(lProductos)

	//lProductos.buscarProducto("arroz")

	minimos := lProductos.listarProductosMínimos()
	fmt.Printf("Productos mínimos son: %v\n", minimos)

	lProductos.aumentarInventarioDeMinimos(minimos)
	fmt.Printf("Productos actualizados: %v\n", lProductos)

	lProductos.ordenarPorNombre()
	fmt.Printf("Productos ordenados: %v\n", lProductos)
}
