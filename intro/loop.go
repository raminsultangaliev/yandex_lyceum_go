package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	var l, r, cnt = 0, 1, 0
	for cnt < 10 {
		if l <= n && n <= r || cnt > 0 {
			cnt += 1
			fmt.Println(r)
		}
		l, r = r, l+r
	}
}
