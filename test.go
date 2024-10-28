package main

import "fmt"

func main() {
	var N int
	var count = 0
	fmt.Scanln(&N)
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			for k := 0; k < N; k++ {
				count++
			}
		}
	}

	fmt.Println("Количество операций:", count)
}
