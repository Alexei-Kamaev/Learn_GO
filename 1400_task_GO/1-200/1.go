//Вывести в строке числа 32, 19 и 80 с пробелом между ними.
package main

import (
	"fmt"
	"strconv"
)
func main () {
	var a, b, c = 32, 19, 80
	var d = " "
	fmt.Println(strconv.Itoa(a) + d + strconv.Itoa(b) + d + strconv.Itoa(c))
}