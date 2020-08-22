package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/montanaflynn/stats"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	// x, i := solveNQueens(100)
	// draw(x, len(x))

	// start := time.Now()
	// _, i := solveNQueens(1000)
	// fmt.Println(time.Since(start))
	// fmt.Println(i)
	benchMark(1024)
}

func benchMark(n int) {
	times := []int64{}
	attempts := []int64{}
	trials := 100
	for i := 0; i < trials; i++ {
		fmt.Print(i)
		start := time.Now()
		_, i := solveNQueens(n)
		times = append(times, int64(time.Since(start)))
		attempts = append(attempts, int64(i))
		fmt.Fprint(os.Stdout, "\r \r")
	}
	interval95(times, "times", float64(trials))
	interval95(attempts, "attempts", float64(trials))
}

func interval95(data []int64, label string, n float64) {
	d := stats.LoadRawData(data)
	mean, _ := d.Mean()
	stdev, _ := d.StandardDeviationSample()
	zstar := stats.NormInterval(0.95, 0, 1)[1]
	stderr := stdev / math.Pow(n, 0.5)
	// fmt.Println(mean, stderr, zstar)
	fmt.Printf("%s: %f0 ± %f1\n", label, mean, stderr*zstar)

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

	//reasonably good
	board := make([]int, n)
	cols := make([]int, len(board))
	ru, lu := make([]int, 2*len(board)), make([]int, 2*len(board))
	for i := 0; i < n; i++ {
		best := 9999999999999
		bestCols := []int{}
		for n := 0; n < len(board); n++ {
			current := cols[n] + ru[i+n] + lu[len(board)-n-1+i]
			if current == best {
				bestCols = append(bestCols, n)
			} else if current < best {
				best = current
				bestCols = []int{n}
			}
		}
		choice := bestCols[rand.Intn(len(bestCols))]
		cols[choice]++
		ru[i+choice]++
		lu[len(board)-choice-1+i]++
		board[i] = choice
	}
	solve(board, len(board), len(board))
	cols2, ru2, lu2 := genIndexes(board, len(board))
	if _, worst := findWorstIndex(board, len(board), cols2, ru2, lu2); worst == 0 {
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
