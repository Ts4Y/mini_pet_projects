package coke

import (
	"fmt"
	"strconv"
)

func CountMoney() {

	total := 0

	cash := map[string]bool{

		"1":  true,
		"5":  true,
		"10": true,
		"25": true,
	}

	for total < 50 {
		fmt.Println("VVEDI MONETU")
		var EnterCoin string
		fmt.Scan(&EnterCoin)

		if _, exist := cash[EnterCoin]; exist {
			IntCoin, err := strconv.Atoi(EnterCoin)
			if err != nil {
				fmt.Println("SDELAY MNE PALIROFKU", err)
			}
			
			total += IntCoin
			if total == 50{
				break
			}
			fmt.Printf("ESHE NUJNO VVESTI %d ", 50-total)

			
		} else {
			fmt.Println("Недопустимая монета. Принимаются только 1, 5, 10 и 25 центов.")
		}

	}
	
	fmt.Println("VOT VASHA KOLA")

}
