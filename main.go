package main

import (
	"fmt"

	operations "github.com/dfgarciasdev/learning_go_module_a"
)

func main() {
	values := operations.Values{A: 1, B: 2}
	fmt.Println(values.Sum())
}
