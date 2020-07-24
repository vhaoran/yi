package cmn

import (
	"strings"
)

func GanChong(gan1, gan2 string) bool {
	//甲庚相冲、乙辛相冲、壬丙相冲、癸丁相冲
	l := []string{"甲庚", "乙辛", "壬丙", "癸丁"}

	for _, v := range l {
		if strings.Contains(v, gan1) &&
			strings.Contains(v, gan1) {
			return true
		}
	}

	return false
}

func ZhiChong(zhi1, zhi2 string) bool {
	////
	l := []string{"子午", "卯酉", "辰戌", "丑未", "寅申", "巳亥"}

	for _, v := range l {
		if strings.Contains(v, zhi1) &&
			strings.Contains(v, zhi2) {
			return true
		}
	}

	return false
}

func FanYinSlice(ganZhi string, lGan ...string) bool {
	for _, v := range lGan {
		if FanYin(ganZhi, v) {
			return true
		}
	}
	return false
}

func FuYinSlice(ganZhi string, lGanZhi ...string) bool {
	for _, v := range lGanZhi {
		if FuYin(ganZhi, v) {
			return true
		}
	}
	return false
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
