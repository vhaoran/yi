package z

import (
	"fmt"
	"testing"

	"github.com/davecgh/go-spew/spew"

	"github.com/vhaoran/yi/cmn"
)

func Test_get_na_yin(t *testing.T) {
	for y := 1970; y <= 2030; y++ {
		//d := calendar.NewSolarFromYmd(y,5,1)
		gz := cmn.GetYearGanZi(y)
		naYin := new(NaYinGet).Get(gz)
		wuxing := new(NaYinGet).GetWuXin(naYin)
		fmt.Println(y, "---", gz, ",", naYin, wuxing)
	}
}
func TestBaZiGet_GetBySolar(t *testing.T) {
	y, m, d, h, mm := 1970, 12, 14, 23, 30
	z := BaZiGetX.GetBySolar(y, m, d, h, mm)
	fmt.Println("-----------------")
	spew.Dump(z)
}
