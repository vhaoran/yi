package service

import (
	"strings"

	"github.com/vhaoran/yi/cmn"
)

var (
	//年上启月
	nianShangQiYue_map = cmn.KV{
		//甲乙丙丁戊己庚辛壬癸}
		"甲": "丙丁戊己庚辛壬癸甲乙丙丁",
		"己": "丙丁戊己庚辛壬癸甲乙丙丁",
		"乙": "戊己庚辛壬癸甲乙丙丁戊己",
		"庚": "戊己庚辛壬癸甲乙丙丁戊己",
		"丙": "庚辛壬癸甲乙丙丁戊己庚辛",
		"辛": "庚辛壬癸甲乙丙丁戊己庚辛",
		"丁": "壬癸甲乙丙丁戊己庚辛壬癸",
		"壬": "壬癸甲乙丙丁戊己庚辛壬癸",
		"戊": "甲乙丙丁戊己庚辛壬癸甲乙",
		"癸": "甲乙丙丁戊己庚辛壬癸甲乙",
	}
)

//尽量不使用库
func GetYueGanZhi(lunarNian, lunarYue int) (gan, zhi string) {
	//
	nianGan, _ := cmn.GetNianGanZhi(lunarNian)
	//nianZhi := cmn.GetYearZhi(lunarNian)
	//FriendReqDispose
	s, ok := nianShangQiYue_map[nianGan]
	if !ok {
		//
	}

	//
	l := strings.Split(s, "")
	gan = l[lunarYue-1]
	//
	yueList := []string{"寅", "卯", "辰", "巳", "午", "未", "申", "酉", "戌", "亥", "子", "丑"}
	zhi = yueList[lunarYue-1]
	return
}
