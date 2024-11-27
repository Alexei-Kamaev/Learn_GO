package main
import (
	"fmt"
)
func main () {
	fmt.Println("Введите название футбольной команды")
	var nameClub string
	fmt.Scanln(&nameClub)
	fmt.Printf("%v - это чемпион!", nameClub)
}