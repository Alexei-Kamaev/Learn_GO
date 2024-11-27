package main
import (
	"fmt"
)
func main () {
	var numUser int
	fmt.Println("Введите целое число")
	fmt.Scanln(&numUser)
	fmt.Printf("Следующее за числом %v число %v.\n",numUser, numUser+1)
	fmt.Printf("Для числа %v предыдущее число %v.\n", numUser, numUser-1)
}