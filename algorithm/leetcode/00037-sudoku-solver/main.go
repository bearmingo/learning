package main

import (
	"bytes"
	"fmt"
	"strings"
)

// Dance Links X
// https://www.cnblogs.com/wujiechao/p/5767124.html
// https://leetcode.com/problems/sudoku-solver/
// [算法实践——舞蹈链（Dancing Links）算法求解数独](https://www.cnblogs.com/grenet/p/3163550.html)

const MaxUint = ^uint(0)
const MinUint = 0
const MaxInt = int(MaxUint >> 1)
const MinInt = -MaxInt - 1

type node struct {
	left  *node
	right *node
	up    *node
	down  *node
	ch    *node // column's header
	count int
	num   int // represent the solution
}

type danceLinkX struct {
	header   *node
	cols     []*node
	n        int // number of row
	m        int // number os col
	flag     bool
	solution []int

	results [][]int
}

const (
	maxColNum = 324 // 9 * 9 * 3 + 81
)

func create(n, m int) *danceLinkX {
	header := node{}
	header.left = &header
	header.right = &header

	cols := make([]*node, maxColNum)
	for i := 0; i < len(cols); i++ {
		n := &node{}
		n.ch = n // column's header is self
		n.up = n
		n.down = n
		n.right = &header
		n.left = header.left
		header.left.right = n
		header.left = n
		cols[i] = n
	}

	d := danceLinkX{
		n:        n,
		m:        m,
		header:   &header,
		cols:     cols,
		solution: make([]int, 81),
		results:  make([][]int, 0),
	}

	return &d
}

func (d *danceLinkX) remove(n *node) {
	ch := n.ch // column head
	ch.right.left = ch.left
	ch.left.right = ch.right
	// delete this column that contains this node
	for c := ch.down; c != ch; c = c.down {
		// delete row
		for row := c.right; row != c; row = row.right {
			row.down.up = row.up
			row.up.down = row.down
			row.ch.count--
		}
	}
}

// resume 删除一个节点时，删除对应的头即可
func (d *danceLinkX) resume(n *node) {
	ch := n.ch

	for c := ch.up; c != ch; c = c.up {
		for r := c.left; r != c; r = r.left {
			r.down.up = r
			r.up.down = r
			r.ch.count++
		}
	}

	ch.right.left = ch
	ch.left.right = ch
}

func (d *danceLinkX) link(num int, cs []int) {
	var rowHeader *node
	for i := 0; i < len(cs); i++ {
		ch := d.cols[cs[i]].ch
		temp := &node{}
		temp.num = num
		if i == 0 {
			rowHeader = temp
			rowHeader.right = rowHeader
			rowHeader.left = rowHeader
		}

		// insert to end of row list
		temp.right = rowHeader
		temp.left = rowHeader.left
		rowHeader.left.right = temp
		rowHeader.left = temp

		// insert to end of columns list
		temp.down = ch
		temp.up = ch.up // 这行最后一个

		temp.ch = ch
		ch.up.down = temp
		ch.up = temp
		ch.count++
	}
}

func (d *danceLinkX) chooseColumn() *node {
	min := MaxInt
	var result *node
	// 选择列中节点最少的列
	for c := d.header.right; c != d.header; c = c.right {
		if c.count < min {
			min = c.count
			result = c
		}
	}

	return result
}

func (d *danceLinkX) saveSolution() {
	copyed := make([]int, len(d.solution))
	copy(copyed, d.solution)

	d.results = append(d.results, copyed)
}

func (d *danceLinkX) search(k int) bool {
	if d.flag == true {
		return true
	}
	if d.header.right == d.header {
		d.flag = true
		d.saveSolution()
		return true
	}

	c := d.chooseColumn()
	d.remove(c)

	for s := c.down; s != c; s = s.down {
		d.solution[k] = s.num // add the solution

		// remove the column in the right
		for r := s.right; r != s; r = r.right {
			d.remove(r)
		}

		// continue searching
		d.search(k + 1)

		// after searched, resume the state
		for l := s.left; l != s; l = l.left {
			d.resume(l)
		}
	}
	d.resume(c)

	return false
}

func (d *danceLinkX) getColumnIndex(r, c, val int) []int {
	indexs := make([]int, 4)
	// 第(r*9+value)列表示r行填写了个数字val
	indexs[0] = r*9 + val - 1
	// 第(81+c*9 + val)列表示c列填写了数字val
	indexs[1] = 81 + c*9 + val - 1

	tr := r / 3
	tc := c / 3
	b := tr*3 + tc // 宫的序号
	// 第n个宫，填写了数字val
	indexs[2] = 162 + b*9 + val - 1

	// 第r行和n列填写了数字
	indexs[3] = 243 + r*9 + c
	return indexs
}

