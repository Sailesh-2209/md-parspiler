package main

import (
	"fmt"
	"os"
)

func main() {
	for _, file := range os.Args[1:] {
		byteArr, err := os.ReadFile(file)
		if err != nil {
			fmt.Fprintf(os.Stderr, "md-parspiler: Unable to read file %s.\n", file)
			continue
		}
		data := string(byteArr)
		fmt.Println(data)
	}
}
