package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func maxProfit(M int, prices []int) int {
	sort.Ints(prices)

	var sum int

	for i := 0; i < M; i++ {
		if prices[i] < 0 {
			sum = sum + prices[i]*-1
		}
	}

	return sum

}

func main() {

	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString(']')

	m := text[0:1]
	fmt.Println(m)
	mUsed, err := strconv.Atoi(m)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(mUsed)
	used := text[2:]
	fmt.Println(used)

	used = strings.Replace(used, "[", "", -1)
	used = strings.Replace(used, "]", "", -1)
	data := strings.Split(used, ",")

	var dataUsed []int

	for i := 0; i < len(data); i++ {
		a, err := strconv.Atoi(data[i])
		if err != nil {
			fmt.Println(err)
			return
		}
		dataUsed = append(dataUsed, a)
	}

	fmt.Println(maxProfit(mUsed, dataUsed))
}
