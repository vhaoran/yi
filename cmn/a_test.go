package cmn

import (
	"fmt"
	"testing"

	"github.com/6tail/lunar-go/calendar"
	"github.com/vhaoran/vchat/common/ytime"
)

//solar
func Test_a(t *testing.T) {
	y, m, d := 1996,7, 21
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
	for k, v := range lunar.GetJieQiTable() {
		//fmt.Println("-----------------")
		s := fmt.Sprintf("%d-%d-%d -%d-%d-%d", v.GetYear(), v.GetMonth(),
			v.GetDay(), v.GetHour(), v.GetMinute(), v.GetSecond())
		fmt.Println(k, " : ", s)
		// spew.Dump(v)
	}

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

func Test_gong12(t *testing.T) {
	for _, gan := range Gan_list {
		for _, zhi := range Zhi_list {
			i, str := GetGong12OfGan(gan, zhi)
			fmt.Println(gan, "-", zhi, "(", i, " ", str, ")")
		}
	}
}
