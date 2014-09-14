package main // invocando el paquete principal

import (
	// importando bibliotecas que se usan
	"container/list"
	"fmt"
	"math"
)

type Matrix [][]int     // slice de slices
type Parent map[int]int // map
type Array []int        // slice

/**
  Encuentra si existe un elemento dentro de un arreglo.

  @param el valor a buscar.

  @return valor boleano que indica si lo encontró.
*/
func (self *Array) ItemNotIn(val int) bool {
	for _, value := range *self {
		if val == value {
			return false
		}
	}

	return true
}

/**
  Type que describe a una cola.
*/
type Queue struct {
	queue list.List
}

/**
  Inicializador del type Queue.

  @param nada.

  @return nada.
*/
func NewQueue() *Queue {
	self := new(Queue)

	return self
}

/**
  Mete en la cola, al principio, un número indeterminado de valores.

  @param uno o varios enteros.

  @return nada.
*/
func (self *Queue) Push(val ...int) {
	for _, val := range val {
		self.queue.PushFront(val)
	}
}

/**
  Elimina un elemento atrás de la cola.

  @param nada.

  @return nada.
*/
func (self *Queue) Pop() int {
	e := self.queue.Back()

	self.queue.Remove(e)

	return e.Value.(int)
}

func (self *Queue) PrintQueue() {
	fmt.Printf("[")

	for e := self.queue.Front(); e != nil; e = e.Next() {
		fmt.Printf(" %d ", e.Value.(int))
	}

	fmt.Printf("]\n")
}

/**
  Contiene las características principales del rompecabezas slide.
*/
type SlidePuzzle struct {
	Width  int
	Height int
	Mat    Matrix
}

/**
  Imprime un matriz (slice de slices) bonito.

  @param matriz de enteros a imprimir.

  @return nada.
*/
func (self *SlidePuzzle) PrintMat(mat Matrix) {
	fmt.Printf("\n")
	for i := 0; i < self.Height; i++ {
		for j := 0; j < self.Width; j++ {
			if mat[i][j] != 0 {
				fmt.Printf("[%d]", mat[i][j])
			} else {
				fmt.Printf("[X]")
			}
		}
		fmt.Printf("\n")
	}
}

/**
  Calcula el ID de una matriz.

  @param matriz de enteros.

  @return ID de la matriz.
*/
func (self *SlidePuzzle) SumMat(mat Matrix) int {
	pwr, sum := 8, 0

	for i := 0; i < self.Height; i++ {
		for j := 0; j < self.Width; j++ {
			if mat[i][j] != 0 {
				sum += mat[i][j] * int(math.Pow(float64(10), float64(pwr)))
			}
			pwr--
		}
	}

	return sum
}

/**
  Describe las características del rempecabezas slide solucionado. Embebe del
  type padre: SlidePuzzle.
*/
type SlidePuzzleSolution struct {
	SlidePuzzle
	Sum int
}

/**
  Inicializador del type SlidePuzzleSolution.

  @param nada.

  @return type SlidePuzzleSolution inicializado.
*/
func NewSlidePuzzleSolution() *SlidePuzzleSolution {
	self := new(SlidePuzzleSolution)

	self.Width = 3
	self.Height = 3
	self.Mat = Matrix{
		{8, 7, 6},
		{5, 4, 3},
		{2, 1, 0},
	}
	self.Sum = self.SumMat(self.Mat)

	return self
}

/**
  Describe las características del puzzle slide desordenado. Embebe del type
  SlidePuzzleSolution.
*/
type SlidePuzzleSolvr struct {
	SlidePuzzle
	Ids Array
}

/**
  Inicializador del type SlidePuzzleSolvr.

  @param nada.
  @param type SlidePuzzleSolvr inicializado.

  @return nada.
*/
func NewSlidePuzzleSolvr() *SlidePuzzleSolvr {
	self := new(SlidePuzzleSolvr)

	self.Width = 3
	self.Height = 3
	self.Mat = Matrix{
		{5, 8, 2},
		{1, 0, 3},
		{4, 6, 7},
	}

	self.Ids = append(self.Ids, self.SumMat(self.Mat))

	return self
}

/**
  Método que regresa la posición del espacio (0).

  @param matriz de enteros que se desea analizar.

  @return posición del espacio (i, j).
*/
func (self *SlidePuzzleSolvr) GetPositionSpace(mat Matrix) (i, j int) {
	for i := 0; i < self.Height; i++ {
		for j := 0; j < self.Width; j++ {
			if mat[i][j] == 0 {
				return i, j
			}
		}
	}
	return i, j
}

/**
  Copia el contenido de una matriz a otra.

  @param matriz de enteros que se desea copiar.

  @return matriz copia de enteros.
*/
func (self *SlidePuzzleSolvr) Copy(mat Matrix) (mat_cp Matrix) {
	mat_cp = make(Matrix, 3, 3)

	for i := 0; i < len(mat[0]); i++ {
		mat_cp[i] = append(mat_cp[i], mat[i]...)
	}

	return
}

/**
  Intercambia el contenido de dos celdas en una matriz.

  @param matriz de enteros que se copiará.
  @param posición de origen (y, x).
  @param posición destino (y_des, x_des).

  @return matriz de enteros con el contenido de dos celdas intercambiadas.
*/
func (self *SlidePuzzleSolvr) Swap(mat Matrix, y, x, y_des, x_des int) (mat_des Matrix) {
	mat_des = self.Copy(mat)

	mat_des[y][x], mat_des[y_des][x_des] = mat_des[y_des][x_des], mat_des[y][x]

	return
}

/**
  Mueve el espacio (0) en la matriz (si se puede) y crea hijos de la matriz de
  padre verificando que no haya elementos repetidos calculando el indentificador
  del hijo.

  @param matriz de enteros que se desea analizar.

  @return nada.
*/
func (self *SlidePuzzleSolvr) MoveSpace(mat Matrix) {
	y, x := self.GetPositionSpace(mat)
	sum := Array{}

	if y-1 > -1 {
		mat_temp := self.Swap(mat, y, x, y-1, x)
		self.PrintMat(mat_temp)
		sum = append(sum, self.SumMat(mat_temp))
	}
	if x+1 < 3 {
		mat_temp := self.Swap(mat, y, x, y, x+1)
		self.PrintMat(mat_temp)
		sum = append(sum, self.SumMat(mat_temp))
	}
	if y+1 < 3 {
		mat_temp := self.Swap(mat, y, x, y+1, x)
		self.PrintMat(mat_temp)
		sum = append(sum, self.SumMat(mat_temp))
	}
	if x-1 > -1 {
		mat_temp := self.Swap(mat, y, x, y, x-1)
		self.PrintMat(mat_temp)
		sum = append(sum, self.SumMat(mat_temp))
	}

	for _, val := range sum {
		if self.Ids.ItemNotIn(val) {
			self.Ids = append(self.Ids, val)
		}
	}

	fmt.Println(self.Ids)
}

// función principal
func main() {
	slide := NewSlidePuzzleSolution()
	fmt.Println(slide.Sum)

	slide_solvr := NewSlidePuzzleSolvr()
	fmt.Println("\nMatriz original")
	slide_solvr.PrintMat(slide_solvr.Mat)
	fmt.Println("\nMatrices copias")
	slide_solvr.MoveSpace(slide_solvr.Mat)
}
