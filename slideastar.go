/**
  IT0425 Introducción a la inteligencia artifical
  slide_bfs.go
  Propósito: Resolver el cubo mágico usando el algoritmo A*.

  @autor Brando Pérez Pacheco
  @autor Uriel Tzel Ciau
  @version 1.1 03/09/14
*/

package main

import (
	"fmt"
	"math"
	"sort" // quicksort
)

type Queue []int
type Matrix [][]int
type Graph map[int][]int
type FVal []Eva // slice de todos los valores F de los puzzles

type Eva struct {
	Id int // ID del slide
	F  int // f = g(u) + h(u)
}

// valores para el puzzle
type Puzzle struct {
	Mat Matrix
	Id  int
	Heu int // valor de huerística
	Lvl int // profundidad
}

// valores del resolvedor de sliding puzzle
type PuzzleSolver struct {
	End      Puzzle
	Start    Puzzle
	IdPuzzle map[int]Puzzle
	Searched Queue
	Graph
}

func main() {
	start := Matrix{
		{5, 1, 6},
		{4, 0, 7},
		{8, 3, 2},
	}

	end := Matrix{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 0},
	}

	p := NewPuzzleSolver(start, end)

	p.AStar()
}

/**
  Introduce un elemento al inicio del slice (arreglo).

  @param valor o valores enteros.

  @return nada.
*/
func (self *Queue) push(values ...int) {
	for _, val := range values {
		*self = append((*self)[:0], append([]int{val}, (*self)[0:]...)...)
	}
}

/**
  Elimina un elemento atrás del slice.

  @param nada.

  @return entero que fue sacado.
*/
func (self *Queue) pop() (val int) {
	val, *self = (*self)[len(*self)-1], (*self)[:len(*self)-1]

	return
}

/**
  Elimina un elemento del slice a partir de su valor.

  @param elemento que deseamos eliminar.

  @return el elemento eliminado.
*/
func (self *Queue) deleteSearch(indx int) (val int) {
	for i, v := range *self {
		if v == indx {
			val, (*self)[i], (*self) =
				(*self)[i], (*self)[len((*self))-1], (*self)[:len((*self))-1]

			return
		}
	}

	return
}

/**
  Verifica si un elemento se encuentra en el slice.

  @param entero que deseamos encontrar.

  @return boleano que nos indica si encontró o no el elemento.
*/
func (self *Queue) itemNotIn(value int) bool {
	for _, val := range *self {
		if val == value {
			return false
		}
	}

	return true
}

/**
  Regresa el tamaño de un type.

  @param nada.

  @return tamaño.
*/
func (self FVal) Len() int { return len(self) }

/**
  Realiza intercambio de valores.

  @param los índices.

  @return nada
*/
func (self FVal) Swap(i, j int) { self[i], self[j] = self[j], self[i] }

/**
  Realiza intercambio de valores.

  @param los índices.

  @return nada.
*/
func (self FVal) Less(i, j int) bool { return self[i].F > self[j].F }

func (self *Matrix) copy() Matrix {
	mat_cp := make(Matrix, 3, 3)

	for i := 0; i < len((*self)[0]); i++ {
		mat_cp[i] = append(mat_cp[i], (*self)[i]...)
	}

	return mat_cp
}

/**
  Obtiene el identificador único para la matriz.

  @param nada.

  @return número entero que es el identificador.
*/
func (self *Matrix) getId() int {
	pwr, sum := 8, 0

	for i := 0; i < len(*self); i++ {
		for j := 0; j < len(*self); j++ {
			if (*self)[i][j] != 0 {
				sum += (*self)[i][j] * int(math.Pow(float64(10), float64(pwr)))
			}
			pwr--
		}
	}

	return sum
}

/**
  Imprime bonito una matriz.

  @param nada.

  @return nada.
*/
func (self *Puzzle) printMat() {
	fmt.Printf("\n")
	for i := 0; i < len(self.Mat); i++ {
		for j := 0; j < len(self.Mat); j++ {
			if self.Mat[i][j] != 0 {
				fmt.Printf("[%d]", self.Mat[i][j])
				continue
			}
			fmt.Printf("[X]")
		}
		fmt.Printf("\n")
	}
}

/**
  Busca el espacio en la matriz (0).

  @param nada.

  @return enteros que son las coordenadas.
*/
func (self *Puzzle) getPositionSpace() (i, j int) {
	for i := 0; i < len(self.Mat); i++ {
		for j := 0; j < len(self.Mat); j++ {
			if self.Mat[i][j] == 0 {
				return i, j
			}
		}
	}

	return i, j
}

/**
  Obtiene la heurística de una matriz a partir de que sus posiciones sean
  incorrectas.

  @param nada.

  @return valor entero de la heurística.
*/
func (self *Matrix) getHeu() int {
	count := 0
	good := Matrix{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 0},
	}

	for i := 0; i < len((*self)[0]); i++ {
		for j := 0; j < len((*self)[0]); j++ {
			if (*self)[i][j] != 0 && (*self)[i][j] != good[i][j] {
				count++
			}
		}
	}

	return count
}

