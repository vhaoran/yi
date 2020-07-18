package cmn

import (
	"fmt"
	"testing"

	"github.com/6tail/lunar-go/calendar"
)

func Test_a(t *testing.T) {
	lunar := calendar.NewLunarFromYmd(2004, 9, 25)
	fmt.Println(lunar.ToFullString())
	fmt.Println(lunar.GetSolar().ToFullString())
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