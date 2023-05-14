package main

import (
	"fmt"
	"template/hello-world/utils"
)

func main() {
	fmt.Println("Hello World!")
	num1 := 10
	num2 := 13
	fmt.Printf("%d + %d = %d\n", num1, num2, utils.Add(num1, num2))
}
