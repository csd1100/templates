package main

import (
	"___packageName___/___projectName___/utils"
	"fmt"
)

func main() {
	fmt.Println("Hello World!")
	num1 := 10
	num2 := 13
	fmt.Printf("%d + %d = %d\n", num1, num2, utils.Add(num1, num2))
}
