package main

import (
	"container/heap"
	"testing"
)

func Test1631(t *testing.T) {

	tests := []struct {
		nums [][]int
		res  int
	}{
		{
			nums: [][]int{{1, 2, 2}, {3, 8, 2}, {5, 3, 5}},
			res:  2,
		},
		{
			nums: [][]int{{1, 2, 3}, {3, 8, 4}, {5, 3, 5}},
			res:  1,
		},
		{
			nums: [][]int{{3}},
			res:  0,
		},
		{
			nums: [][]int{{1, 10, 6, 7, 9, 10, 4, 9}},
			res:  9,
		},
	}

	for _, tt := range tests {
		res := minimumEffortPath(tt.nums)
		if res != tt.res {
			t.Errorf("Error in 1631. Expected %v, got %v", tt.res, res)
		}
	}
}

type MinHeap [][]int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i][2] < h[j][2] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h MinHeap) Peek() []int        { return h[0] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.([]int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// tc: O(wh log(wh)), w & h: width & height of points
func minimumEffortPath(heights [][]int) int {
	w, h := len(heights[0])-1, len(heights)-1
	mh := &MinHeap{}
	heap.Init(mh)

	visited := make([][]bool, len(heights))
	for i := range visited {
		visited[i] = make([]bool, len(heights[0]))
	}
	heap.Push(mh, []int{0, 0, 0})

	for mh.Len() > 0 {
		top := heap.Pop(mh).([]int)

		if visited[top[0]][top[1]] {
			continue
		}
		visited[top[0]][top[1]] = true

		if top[0] == h && top[1] == w {
			return top[2]
		}

		for _, d := range dir {
			newY, newX := top[0]+d[0], top[1]+d[1]

			if validPoint(heights, newX, newY) && !visited[newY][newX] {
				heap.Push(mh, []int{
					newY, newX, max(top[2],
						abs(heights[newY][newX]-heights[top[0]][top[1]])),
				})
			}
		}
	}

	return 0
}

var dir = [][]int{
	{0, 1},
	{0, -1},
	{1, 0},
	{-1, 0},
}

func validPoint(heights [][]int, x, y int) bool {
	return x >= 0 && y >= 0 && x < len(heights[0]) && y < len(heights)
}

func max(i, j int) int {
	if i >= j {
		return i
	}
	return j
}

func min(i, j int) int {
	if i <= j {
		return i
	}
	return j
}

func abs(i int) int {
	if i >= 0 {
		return i
	}
	return -i
}

//type Coords struct {
//x int
//y int
//}

//type Node struct {
//visited bool
//value   float64
//cost    float64
//coords  Coords
//}

//func newNodeList(heights [][]int) [][]*Node {
//nodes := [][]*Node{}

//for i := range heights {
//ns := []*Node{}
//for j := range heights[i] {
//ns = append(ns, &Node{false, float64(heights[i][j]), math.Inf(1), Coords{i, j}})
//}
//nodes = append(nodes, ns)
//}

//return nodes
//}

//func minimumEffort(heights [][]int) int {
//nodes := newNodeList(heights)
//var visit func(int, *Node)

//visit = func(prev int, n *Node) {

//if n.visited {
//return
//}

//visitNeighbor := func(x, y int) {
//neighbor := nodes[x][y]
//if neighbor.cost > math.Abs(n.value-neighbor.value) && !neighbor.visited {
//neighbor.cost = math.Abs(n.value - neighbor.value)
//}
//}

//neighbors := []Coords{}

//if n.coords.x > 0 {
//visitNeighbor(n.coords.x-1, n.coords.y)
//neighbors = append(neighbors, Coords{n.coords.x - 1, n.coords.y})
//}
//if n.coords.y > 0 {
//visitNeighbor(n.coords.x, n.coords.y-1)
//neighbors = append(neighbors, Coords{n.coords.x, n.coords.y - 1})
//}
//if n.coords.x < len(nodes)-1 {
//visitNeighbor(n.coords.x+1, n.coords.y)
//neighbors = append(neighbors, Coords{n.coords.x + 1, n.coords.y})
//}
//if n.coords.y < len(nodes[0])-1 {
//visitNeighbor(n.coords.x, n.coords.y+1)
//neighbors = append(neighbors, Coords{n.coords.x, n.coords.y + 1})
//}

//n.visited = true

//for _, neighbor := range neighbors {
//visit(int(n.value), nodes[neighbor.x][neighbor.y])
//}

//}

//visit(0, nodes[0][0])
//var goBack func(Coords) float64

//goBack = func(coords Coords) float64 {
//node := nodes[coords.x][coords.y]
//if coords.x == 0 && coords.y == 0 {
//return 0
//}
//neighbors := []*Node{}

//if node.coords.x > 0 {
//neighbors = append(neighbors, nodes[node.coords.x-1][node.coords.y])
//}

//if node.coords.y > 0 {
//neighbors = append(neighbors, nodes[node.coords.x][node.coords.y-1])
//}

//if node.coords.x > len(nodes)-1 {
//neighbors = append(neighbors, nodes[node.coords.x+1][node.coords.y])
//}

//if node.coords.y > len(nodes[0])-1 {
//neighbors = append(neighbors, nodes[node.coords.x][node.coords.y+1])
//}

//for _, neighbor := range neighbors {
//if node.cost == math.Abs(node.value-neighbor.value) {
//return math.Max(node.cost, goBack(neighbor.coords))
//}
//}
//return 0
//}

//return int(goBack(Coords{len(nodes) - 1, len(nodes[0]) - 1}))
//}

//func minimumEffort(heights [][]int) int {
//var recursiveEffort func([2]int, []*int, int, int) int

//recursiveEffort = func(coords [2]int, history []*int, prev, diff int) int {
//val := &heights[coords[0]][coords[1]]

//if coords[0] == len(heights)-1 && coords[1] == len(heights[0])-1 {
//return int(math.Abs(float64(prev - *val)))
//}

//for _, h := range history {
//if h == val {
//return -1
//}
//}

//history = append(history, val)

//var possibleResults []int

//if coords[0] > 0 {
//possibleResults = append(possibleResults, recursiveEffort([2]int{coords[0] - 1, coords[1]}, history, *val, diff))
//}
//if coords[0] < len(heights)-1 {
//possibleResults = append(possibleResults, recursiveEffort([2]int{coords[0] + 1, coords[1]}, history, *val, diff))
//}

//if coords[1] > 0 {
//possibleResults = append(possibleResults, recursiveEffort([2]int{coords[0], coords[1] - 1}, history, *val, diff))
//}
//if coords[1] < len(heights[0])-1 {
//possibleResults = append(possibleResults, recursiveEffort([2]int{coords[0], coords[1] + 1}, history, *val, diff))
//}

//minimum_val := -1
//for _, v := range possibleResults {
//if v == -1 {
//continue
//}
//if minimum_val == -1 || v < minimum_val {
//minimum_val = v
//}
//}

//if minimum_val == -1 {
//return minimum_val
//}
//return int(math.Max(float64(minimum_val), math.Abs(float64(prev-*val))))
//}
//return recursiveEffort([2]int{0, 0}, []*int{}, heights[0][0], 0)
//}
