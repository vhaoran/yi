package cmn

import (
	"fmt"
	"testing"

	"github.com/6tail/lunar-go/calendar"
	"github.com/vhaoran/vchat/common/ytime"
)

//solar
func Test_a(t *testing.T) {
	y, m, d := 2020, 2, 3
	dt := ytime.OfInt(y, m, d).Time
	lunar := calendar.NewLunarFromDate(dt)
	fmt.Println(lunar.ToFullString())
	fmt.Println("-----------------")
	fmt.Println(lunar.GetSolar().ToFullString())
	fmt.Println("-----------------")
	fmt.Println(lunar.GetJie())
	fmt.Println("-----------------")
	fmt.Println(lunar.GetJieQiTable())
	fmt.Println("-----------------")
}

// 家国取四柱
func Test_b(t *testing.T) {
	//dt := ytime.OfInt(1970, 11, 16, 4).Time
	d := calendar.NewLunar(1970, 11, 16, 4, 0, 0)

	fmt.Println(d.ToFullString())
	fmt.Println(d.ToFullString())
	fmt.Println("-----------------")
	fmt.Println("-----------------")

	fmt.Println(d.GetBaZi())
}
func Test_GetYearGanZi(t *testing.T) {

}
