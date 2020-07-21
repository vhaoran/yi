package cmn

import (
	"strings"

	"github.com/vhaoran/vchat/common/ytime"
)

var (
	//日上起时
	riShangQishi_list = KV{
		//甲乙丙丁戊己庚辛壬癸
		"甲": "甲乙丙丁戊己庚辛壬癸甲乙",
		"己": "甲乙丙丁戊己庚辛壬癸甲乙",
		"乙": "丙丁戊己庚辛壬癸甲乙丙丁",
		"庚": "丙丁戊己庚辛壬癸甲乙丙丁",
		"丙": "戊己庚辛壬癸甲乙丙丁戊己",
		"辛": "戊己庚辛壬癸甲乙丙丁戊己",
		"丁": "庚辛壬癸甲乙丙丁戊己庚辛",
		"壬": "庚辛壬癸甲乙丙丁戊己庚辛",
		"戊": "壬癸甲乙丙丁戊己庚辛壬癸",
		"癸": "壬癸甲乙丙丁戊己庚辛壬癸",
	}
)

//
func GetNianGanZhi(lunarYear int) (gan, zhi string) {
	if lunarYear < 1900 {
		return "", ""
	}

	g := []string{"甲", "乙", "丙", "丁", "戊", "己", "庚", "辛", "壬", "癸"}
	z := []string{"子", "丑", "寅", "卯", "辰", "巳", "午", "未", "申", "酉", "戌", "亥"}
	//1900 庚子年
	//     6
	//     0
	offset := lunarYear - 1900
	gIdx := (offset%10 + 6) % 10
	zIdx := offset % 12
	gan, zhi = g[gIdx], z[zIdx]
	return
}

func GetRiGanZhi(year, month, day int) (gan, zhi string) {
	if year < 1900 {
		return "", ""
	}
	//1900-03-01  癸酉
	//
	t0 := ytime.OfInt(1900, 3, 1, 0, 0, 0).Time
	t1 := ytime.OfInt(year, month, day, 0, 0, 0).Time
	//
	days := int64(t1.Sub(t0).Hours() / 24)

	g := []string{"甲", "乙", "丙", "丁", "戊", "己", "庚", "辛", "壬", "癸"}
	z := []string{"子", "丑", "寅", "卯", "辰", "巳", "午", "未", "申", "酉", "戌", "亥"}
	//1900 庚子年
	//     6
	//     0
	gIdx := (days + 9) % 10
	zIdx := (days + 9) % 12
	gan, zhi = g[gIdx], z[zIdx]
	return
}

func GetShiGan(riGan, shiZhi string) string {
	s, ok := riShangQishi_list[riGan]
	//
	if !ok {
		//do nothing
	}
	l := strings.Split(s, "")
	//
	i := ZhiIndex(shiZhi)
	//
	if i > 0 {
		return l[i]
	}
	return ""
}
