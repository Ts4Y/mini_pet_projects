package fuel

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func CheckFuel(f string) {
	strs := strings.Split(f, "/")
	x := strs[0]
	y := strs[len(strs)-1]

	intx, err := strconv.Atoi(x)

	if err != nil {
		fmt.Println("POSHEL V JOPU", err)
	}

	inty, err := strconv.Atoi(y)

	if err != nil {
		fmt.Println("POSHEL V JOPU", err)
	}

	if inty == 0 {
		fmt.Println("NA NOL NELZYA DELIT DAUN")
		return
	} else {

	
		percent := (float32(intx) / float32(inty)) * 100

		if percent <= 1 {
			fmt.Println("E")
		} else if percent >= 99 {
			fmt.Println("F")
		} else {
			fmt.Println(math.Round(float64(percent)),"%")
		}
	}

}
