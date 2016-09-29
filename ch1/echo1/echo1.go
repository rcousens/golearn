package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	start := time.Now()

	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}

	fmt.Printf("%.2fs elapsed\n", time.Since(start).Nanoseconds())
	fmt.Println(s)
}
