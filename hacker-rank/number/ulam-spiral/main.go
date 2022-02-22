package main

import "fmt"

var funcs []func(int, int) (int, int)

func main() {

	funcs = append(funcs,
		func(x int, y int) (int, int) {
			return x, y + 1
		},
		func(x int, y int) (int, int) {
			return x - 1, y
		},
		func(x int, y int) (int, int) {
			return x, y - 1
		},
		func(x int, y int) (int, int) {
			return x + 1, y
		},
	)

	x := 7

	if x%2 == 0 {
		panic("Not an odd number")
	}

	metric := make([][]int, x)
	for i := range metric {
		metric[i] = make([]int, x)
	}

	mid := x / 2
	position := []int{mid, mid}
	metric[position[0]][position[1]] = 1

	mv := 1
	ptr := 0
	val := 2
	counter := 1
	valid := true

	for valid {

		if val != 2 {
			ptr++
		}

		for i := 0; i < mv; i++ {

			ptr, position[0], position[1] = move(ptr, position[0], position[1])
			metric[position[0]][position[1]] = val

			if val == x*x {
				valid = false
				break
			}
			val++
		}

		if counter%2 == 0 {
			mv = mv + 1
		}

		counter++

	}

	print(metric)
}

func print(metric [][]int) {
	for i := 0; i < len(metric); i++ {
		for j := 0; j < len(metric[i]); j++ {
			fmt.Print(metric[i][j], "\t")
		}
		fmt.Println()
	}
}

func move(idx int, x int, y int) (int, int, int) {

	if idx == 4 {
		idx = 0
	}

	x, y = funcs[idx](x, y)

	return idx, x, y
}
