/**
  IT0425 Introducción a la inteligencia artifical
  slide_bfs.go
  Propósito: Resolver el cubo mágico usando el algoritmo DFS.

  @autor Brando Pérez Pacheco
  @autor Uriel Tzel Ciau
  @version 1.1 03/09/14
*/

package main

import (
	"fmt"
	"math"
)

type Stack []int
type Matrix [][]int
type Graph map[int][]int

// valores del sliding puzzle
type Puzzle struct {
	Mat Matrix
	Id  int
}

// valores del algoritmo DFS
type AttrDfs struct {
	Path  Stack
	Vcted Stack
}

// valores del sliding puzzle
type PuzzleSolver struct {
	End      Puzzle
	Start    Puzzle
	IdPuzzle map[int]Puzzle
	Searched Stack // "puzzles" pintados
	AttrDfs
	Graph
}

// función principal
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
	p.Dfs(p.Start.Id, p.End.Id)
	p.printPath()
}

/**
  Introduce un elemento al final del slice (arreglo).

  @param valor o valores enteros.

  @return nada.
*/
func (self *Stack) push(values ...int) {
	for _, val := range values {
		*self = append((*self)[0:], append([]int{val}, (*self)[:0]...)...)
	}
}

/**
  Elimina un elemento atrás del slice.

  @param nada.

  @return entero que fue sacado.
*/
func (self *Stack) pop() (val int) {
	val, *self = (*self)[len(*self)-1], (*self)[:len(*self)-1]

	return
}

/**
  Verifica si un elemento se encuentra en el slice.

  @param entero que deseamos encontrar.

  @return boleano que nos indica si encontró o no el elemento.
*/
func (self *Stack) itemNotIn(value int) bool {
	for _, val := range *self {
		if val == value {
			return false
		}
	}

	return true
}

/**
  Copia los elementos de una matriz a otra.

  @param nada.

  @return matriz copia.
*/
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

	return swaped
}

// método inicializador del type PuzzleSolver
func NewPuzzleSolver(start, end Matrix) *PuzzleSolver {
	self := new(PuzzleSolver)

	self.End = Puzzle{
		Mat: end,
		Id:  end.getId(),
	}

	self.Start = Puzzle{
		Mat: start,
		Id:  start.getId(),
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

  @param nada.

  @return nada.
*/
func (self *PuzzleSolver) printPath() {
	for i := 0; i < len(self.Path); i++ {
		mat := self.IdPuzzle[self.Path[i]]
		mat.printMat()
	}
}

/**
  Método que realiza el algoritmo BFS.

  @param nada.

  @return boleano que indica éxito o fracaso.
*/
func (self *PuzzleSolver) Dfs(start, end int) {
	self.AttrDfs.Vcted.push(start) // "pintar" actual

	if start == end { // si es el objetivo: termina
		self.AttrDfs.Path.push(self.AttrDfs.Vcted...)
	}

	self.moveSpaceGetChild(self.IdPuzzle[start]) // generar hijos

	for _, v := range self.Graph[start] {
		if self.AttrDfs.Vcted.itemNotIn(v) { // sino han sido visitados
			self.Dfs(v, end)
		}
	}
}
