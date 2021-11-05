package main

import "fmt"

func main() {
	index := 5
	arr := make([][]int, index)
	for i := 0; i < len(arr); i++ {
		arr[i] = make([]int, index)
	}

	val := 1
	row := 0
	col := -1
	move := 1
	max := index

	for i := 0; i < index; i++ {
		for j := 0; j < max; j++ {
			col = col + move
			arr[row][col] = val
		}
		val++
		max--
		for j := 0; j < max; j++ {
			row = row + move
			arr[row][col] = val
		}
		val++
		move = move * -1
	}

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[0]); j++ {
			fmt.Print(arr[i][j], "\t")
		}
		fmt.Println()
	}
}
