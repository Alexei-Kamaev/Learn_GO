package main
import (
	"fmt"
)
func main () {
	fmt.Println("Введите Ваше имя:")
	var nameUser string
	fmt.Scanln(&nameUser)
	fmt.Println("Ваше имя", nameUser)
}