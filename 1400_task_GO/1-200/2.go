//Вывести на одной строке числа 48, 53 и 151 с двумя пробелами.

package main

import (
	"fmt"
	"strconv"
)
func main () {
	var a, b, c = 48, 53, 151
	d:= "  "
	fmt.Println(strconv.Itoa(a) + d + strconv.Itoa(b) + d + strconv.Itoa(c))
}