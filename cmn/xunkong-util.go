package cmn

import (
	"strings"
)

var XUNKONG = KV{
	"甲子": "戌亥",
	"乙丑": "戌亥",
	"丙寅": "戌亥",
	"丁卯": "戌亥",
	"戊辰": "戌亥",
	"己巳": "戌亥",
	"庚午": "戌亥",
	"辛未": "戌亥",
	"壬申": "戌亥",
	"癸酉": "戌亥",


	"甲戌": "申酉",
	"乙亥": "申酉",
	"丙子": "申酉",
	"丁丑": "申酉",
	"戊寅": "申酉",
	"己卯": "申酉",
	"庚辰": "申酉",
	"辛巳": "申酉",
	"壬午": "申酉",
	"癸未": "申酉",

	//-----------------
	"甲申": "午未",
	"乙酉": "午未",
	"丙戌": "午未",
	"丁亥": "午未",
	"戊子": "午未",
	"己丑": "午未",
	"庚寅": "午未",
	"辛卯": "午未",
	"壬辰": "午未",
	"癸巳": "午未",

	//--------------------
	"甲午": "辰巳",
	"乙未": "辰巳",
	"丙申": "辰巳",
	"丁酉": "辰巳",
	"戊戌": "辰巳",
	"己亥": "辰巳",
	"庚子": "辰巳",
	"辛丑": "辰巳",
	"壬寅": "辰巳",
	"癸卯": "辰巳",

	//--------------------
	"甲辰": "寅卯",
	"乙巳": "寅卯",
	"丙午": "寅卯",
	"丁未": "寅卯",
	"戊申": "寅卯",
	"己酉": "寅卯",
	"庚戌": "寅卯",
	"辛亥": "寅卯",
	"壬子": "寅卯",
	"癸丑": "寅卯",

	//--------------------
	"甲寅": "子丑",
	"乙卯": "子丑",
	"丙辰": "子丑",
	"丁巳": "子丑",
	"戊午": "子丑",
	"己未": "子丑",
	"庚申": "子丑",
	"辛酉": "子丑",
	"壬戌": "子丑",
	"癸亥": "子丑",
}

func XuKong(ganZhi string) []string {
	//
	//甲乙丙丁戊己庚辛壬癸
	//子丑寅卯辰巳午未申酉戌亥
	//-------- -----------------------------
	//甲乙丙丁戊己庚辛壬癸
	//戌亥子丑寅卯辰巳午未申酉
	//-------- -----------------------------
	//甲乙丙丁戊己庚辛壬癸
	//申酉戌亥子丑寅卯辰巳午未
	//-------- -----------------------------
	//甲乙丙丁戊己庚辛壬癸
	//午未申酉戌亥子丑寅卯辰巳
	//-------- -----------------------------
	//甲乙丙丁戊己庚辛壬癸
	//辰巳午未申酉戌亥子丑寅卯
	//-------- -----------------------------
	//甲乙丙丁戊己庚辛壬癸
	//寅卯辰巳午未申酉戌亥子丑
	s, ok := XUNKONG[ganZhi]
	if ok {
		l := strings.Split(s, "")
		return l
	}

	return nil
}