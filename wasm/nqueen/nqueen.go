package main

import (
	"fmt"
	"math/rand"
	"strings"
	"syscall/js"
)

func main() {
	c := make(chan struct{}, 0)
	js.Global().Set("NQueen", js.FuncOf(wrapper))
	<-c
}

func wrapper(this js.Value, i []js.Value) interface{} {
	n := i[0].Int()
	arr, _ := solveNQueens(n)
	ans := make([]interface{}, len(arr))
	for i := range arr {
		ans[i] = arr[i]
	}
	fmt.Println(arr)
	return js.ValueOf(ans)
}

func solveNQueens(n int) ([]int, int) {
	i := 0
	for {
		i++
		yes, board := minConflicts(n)
		if yes {
			return board, i
		}
	}
}

func minConflicts(n int) (bool, []int) {

	//intialize board
	board := make([]int, n)
	for i := 0; i < n; i++ {
		board[i] = i
	}
	//shuffle board
	for i := 0; i < n; i++ {
		r := rand.Intn(n)
		board[i], board[r] = board[r], board[i]
	}
	//reasonably good?
	for i := 0; i < n; i++ {
		cols, ru, lu := genIndexes(board, len(board))
		col, _ := findBestCol(board, len(board), i, cols, ru, lu)
		board[i] = col
	}
	solve(board, len(board), len(board))
	cols, ru, lu := genIndexes(board, len(board))
	if _, worst := findWorstIndex(board, len(board), cols, ru, lu); worst == 0 {
		return true, board
	}
	return false, nil
}

func genIndexes(board []int, width int) (*[]int, *[]int, *[]int) {
	cols := make([]int, len(board))
	ru, lu := make([]int, width+len(board)), make([]int, width+len(board))
	for i, v := range board {
		cols[v]++
		ru[i+v]++
		lu[width-v-1+i]++
	}
	return &cols, &ru, &lu
}

func findWorstIndex(board []int, width int, cols, ru, lu *[]int) (int, int) {
	worstIndex := -1
	worst := -1
	for i := 0; i < len(board); i++ {
		c := board[i]
		current := (*cols)[c] + (*ru)[i+c] + (*lu)[width-c-1+i] - 3
		if current > worst {
			worst = current
			worstIndex = i
		}
	}
	return worstIndex, worst
}

func findBestCol(board []int, width, i int, cols, ru, lu *[]int) (int, int) {
	best := 9999999999999
	bestCol := -1
	c := board[i]
	(*cols)[c]--
	(*ru)[i+c]--
	(*lu)[width-c-1+i]--
	for n := 0; n < width; n++ {
		current := (*cols)[n] + (*ru)[i+n] + (*lu)[width-n-1+i]
		if current < best {
			best = current
			bestCol = n
		}
	}
	(*cols)[c]++
	(*ru)[i+c]++
	(*lu)[width-c-1+i]++
	return bestCol, best
}

func solve(board []int, width, trials int) (bool, []int) {
	cols, ru, lu := genIndexes(board, width)
	for n := 0; n < trials; n++ {
		worstIndex, worst := findWorstIndex(board, width, cols, ru, lu)
		if worst == 0 {
			// draw(board, len(board))
			return true, board
		}
		bestCol, _ := findBestCol(board, width, worstIndex, cols, ru, lu)
		(*cols)[board[worstIndex]]--
		(*ru)[worstIndex+board[worstIndex]]--
		(*lu)[width-board[worstIndex]-1+worstIndex]--
		board[worstIndex] = bestCol
		(*cols)[bestCol]++
		(*ru)[worstIndex+bestCol]++
		(*lu)[width-bestCol-1+worstIndex]++
	}
	return false, nil
}

func draw(board []int, width int) {
	for _, row := range board {
		str := strings.Repeat(".", row) + "Q" + strings.Repeat(".", width-row-1)
		fmt.Println(str)
	}
}

//old code below
/*
//NQueen returns a nqueen solution for an nxn board
func NQueen(this js.Value, i []js.Value) interface{} {
	n := i[0].Int()
	if n < 4 {
		return js.ValueOf("")
	}
	ans := solveNQueens(n)[0]
	response := ""
	for _, str := range ans {
		response += str + "\n"
	}
	return js.ValueOf(response)
}

func solveNQueens(n int) [][]string {
	ans := [][]string{}

	cols := make([]bool, n)
	dag1 := make([]bool, 2*n-1)
	dag2 := make([]bool, 2*n-1)
	board := make([]int, n)
	// init the board
	// for i := 0; i < n; i++ {
	//     board[i] = -1
	// }

	magic(n, 0, cols, dag1, dag2, board, &ans)
	return ans
}

func magic(n int, row int, cols, dag1, dag2 []bool, board []int, ans *[][]string) bool {
	if row == n {
		c := render(board)
		*ans = append(*ans, c)
		return true
	}
	for c := 0; c < n; c++ {
		if !(cols[c] || dag1[c+row] || dag2[row-c+n-1]) {
			cols[c] = true
			dag1[c+row] = true
			dag2[row-c+n-1] = true
			board[row] = c

			yes := magic(n, row+1, cols, dag1, dag2, board, ans)
			if yes {
				return true
			}

			cols[c] = false
			dag1[c+row] = false
			dag2[row-c+n-1] = false
		}
	}
	return false
}

func render(board []int) []string {
	res := make([]string, len(board))
	for i, v := range board {
		for j := 0; j < v; j++ {
			res[i] += "."
		}
		res[i] += "Q"
		for j := v + 1; j < len(board); j++ {
			res[i] += "."
		}
	}
	return res
}
*/
