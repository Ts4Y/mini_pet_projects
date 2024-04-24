package enstein

import (
	"fmt"
)

func main() {

	var m int
	fmt.Println("Введите число")
	fmt.Scan(&m)

	fmt.Println(speed(300000000, m))


}

func speed(c, m int) int {

	return m * c * c

}
