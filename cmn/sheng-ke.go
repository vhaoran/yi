package cmn

import (
	"strings"

	"github.com/vhaoran/vchat/common/g"
)

/* 五行生克关系 */
type RX_WuXing int

const (
	//生我
	RX_ShengWo = 1
	//我生
	RX_WoSheng = 2
	//克我
	RX_KeWo = 3
	//我克
	RX_WoKe = 4
	//同我（五行相同）
	RX_TongWo = 5
)

//得到天干 、地支的五行
func GetWuXing(ganOrZi string) string {
	//"金","木","水","火","土"
	m := KV{
		"金": "庚辛申酉",
		"木": "甲乙寅卯",
		"水": "壬癸子亥",
		"火": "丙丁巳午",
		"土": "戊己丑未辰戌",
	}

	for k, v := range m {
		if strings.Contains(v, ganOrZi) {
			return k
		}
	}
	return ""
}

//得到我与其它五行的
//我的天十或地支，其它的填或支
//isSame同属笥
func GetRX(woGanOrZhi, otherGaoOrZhi string) (rx RX_WuXing, isSame bool) {
	//五行属性
	w, o := GetWuXing(woGanOrZhi), GetWuXing(otherGaoOrZhi)
	//阴阳属性
	ws, bs := GetShuXing(woGanOrZhi), GetShuXing(otherGaoOrZhi)

	isSame, rx = ws == bs, 0

	//--------我生 -----------------------------
	//金木水火土
	woSheng := []string{"金水", "木火", "水木", "火土", "土金"}
	if g.InSlice(w+o, woSheng) {
		rx = RX_WoSheng
		return
	}
	//---生我 ---
	if g.InSlice(o+w, woSheng) {
		rx = RX_ShengWo
		return
	}
	//--------我克 -----------------------------
	//金木水火土
	woKe := []string{"金木", "木土", "土水", "水火", "火金"}
	if g.InSlice(w+o, woKe) {
		rx = RX_WoKe
		return
	}
	//---克我 ---
	if g.InSlice(o+w, woKe) {
		rx = RX_KeWo
		return
	}

	//---同我------------------------------
	if w == o {
		rx = RX_TongWo
		return
	}
	return
}
