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

func Test_GetDaysByMonth(t *testing.T) {
	for y := 2016; y < 2020; y++ {
		fmt.Println("-----------", y, "----------")
		for m := 1; m <= 12; m++ {
			fmt.Print(m, ":", GetDaysByMonth(y, m), "/")
		}
		fmt.Println("")
	}
}
