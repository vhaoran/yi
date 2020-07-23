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
		naYin := new(NaYinGet).GetSingle(g + z)
		wuxing := new(NaYinGet).GetWuXin(naYin)
		fmt.Println(y, "---", g+z, ",", naYin, wuxing)
	}
}

func Test_bz_get(t *testing.T) {
	y, m, d, h, minute := 1970, 12, 14, 4, 50
	z := BaZiGetX.FromSolar(y, m, d, h, minute)
	//
	fmt.Println("-----------------")

	spew.Dump(z)
	fmt.Println("-----------------")

	l := service.GetJie(1970)
	for i, v := range l {
		fmt.Println(i+1, " --- ", v.ToString())
	}

	//-------- -----------------------------
	fmt.Println("--------地支藏干---------")
	l_DiZhiCanGanGet := DiZhiCanGanGetX.Get(z)
	for _, v := range l_DiZhiCanGanGet {
		fmt.Print(v.ToString(), ",")
	}
	fmt.Println("")

	//--------纳音-----------------------------
	fmt.Println("-------纳音----------")
	l_NaYinGetX := NaYinGetX.Get(z)
	for _, v := range l_NaYinGetX {
		fmt.Print(v.ToString(), ",")
	}
	fmt.Println("")

	//---天干十神---------------------------------------
	fmt.Println("-----天干十神:------------")
	sheng10 := ShiShenGetX.Get(z)
	fmt.Println(sheng10.ToString())
	//--------------------------------------------
	fmt.Println("-----地支合化:------------")
	l_hehu := DiZhiHeHuiGetX.Get(z)
	for _, v := range l_hehu {
		fmt.Print(v.ToString(), ",")
	}
	fmt.Println("")
	//--------------------------------------------
	fmt.Println("-----地支刑冲:------------")
	l_DiZiXingChongHaiGetX := DiZiXingChongHaiGetX.Get(z)
	for _, v := range l_DiZiXingChongHaiGetX {
		fmt.Print(v, ",")
	}
	fmt.Println("")

	fmt.Println("------起运-----------")
	nianShu, yueShu := new(QiYunGet).Call(z)
	fmt.Println(nianShu, " 年", yueShu, "月")
	fmt.Println("-----------------")
	//--------guiren -----------------------------
	fmt.Println("--------吉神---------")
	lguiren := new(GuiRenGet).Call(z)
	for _, v := range lguiren {
		fmt.Println(v.ToString())
	}

	//--------xiongsheng -----------------------------
	fmt.Println("--------凶神---------")
	lXiongShen := new(XiongShenGet).Call(z)
	for _, v := range lXiongShen {
		fmt.Println(v.ToString())
	}

}
