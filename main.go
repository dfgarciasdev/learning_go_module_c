package main

import (
	"fmt"

	operations "github.com/dfgarciasdev/learning_go_module_a"
	operationsV2 "github.com/dfgarciasdev/learning_go_module_a/v2"
)

func main() {
	values := operations.Values{A: 1, B: 2}
	fmt.Println(values.Sum())
	fmt.Println(operationsV2.Sum(values.A, values.B))
}