/**
  Obtiene los hijos apartir de una matriz padre. Intercambiando dos elementos
  dentro de la matriz del type Puzzle y encuentra el ID de la matriz modificada.

  @param posición del espacio (y, x).
  @param posición destino espacio (y_des, x_des).

  @return type Puzzle.
*/
func (self *Puzzle) getChild(y, x, y_des, x_des int) Puzzle {
	swaped := Puzzle{Mat: self.Mat.copy()}

	swaped.Mat[y][x], swaped.Mat[y_des][x_des] =
		swaped.Mat[y_des][x_des], swaped.Mat[y][x]

	swaped.Id = swaped.Mat.getId() // obteniendo ID
	swaped.Heu = self.Mat.getHeu() // heurísitica
	swaped.Lvl = self.Lvl + 1      // nivel de profundidad

	return swaped
}

// método inicializador del type PuzzleSolver
func NewPuzzleSolver(start, end Matrix) *PuzzleSolver {
	self := new(PuzzleSolver)

	self.End = Puzzle{
		Mat: end,
		Id:  end.getId(),
		Heu: 0,
		Lvl: 0,
	}

	self.Start = Puzzle{
		Mat: start,
		Id:  start.getId(),
		Heu: start.getHeu(),
		Lvl: 0,
	}

	self.IdPuzzle = make(map[int]Puzzle)
	self.IdPuzzle[self.Start.Id] = Puzzle{Mat: self.Start.Mat}

	self.Graph = make(Graph)

	return self

}

/**
  Encuentra los hijos de la matriz padre y verifica que no se hayan sido
  visitados.

  @param slide puzzle padre.

  @return nada.
*/
func (self *PuzzleSolver) moveSpaceGetChild(puz Puzzle) {
	y, x := puz.getPositionSpace()
	y_pos := []int{y - 1, y, y + 1, y}
	x_pos := []int{x, x + 1, x, x - 1}
	sum := []Puzzle{}

	for i := 0; i < 4; i++ {
		if y_pos[i] > -1 && y_pos[i] < 3 && x_pos[i] > -1 && x_pos[i] < 3 {
			sum = append(sum, puz.getChild(y, x, y_pos[i], x_pos[i]))
		}
	}

	for _, val := range sum {

		// verificando si los hijos no han sido visitados
		if self.Searched.itemNotIn(val.Id) {
			self.Searched.push(val.Id)
			self.Graph[puz.Mat.getId()] =
				append(self.Graph[puz.Mat.getId()], val.Id)
			self.IdPuzzle[val.Id] = val
		}
	}
}

/**
  Imprime pasos para resolver el slide puzzle.

  @param mapa que contiene los padres de los hijos.

  @return nada.
*/
func (self *PuzzleSolver) printPath(path map[int]int, start, end int) {
	paint := Queue{}

	paint.push(end)

	for i := path[paint[0]]; paint[0] != start; i = path[paint[0]] {
		paint.push(i)
	}

	for i := 0; i < len(paint); i++ {
		mat := self.IdPuzzle[paint[i]]
		mat.printMat()
	}
}

/**
  Método que realiza el algoritmo A*.

  @param nada.

  @return boleano que indica éxito o fracaso.
*/
func (self *PuzzleSolver) AStar() bool {
	open := Queue{}
	close := Queue{}
	path := map[int]int{}
	g := map[int]int{} // heurística de los slide
	f := FVal{}        // ID y valores de F de los slide

	open.push(self.Start.Id)

	// distancia al padre (g) y valor de f = g + h
	g[self.Start.Id] = 0
	f = append(f,
		Eva{Id: self.Start.Id, F: g[self.Start.Id] + self.Start.Heu})

	for len(open) > 0 { // mientras no esté vacío
		sort.Sort(f) // aplicando quicksort

		v := open.deleteSearch(f[len(f)-1].Id) // encontrando el mejor valor
		f = f[:len(f)-1]                       // eliminado del arreglo de F's

		if v == self.End.Id { // si es el objetivo: termina
			self.printPath(path, self.Start.Id, self.End.Id)
			return true
		}

		close.push(v) // marcar como nodo cerrado el actual

		self.moveSpaceGetChild(self.IdPuzzle[v]) // generando hijos

		for _, u := range self.Graph[v] { // por cada hijo
			if !close.itemNotIn(u) { // si está cerrado
				continue
			}

			tg := g[v] + self.IdPuzzle[u].Lvl // profundidad

			// nodo abierto o menor profundidad que nodo actual
			if open.itemNotIn(u) || tg < g[u] {

				// apadrinar, nuevo nivel de profundidad y f
				path[u] = v
				g[u] = tg
				f = append(f,
					Eva{Id: u, F: g[u] + self.Start.Heu})
				if open.itemNotIn(u) { // sino está abierto
					open.push(u) // ábrelo
				}
			}
		}
	}

	return false
}