func (d *danceLinkX) construct(puzzle [][]byte) {
	for pr := 0; pr < len(puzzle); pr++ {
		for pc := 0; pc < len(puzzle[pr]); pc++ {
			i := pr*9 + pc
			word := puzzle[pr][pc]
			r, c := i/9, i%9

			if word == '.' {
				for val := 1; val <= 9; val++ {
					d.link(r*100+c*10+val, d.getColumnIndex(r, c, val))
				}
			} else {
				num := int(word) - int('0')
				d.link(r*100+c*10+num, d.getColumnIndex(r, c, num))
			}
		}
	}
}

func (d danceLinkX) printSolution(board [][]byte) {
	// result := make([][]byte, 9)
	// for i := 0; i < 9; i++ {
	// 	result[i] = make([]byte, 9)
	// }
	sol := d.results[0]
	for _, n := range sol {
		r, c := n/100, n/10%10
		v := n % 10
		board[r][c] = byte(int('0') + v)
	}
}

func printResult(res [][]string) {
	buf := bytes.Buffer{}
	buf.WriteString("[")
	for _, row := range res {
		buf.WriteString("[")
		buf.WriteString(strings.Join(row, ","))
		buf.WriteString("]")
	}
	buf.WriteString("]")
	fmt.Print(buf.String())
}

func solveSudoku(board [][]byte) {
	d := create(9, 9)
	d.construct(board)
	d.search(0)
	d.printSolution(board)
}

// --------------------------
func printSolvedPuzzle(board [][]byte) {
	buf := bytes.Buffer{}
	for _, row := range board {
		for _, b := range row {
			buf.WriteByte(b)
			buf.WriteByte(',')
		}
		buf.WriteByte('\n')
	}

	fmt.Print(buf.String())
}

func main() {

	/*
		puzzle := make([][]byte, 9, 9)
		puzzle[0] = []byte{'5', '3', '.', '.', '7', '.', '.', '.', '.'}
		puzzle[1] = []byte{'6', '.', '.', '1', '9', '5', '.', '.', '.'}
		puzzle[2] = []byte{'.', '9', '8', '.', '.', '.', '.', '6', '.'}
		puzzle[3] = []byte{'8', '.', '.', '.', '6', '.', '.', '.', '3'}
		puzzle[4] = []byte{'4', '.', '.', '8', '.', '3', '.', '.', '1'}
		puzzle[5] = []byte{'7', '.', '.', '.', '2', '.', '.', '.', '6'}
		puzzle[6] = []byte{'.', '6', '.', '.', '.', '.', '2', '8', '.'}
		puzzle[7] = []byte{'.', '.', '.', '4', '1', '9', '.', '.', '5'}
		puzzle[8] = []byte{'.', '.', '.', '.', '8', '.', '.', '7', '9'}
	*/
	/*puzzle := [][]byte{
		[]byte{'.', '.', '9', '7', '4', '8', '.', '.', '.'},
		[]byte{'7', '.', '.', '.', '.', '.', '.', '.', '.'},
		[]byte{'.', '2', '.', '1', '.', '9', '.', '.', '.'},
		[]byte{'.', '.', '7', '.', '.', '.', '2', '4', '.'},
		[]byte{'.', '6', '4', '.', '1', '.', '5', '9', '.'},
		[]byte{'.', '9', '8', '.', '.', '.', '3', '.', '.'},
		[]byte{'.', '.', '.', '8', '.', '3', '.', '2', '.'},
		[]byte{'.', '.', '.', '.', '.', '.', '.', '.', '6'},
		[]byte{'.', '.', '.', '2', '7', '5', '9', '.', '.'},
	}*/

	puzzle := [][]byte{
		[]byte{'.', '.', '9', '7', '4', '8', '.', '.', '.'},
		[]byte{'7', '.', '.', '.', '.', '.', '.', '.', '.'},
		[]byte{'.', '2', '.', '1', '.', '9', '.', '.', '.'},
		[]byte{'.', '.', '7', '.', '.', '.', '2', '4', '.'},
		[]byte{'.', '6', '4', '.', '1', '.', '5', '9', '.'},
		[]byte{'.', '9', '8', '.', '.', '.', '3', '.', '.'},
		[]byte{'.', '.', '.', '8', '.', '3', '.', '2', '.'},
		[]byte{'.', '.', '.', '.', '.', '.', '.', '.', '6'},
		[]byte{'.', '.', '.', '2', '7', '5', '9', '.', '.'},
	}

	solveSudoku(puzzle)
	printSolvedPuzzle(puzzle)
}
