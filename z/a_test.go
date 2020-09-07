package z

import (
	"fmt"
	"testing"

	"github.com/davecgh/go-spew/spew"

	"github.com/vhaoran/yi/cmn"
	"github.com/vhaoran/yi/service"
)

func Test_fuyin(t *testing.T) {
	//子丑寅卯辰巳午未申酉戌亥
	//甲乙丙丁戊己庚辛壬癸

	a, str := cmn.FuYinSlice("庚寅", "庚戌", "戊子", "戊辰", "甲寅")
	fmt.Println(a, "---", str)
}

func Test_zhiChong(t *testing.T) {
	b := cmn.ZhiChong("寅", "寅")
	fmt.Println("-----------------")
	fmt.Println(b)
}

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
	y, m, d, h, minute := 1992, 5, 6, 7, 50
	z := BaZiGetX.FromSolar(y, m, d, h, minute)
	//
	fmt.Println("-----------------")

	spew.Dump(z)
	fmt.Println("-----------------",
		z.NianZhu(), " ",
		z.YueZhu(), " ",
		z.RiZhu(), " ",
		z.ShiZhu(), " ",
	)

	l := service.GetJie(y)
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

	//--------dayun -----------------------------
	fmt.Println("--------起大运---------")
	qiYunSuiShe := nianShu
	if yueShu > 6 {
		qiYunSuiShe += 1
	}

	z.IsMale = true
	l_PaiDaYunExec := new(PaiDaYunItemExec).Exec(z)
	for _, v := range l_PaiDaYunExec {
		fmt.Println(v.ToString())
		//fmt.Println("-----------------")
		//spew.Dump(v)
	}

	//-------- -----------------------------
}
