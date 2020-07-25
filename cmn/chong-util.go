package cmn

import (
	"strings"

	"github.com/vhaoran/vchat/common/g"
)

func GanChong(gan1, gan2 string) bool {
	s1 := gan1 + gan2
	s2 := gan2 + gan2

	//甲庚相冲、乙辛相冲、壬丙相冲、癸丁相冲
	l := []string{"甲庚", "乙辛", "壬丙", "癸丁"}

	return g.InSlice(s1, l) || g.InSlice(s2, l)
}

func ZhiChong(zhi1, zhi2 string) bool {
	s1, s2 := zhi1+zhi2, zhi2+zhi1

	////
	l := []string{"子午", "卯酉", "辰戌", "丑未", "寅申", "巳亥"}

	return g.InSlice(s1, l) || g.InSlice(s2, l)
}

func FanYinSlice(ganZhi string, lGanZhi ...string) (fanYin bool, fanYinganZhi string) {
	for _, v := range lGanZhi {
		if FanYin(ganZhi, v) {
			fanYin, fanYinganZhi = true, v
			return
		}
	}
	return false, ""
}

func FuYinSlice(ganZhi string, lGanZhi ...string) (fuYin bool, fuYinGanZhi string) {
	for _, v := range lGanZhi {
		if FuYin(ganZhi, v) {
			fuYin, fuYinGanZhi = true, v
			return
		}
	}
	return false, ""
}

func FanYin(ganZhi1, ganZhi2 string) bool {
	src := strings.Split(ganZhi1, "")
	dst := strings.Split(ganZhi2, "")
	if len(src) < 2 || len(dst) < 2 {
		return false
	}

	return GanChong(src[0], dst[0]) && ZhiChong(src[1], dst[1])
}

func FuYin(ganZhi1, ganZhi2 string) bool {
	return ganZhi1 == ganZhi2
}
