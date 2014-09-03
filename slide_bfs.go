package main // invocando el paquete principal

import (
	// importando bibliotecas que se usan
	"container/list"
	"fmt"
	"math"
)

type Matrix [][]int          // slice de slices
type Parent map[int]int      // map
type Graph map[int][]Pattern // map de slices de type Pattern
type Array []int             // slice
type IdMatrix map[int]Matrix // map of Matrix

/**
  Contiene la matriz y su identificador.
*/
type Pattern struct {
	Mat Matrix
	Id  int
}

/**
  Type que describe a una cola.
*/
type Queue struct {
	queue list.List
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
  Describe las características del puzzle slide desordenado. Embebe del type
  SlidePuzzleSolution.
*/
type SlidePuzzleSolvr struct {
	SlidePuzzle
	Ids   Array
	Gra   Graph
	IdMat IdMatrix
	Path  Queue
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
  Obtiene el tamaño actual de la cola.

  @param nada.

  @return nada.
*/
func (self *Queue) Len() int {
	return self.queue.Len()
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
func (self *SlidePuzzle) SumMat(mat Matrix) (sum int) {
	pwr, sum := 8, 0

	for i := 0; i < self.Height; i++ {
		for j := 0; j < self.Width; j++ {
			if mat[i][j] != 0 {
				sum += mat[i][j] * int(math.Pow(float64(10), float64(pwr)))
			}
			pwr--
		}
	}

	return
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
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 0},
	}
	self.Sum = self.SumMat(self.Mat)

	return self
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
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 0},
	}
	self.Gra = make(Graph)
	self.Ids = append(self.Ids, self.SumMat(self.Mat))
	self.IdMat = IdMatrix{self.Ids[0]: self.Mat}

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
  Crea type Pattern apartir de una matriz, que contiene un hijo matriz y su
  ID.

  @param matriz.
  @param posición del espacio en y (0).
  @param posición del espacio en x (0).
  @param posición donde se mueve el espacio en y (0)
  @param posición donde se mueve el espacio en x (0)

  @return el type Pattern con la matriz y su ID.
*/
func (self *SlidePuzzleSolvr) GetChild(mat Matrix, y, x, y_des, x_des int) (sum Pattern) {
	mat_temp := self.Swap(mat, y, x, y_des, x_des)
	sum_temp := self.SumMat(mat_temp)
	sum = Pattern{
		Mat: mat_temp,
		Id:  sum_temp,
	}

	return sum
}

/**
  Mueve el espacio (0) en la matriz (si se puede) y crea hijos de la matriz de
  padre verificando que no haya elementos repetidos calculando el indentificador
  del hijo.

  @param matriz de enteros que se desea analizar.

  @return nada.
*/
func (self *SlidePuzzleSolvr) MoveSpaceGetChild(mat Matrix) {
	y, x := self.GetPositionSpace(mat)
	sum := []Pattern{}
	id := self.SumMat(mat)

	if y-1 > -1 {
		sum = append(sum, self.GetChild(mat, y, x, y-1, x))
	}

	if x+1 < 3 {
		sum = append(sum, self.GetChild(mat, y, x, y, x+1))
	}

	if y+1 < 3 {
		sum = append(sum, self.GetChild(mat, y, x, y+1, x))
	}

	if x-1 > -1 {
		sum = append(sum, self.GetChild(mat, y, x, y, x-1))
	}

	for _, val := range sum {
		if self.Ids.ItemNotIn(val.Id) {
			self.Ids = append(self.Ids, val.Id)
			self.Gra[id] = append(self.Gra[id], val)
			self.IdMat[val.Id] = val.Mat
		}
	}
}

/**
  Obtiene e imprime el camino del resultado al padre, que será la respuesta
  para resolver el rompecabezas.

  @param grafo que contiene los hijos de los padres.
  @param ID del rompecabezas revuelto.
  @param ID del rompecabezas resuelto.

  @return nada.
*/
func (self *SlidePuzzleSolvr) Backtrace(parent Parent, start, end int) (path Queue) {
	path.queue.PushFront(end)

	for e := path.queue.Front(); e.Value.(int) != start; e = path.queue.Front() {
		path.queue.PushFront(parent[e.Value.(int)])
	}

	return
}

/**
  Realiza el BFS para encontrar el resultado del puzzle inicial.

  @param "grafo" de type Graph que contendrá todos los hijos del nodo padre
  y su identificador
  @param ID del puzzle de comienzo.
  @param ID del puzzle armado.

  @return los pasos que se tienen que hacer para armar el sliding puzzle.
*/
func (self *SlidePuzzleSolvr) Bfs(start, end int) Queue {
	queue := NewQueue()
	vcted := Array{}
	parent := Parent{}

	queue.Push(start)
	vcted = append(vcted, start)

	for queue.Len() > 0 {
		v := queue.Pop()

		if v == end {
			return self.Backtrace(parent, start, end)
		}

		self.MoveSpaceGetChild(self.IdMat[v])

		for _, u := range self.Gra[v] {
			if vcted.ItemNotIn(u.Id) {
				parent[u.Id] = v
				vcted = append(vcted, u.Id)
				queue.Push(u.Id)
			}
		}
	}

	return Queue{}
}

/**
  Imprime los pasos que hay que hacer para resolver el rompecabezas.

  @param nada.

  @return nada.
*/
func (self *SlidePuzzleSolvr) PrintPath() {
	for e := self.Path.queue.Front(); e != nil; e = e.Next() {
		self.PrintMat(self.IdMat[e.Value.(int)])
	}

	fmt.Println("\n")
}

// función principal
func main() {
	slide_solvd := NewSlidePuzzleSolution()
	slide_solvr := NewSlidePuzzleSolvr()

	slide_solvr.Path = slide_solvr.Bfs(slide_solvr.Ids[0], slide_solvd.Sum)
	slide_solvr.PrintPath()
}
