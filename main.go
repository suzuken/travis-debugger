package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("=== start tarvis-debugger ===")

	fmt.Println("=== print environment variables ===")

	for _, e := range os.Environ() {
		fmt.Printf("%s\n", e)
	}

	fmt.Println("=== print travis specific environment variables ===")
	envs := os.Environ()
	for _, e := range envs {
		if strings.HasPrefix(e, "TRAVIS") {
			fmt.Printf("%s\n", e)
		}
	}

	fmt.Println("=== finished tarvis-debugger ===")
}
