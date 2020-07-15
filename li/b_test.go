package li

import (
	"fmt"
	"testing"
)

func Test_get_leap_month(t *testing.T) {
	for i := 1974; i < 1990; i++ {
		j := GetLeapMonth(i)
		days := GetLeapDays(i)
		fmt.Println("------", i, "-----------")
		if j > 0 {
			fmt.Println(" run: ", j, " :", days)
		}
	}
}
