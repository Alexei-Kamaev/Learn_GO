package main
import (
	"fmt"
)
func main () {
	var a, b, c int
	fmt.Println("Введите три целых числа")
	fmt.Scanln(&a, &b, &c)
	fmt.Printf("%v  %v  %v", a, b, c)
}