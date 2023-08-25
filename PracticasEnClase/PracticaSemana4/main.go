package main

import "fmt"

type aula int

const (
	Aula0 aula = iota
	Aula1
	Aula2
	Aula3
	Aula4
)

type curso struct {
	nombre string
	dia    string
	hora   string
	aula   aula
	notas  []float64 // Agregamos un slice de notas
}

type listaCursos []curso

var lCursos listaCursos

func (l *listaCursos) verificarCurso(dia string, hora string, aula aula) bool {
	for _, cur := range *l {
		if cur.dia == dia && cur.hora == hora {
			// Si el aula es diferente, no hay choque de horario
			if cur.aula != aula {
				return true
			}
			// Hay choque de horario (no se puede agregar)
			return false
		}
	}
	// Se puede agregar (no hay cursos en ese horario)
	return true
}

func (l *listaCursos) agregarCurso(nombre string, dia string, hora string, aula aula) bool {
	cursoVerificado := l.verificarCurso(dia, hora, aula)

	if cursoVerificado == false {
		fmt.Println("No es posible agregar curos, hay choque de horario")
		return false
	}
	*l = append(*l, curso{nombre: nombre, dia: dia, hora: hora, aula: aula})
	return true
}

func llenarCursos() {
	lCursos.agregarCurso("Lenguajes de programacion", "Lunes", "Mañana", Aula0)
	lCursos.agregarCurso("Bases de datos I", "Martes", "Tarde", Aula1)
	lCursos.agregarCurso("Compiladores", "Martes", "Mañana", Aula1)
	lCursos.agregarCurso("Requerimientos de Software", "Viernes", "Mañana", Aula2)
}

func agregarNotasACurso(cursoIdx int, nuevasNotas []float64) {
	lCursos[cursoIdx].notas = append(lCursos[cursoIdx].notas, nuevasNotas...)
}

func aulaToString(a aula) string {
	switch a {
	case Aula0:
		return "Aula 0"
	case Aula1:
		return "Aula 1"
	case Aula2:
		return "Aula 2"
	case Aula3:
		return "Aula 3"
	case Aula4:
		return "Aula 4"
	default:
		return "Desconocido"
	}
}

func map1[P1, P2 any](list []P1, f func(P1) P2) []P2 {
	mapped := make([]P2, len(list))

	for i, e := range list {
		mapped[i] = f(e)
	}
	return mapped
}

func filter1[P1 any](list []P1, f func(P1) bool) []P1 {
	filtered := make([]P1, 0)

	for _, element := range list {
		if f(element) {
			filtered = append(filtered, element)
		}
	}
	return filtered
}

func cantidadAprobadosReprobados(notas []float64) (int, int) {
	// Filtrar el slice de notas para obtener solo las notas aprobadas
	notasAprobadas := filter1(notas, func(n float64) bool {
		return n >= 70 // 70 es la nota mínima para aprobar
	})

	// Contar la cantidad de elementos en el slice de notas aprobadas
	cantidadAprobados := len(notasAprobadas)

	cantidadReprobados := len(notas) - cantidadAprobados

	// Devolver la cantidad de estudiantes que aprobaron
	return cantidadAprobados, cantidadReprobados
}

func promedioNotas(notas []float64) float64 {
	if len(notas) == 0 {
		return 0.0
	}

	sumaNotas := 0.0
	for _, nota := range notas {
		sumaNotas += nota
	}
	return sumaNotas / float64(len(notas))
}

func main() {
	llenarCursos()
	fmt.Println(lCursos)

	// Ejemplo si choca el horario
	lCursos.agregarCurso("Bases de Datos II", "Martes", "Mañana", Aula1)

	// Ejemplo si no choca
	lCursos.agregarCurso("Bases de Datos II", "Martes", "Mañana", Aula4)
	fmt.Println(lCursos)

	cursoIdx := 0                                 // Índice del curso al que deseas agregar notas
	nuevasNotas := []float64{20, 30, 90, 100, 70} // Nuevas notas a agregar
	agregarNotasACurso(cursoIdx, nuevasNotas)

	aprobados, reprobados := cantidadAprobadosReprobados(lCursos[0].notas)
	fmt.Printf("La cantidad de aprobados es: %d\n", aprobados)
	fmt.Printf("La cantidad de reprobados es: %d\n", reprobados)
	promedio := promedioNotas(lCursos[0].notas)
	fmt.Printf("El promedio es: %.2f\n", promedio)

}
