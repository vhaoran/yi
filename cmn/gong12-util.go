package cmn

import (
	"strings"
)

const gong_12_str = "生沐冠临帝衰病死墓绝胎养"

type RXGong12 int

const (
	//生
	RXGong12_sheng = 0
	//沐
	RXGong12_muYu = 1
	//冠
	RXGong12_guan = 2
	//临
	RXGong12_lin = 3
	//帝
	RXGong12_di = 4
	//衰
	RXGong12_shuai = 5
	//病
	RXGong12_bing = 6
	//死
	RXGong12_si = 7
	//墓
	RXGong12_mu = 8
	//绝
	RXGong12_jue = 9
	//胎
	RXGong12_tai = 10
	//养
	RXGong12_yang = 11
)

func Gong12List() []string {
	l := strings.Split(gong_12_str, "")
	return l
}

func Gong12Idx(gong string) int {
	for i, v := range Gong12List() {
		if v == gong {
			return i
		}
	}
	return -1
}

func Gong12FullStr(gongShort string) string {
	m := KV{
		"生": "长生",
		"沐": "沐浴",
		"冠": "冠帶",
		"临": "临官",
		"帝": "帝旺",
	}
	s, ok := m[gongShort]
	if ok {
		return s
	}

	return gongShort
}

func Gong12Str(index RXGong12) string {
	return Gong12List()[index]
}

func GetGong12OfGan(gan, zhi string) (gong RXGong12, gongName string) {
	gong, gongName = -1, ""
	//子丑寅卯辰巳午未申酉戌亥
	//甲乙丙丁戊己庚辛壬癸
	m := KV{
		//     生沐冠临帝衰病死墓绝胎养"
		"甲": "亥子丑寅卯辰巳午未申酉戌",
		"乙": "午巳辰卯寅丑子亥戌酉申未",
		"丙": "寅卯辰巳午未申酉戌亥子丑",
		"丁": "酉申未午巳辰卯寅丑子亥戌",
		"戊": "寅卯辰巳午未申酉戌亥子丑",
		"己": "酉申未午巳辰卯寅丑子亥戌",
		"庚": "巳午未申酉戌亥子丑寅卯辰",
		"辛": "子亥戌酉申未午巳辰卯寅丑",
		"壬": "申酉戌亥子丑寅卯辰巳午未",
		"癸": "卯寅丑子亥戌酉申未午巳辰",
	}

	s, ok := m[gan]
	if !ok {
		return -1, ""
	}

	//
	l := strings.Split(s, "")
	i := StrIndexOfSlice(zhi, l...)
	if i >= 0 && i < 12 {
		gong = RXGong12(i)
		gongName = Gong12Str(gong)
	}
	//
	return
}
