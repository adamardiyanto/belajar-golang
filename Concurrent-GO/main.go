package main

import (
	"fmt"
	"runtime"
)

func print(count int, msg string) {
	for i := 0; i < count; i++ {
		fmt.Println((i + 1), msg)
	}
}

func main() {
	runtime.GOMAXPROCS(6)

	print(5, "rutin 2")
	go print(5, "rutin 1")
	var input string
	fmt.Scanln(&input)
}
