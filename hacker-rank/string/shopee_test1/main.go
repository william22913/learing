package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func minimumDolls(dollSizes []int) int {
	sort.Ints(dollSizes)

	var used []int
	var doll map[int][]int
	var counter int
	doll = make(map[int][]int)

	for i := 0; i < len(dollSizes)-1; i++ {
		usedData := dollSizes[i]
		for j := i + 1; j < len(dollSizes); j++ {
			found := false
			for k := 0; k < len(used); k++ {
				if used[k] == i || used[k] == j {
					found = true
					break
				}
			}

			if !found {
				if usedData+1 == dollSizes[j] {
					counter++
					usedData++
					used = append(used, j)
					doll[i] = append(doll[i], j)
				}
			}
		}
	}

	left := (len(dollSizes) - (counter + len(doll)))

	return len(doll) + left
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString(']')
	text = strings.Replace(text, "[", "", -1)
	text = strings.Replace(text, "]", "", -1)
	data := strings.Split(text, ",")

	var dataUsed []int

	for i := 0; i < len(data); i++ {
		a, err := strconv.Atoi(data[i])
		if err != nil {
			fmt.Println(err)
			return
		}
		dataUsed = append(dataUsed, a)
	}

	fmt.Println(minimumDolls(dataUsed))
}
