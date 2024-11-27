package main
import (
	"fmt"
)
func main () {
	fmt.Println("Введите число.")
	var num int
	fmt.Scanln(&num)
	fmt.Println("Вы ввели число:", num)
}