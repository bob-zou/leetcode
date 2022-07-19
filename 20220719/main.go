package main

import (
	"fmt"
	"time"
)

func generate(numRows int) [][]int {
	var triangle [][]int
	for i := 0; i < numRows; i++ {
		if i == 0 {
			triangle = append(triangle, []int{1})
			continue
		}
		if i == 1 {
			triangle = append(triangle, []int{1, 1})
			continue
		}

		triangle = append(triangle, []int{})

		for j := 0; j <= i; j++ {
			if j == 0 {
				triangle[i] = append(triangle[i], 1)
				continue
			}
			if j == i {
				triangle[i] = append(triangle[i], 1)
				continue
			}

			triangle[i] = append(triangle[i], triangle[i-1][j-1]+triangle[i-1][j])
		}
	}
	return triangle
}

func main() {
	defer func(start time.Time) {
		fmt.Printf("cost: %dms\n", time.Since(start).Milliseconds())
	}(time.Now())

	fmt.Println(generate(30))
}
