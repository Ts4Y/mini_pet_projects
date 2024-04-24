package tip

import (
	"fmt"
	"strconv"
	"strings"
)

func DollarsToFloat(d string)float32 {

PreffixOff := strings.TrimPrefix(d, "$")

dtf, err:= strconv.ParseFloat(PreffixOff, 32)

if err != nil{
	fmt.Println("Ошибка:", err)
}
return float32(dtf)
}

func PercentToFloat(p string)float32{

SuffixOff:= strings.TrimSuffix(p,"%")

ptf, err := strconv.ParseFloat(SuffixOff, 32)

if err != nil{
	fmt.Println("Ошибка:", err)
}

tip:= ptf/100

return float32(tip)

}