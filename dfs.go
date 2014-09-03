package main

import (
	"container/list"
	"fmt"
)

type Map map[int][]int
type Parent map[int]int
type Array []int

type Queue struct {
	queue list.List
}

func NewQueue() *Queue {
	self := new(Queue)

	return self
}

func (self *Queue) Push(val ...int) {
	for _, val := range val {
		self.queue.PushBack(val)
	}
}

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

func (self *Array) ItemNotIn(val int) bool {
	for _, value := range *self {
		if val == value {
			return false
		}
	}

	return true
}

func (self *Queue) Len() int {
	return self.queue.Len()
}

func dfs(graph Map, vcted Array, path *Queue, start, end int) {
	vcted = append(vcted, start)

	if start == end {
		path.Push(vcted...)
	}

	for _, v := range graph[start] {
		if vcted.ItemNotIn(v) {
			dfs(graph, vcted, path, v, end)
		}
	}
}

func main() {
	graph := Map{
		1:  {2, 3, 4},
		2:  {1, 5, 6},
		3:  {1},
		4:  {1, 7, 8},
		5:  {2, 9, 10},
		6:  {2},
		7:  {4, 11, 12},
		8:  {4},
		9:  {5},
		10: {5},
		11: {7},
		12: {7},
	}
	vcted := Array{}
	path := NewQueue()

	dfs(graph, vcted, path, 9, 12)
	path.PrintQueue()
}
