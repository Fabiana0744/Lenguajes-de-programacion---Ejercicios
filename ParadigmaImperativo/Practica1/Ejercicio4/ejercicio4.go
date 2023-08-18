package main

import "fmt"

// EJERCICIO 4

type calzado struct {
	modelo string
	precio int
	talla  int
	stock  int
}
type listaCalzados []calzado

var lCalzados listaCalzados

func (l *listaCalzados) buscarCalzado(modelo string, talla int) (*calzado, int) {
	var error = 0
	for i := 0; i < len(*l); i++ {
		if (*l)[i].modelo == modelo && (*l)[i].talla == talla {
			return &(*l)[i], error
		}
	}
	error = -1
	return nil, error
}

func (l *listaCalzados) agregarCalzado(modelo string, precio int, talla int, stock int) {

	if talla < 34 || talla > 44 {
		fmt.Println("Talla fuera del rango permitido (34-44).")
		return
	}

	var calz, error = l.buscarCalzado(modelo, talla)

	if error != -1 {
		calz.stock += stock
		//fmt.Println("El producto ya existe con la misma talla.")
		return
	}

	*l = append(*l, calzado{modelo: modelo, precio: precio, talla: talla, stock: stock})

}

func (l *listaCalzados) eliminarCalzado(modelo string, talla int) {
	var indiceAEliminar = -1
	for i, calz := range *l {
		if calz.modelo == modelo && calz.talla == talla {
			indiceAEliminar = i
			break
		}
	}

	if indiceAEliminar != -1 {
		if indiceAEliminar == len(*l)-1 { // Si el elemento a eliminar es el último
			*l = (*l)[:indiceAEliminar] // Eliminar simplemente ajustando la longitud
		} else {
			*l = append((*l)[:indiceAEliminar], (*l)[indiceAEliminar+1:]...)
		}
	}
}

func (l *listaCalzados) venderCalzado(modelo string, talla int, cantidad int) {
	var calz, error = l.buscarCalzado(modelo, talla)
	if error != -1 {
		// Validacion
		if cantidad > calz.stock {
			fmt.Printf("No hay cantidad suficiente de '%s' para vender.\n", modelo)
			return
		}
		calz.stock = calz.stock - cantidad

		if calz.stock <= 0 {
			l.eliminarCalzado(modelo, talla)
			fmt.Println("Producto vendido y agotado")
			return
		} else {
			fmt.Println("Producto vendido")
			return
		}
	}
	fmt.Println("No existe el producto")
}

func llenaInventario() {
	lCalzados.agregarCalzado("Adidas", 45000, 40, 10)
	lCalzados.agregarCalzado("Nike", 80000, 38, 20)
	lCalzados.agregarCalzado("Puma", 45000, 39, 10)
	lCalzados.agregarCalzado("New Balance", 60000, 42, 15)
	lCalzados.agregarCalzado("Vans", 40000, 44, 6)
}

func main() {
	llenaInventario()
	fmt.Println(lCalzados)

	lCalzados.agregarCalzado("Adidas", 4500, 40, 12)
	fmt.Println(lCalzados)

	lCalzados.venderCalzado("Nike", 38, 5)
	fmt.Println(lCalzados)

}
