package main
import (
	"fmt"
)
func main () {
	eLog := 2.71828
	fmt.Println("Начальное число:", eLog)
	eLog_1 := int64(eLog*10)
	eLog_2 := float64(eLog_1)/10
	fmt.Println("Конечное число:", eLog_2)
}