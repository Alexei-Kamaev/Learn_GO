package main

import (
	"fmt"
)
func main () {
	pi := 3.1415926
	fmt.Printf("Дано такое число %f, надо уменьшить разрядность. \n", pi)
	piInt := int64(pi*1000)
	piFloat := float64(piInt)/1000
	fmt.Println("Уменьшенная разрядность, без округлений:", piFloat)
}