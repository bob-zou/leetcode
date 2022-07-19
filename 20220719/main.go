package main

import (
	"fmt"
	"time"
)

func generate(numRows int) [][]int {
	ans := make([][]int, numRows)
	for i := range ans {
		ans[i] = make([]int, i+1)
		ans[i][0], ans[i][i] = 1, 1 // 首尾固定1
		for j := 1; j < i; j++ {
			ans[i][j] = ans[i-1][j-1] + ans[i-1][j]
		}
	}
	return ans
}

func main() {
	defer func(start time.Time) {
		fmt.Printf("cost: %dms\n", time.Since(start).Milliseconds())
	}(time.Now())

	fmt.Println(generate(30))
}
