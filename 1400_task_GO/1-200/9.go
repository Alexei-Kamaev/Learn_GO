package main
import (
	"fmt"
)
func main () {
	fmt.Println("Введите число")
	var num int
	fmt.Scanln(&num)
	fmt.Printf("%v - вот какое число Вы ввели", num)
}