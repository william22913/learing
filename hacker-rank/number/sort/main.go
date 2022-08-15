package main

import (
	"fmt"
)

func main() {
	var data = []int{2, 3, 4, 5, 6, 7, 8, 9, 10, 1}

	for i := len(data); i > 0; i-- {
		min := 0
		for j := i - 1; j >= 0; j-- {
			var change = false
			if j == i-1 {
				min = data[j]
				continue
			}

			if min < data[j] {
				data[j+1] = data[j]
				change = true
			} else {
				if j+2 < len(data) {
					if data[j+1] == data[j+2] {
						data[j+1] = min
						change = true
					}
				}
			}

			if change {
				fmt.Println(data)
			}
			
			if j == 0 {
				if data[j] == data[j+1] {
					data[j] = min
					change = true
				}

				if change {
					fmt.Println((print(data)))
				}
			}

		}

	}

}

func print(data []int) (result string) {
	for i := range data {
		result = result + fmt.Sprintf("%d ", data[i])
	}
	return result
}
