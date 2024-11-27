package main
import (
	"fmt"
)
func main () {
	fmt.Println("Как Вас зовут?")
	var nameUser string
	fmt.Scanln(&nameUser)
	fmt.Printf("Привет, %v!", nameUser)
}