package y

import (
	"strings"
)

//
var YAO_LIUSHEN = map[string]string{
	//子丑寅卯辰巳午未申酉戌亥]
	"甲": "青龙|朱雀|勾陈|滕蛇|白虎|玄武",
	"乙": "青龙|朱雀|勾陈|滕蛇|白虎|玄武",
	"丙": "朱雀|勾陈|滕蛇|白虎|玄武|青龙",
	"丁": "朱雀|勾陈|滕蛇|白虎|玄武|青龙",
	//---不同-<<<
	"戊": "勾陈|滕蛇|白虎|玄武|青龙|朱雀",
	//---不同->>>
	"己": "滕蛇|白虎|玄武|青龙|朱雀|勾陈",
	"庚": "滕蛇|白虎|玄武|青龙|朱雀|勾陈",
	"辛": "滕蛇|白虎|玄武|青龙|朱雀|勾陈",
	"壬": "玄武|青龙|朱雀|勾陈|滕蛇|白虎",
	"癸": "玄武|青龙|朱雀|勾陈|滕蛇|白虎",
}

//输出按从六爻到一多的顺序列出
//出0：6爻  1：2爻。。。。5：初爻
func GetLiuShen(gan string) []string {
	s, ok := YAO_LIUSHEN[gan]
	//
	if !ok {
		return nil
	}
	//
	l := strings.Split(s, "|")
	for i := 0; i < len(l)/2; i++ {
		l[i], l[len(l)-1-i] = l[len(l)-1-i], l[i]
	}
	return l
}
