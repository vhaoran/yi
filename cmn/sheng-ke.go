package cmn

import (
	"strings"
)

func GetWuXing(ganOrZi string) string {
	//"金","木","水","火","土"
	m := KV{
		"甲": "木",
		"乙": "木",
		"寅": "木",
		"卯": "木",


		"丙": "火",
		"丁": "火",
		"巳": "火",
		"午": "火",

		"戊": "土",
		"己": "土",
		"丑": "土",
		"未": "土",
		"辰": "土",
		"戌": "土",

		"庚": "金",
		"辛": "金",
		"酉": "金",
		"申": "金",


		"壬": "水",
		"癸": "水",
		"子": "水",
		"亥": "水",
	}

	s, ok := m[ganOrZi]
	if ok {
		return s
	}
	return ""
}

//a,b必须同时是天幹或地支
func Sheng(a, b string) bool {
	xa, xb := GetWuXing(a), GetWuXing(b)

	//WuXing_list = []string{
	//"金", "木", "水", "火", "土"}
	m := KV{
		"金": "水",
		"木": "火",
		"水": "木",
		"火": "土",
		"土": "金",
	}

	for k, v := range m {
		if k == xa && v == xb {
			return true
		}
	}
	return false
}

func Ke(a, b string) bool {
	//WuXing_list = []string{"金", "木", "水", "火", "土"}
	xa, xb := GetWuXing(a), GetWuXing(b)

	m := KV{
		"金": "木",
		"木": "土",
		"水": "水",
		"火": "金",
		"土": "水",
	}

	for k, v := range m {
		if k == xa && v == xb {
			return true
		}
	}
	return false
}

func Bi(a, b string) bool {
	xa, xb := GetWuXing(a), GetWuXing(b)

	return strings.Compare(xa, xb) == 0
}
