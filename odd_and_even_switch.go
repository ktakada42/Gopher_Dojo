package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		switch {
		case i%2 == 1:
			fmt.Printf("%d-奇数\n", i)
		case i%2 == 0:
			fmt.Printf("%d-偶数\n", i)
		}
	}
}
