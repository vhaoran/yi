package z

import (
	"fmt"
	"testing"

	"github.com/davecgh/go-spew/spew"

	"github.com/vhaoran/yi/cmn"
	"github.com/vhaoran/yi/service"
)

func Test_get_na_yin(t *testing.T) {
	for y := 1970; y <= 2030; y++ {
		//d := calendar.NewSolarFromYmd(y,5,1)
		g, z := cmn.GetNianGanZhi(y)
		naYin := new(NaYinGet).Get(g + z)
		wuxing := new(NaYinGet).GetWuXin(naYin)
		fmt.Println(y, "---", g+z, ",", naYin, wuxing)
	}
}

func Test_bz_get(t *testing.T) {
	y, m, d, h, minute := 1970, 12, 14, 4, 50
	bz := BaZiGetX.FromSolar(y, m, d, h, minute)
	//
	fmt.Println("-----------------")

	spew.Dump(bz)
	fmt.Println("-----------------")

	l := service.GetJie(1970)
	for i, v := range l {
		fmt.Println(i+1, " --- ", v.ToString())
	}
	fmt.Println("-----------------")
	nianShu, yueShu := new(QiYunGet).Call(bz)
	fmt.Println("起运", nianShu, " 年", yueShu, "月")
	fmt.Println("-----------------")
	//--------guiren -----------------------------
	lguiren := new(GuiRenGet).Call(bz)
	for _, v := range lguiren {
		fmt.Println(v.ToString())
	}

	//--------xiongsheng -----------------------------
	lXiongShen := new(XiongShenGet).Call(bz)
	for _, v := range lXiongShen {
		fmt.Println(v.ToString())
	}

}
