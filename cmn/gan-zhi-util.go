package cmn

//
func GetYearGanZi(lunarYear int) string {
	if lunarYear < 1900 {
		return ""
	}

	g := []string{"甲", "乙", "丙", "丁", "戊", "己", "庚", "辛", "壬", "癸"}
	z := []string{"子", "丑", "寅", "卯", "辰", "巳", "午", "未", "申", "酉", "戌", "亥"}
	//1900 庚子年
	//     6
	//     0
	offset := lunarYear - 1900
	gi := (offset%10 + 6) % 10
	zi := offset % 12
	return g[gi] + z[zi]
}

func GetYearGan(lunarYear int) string {
	if lunarYear < 1900 {
		return ""
	}

	g := []string{"甲", "乙", "丙", "丁", "戊", "己", "庚", "辛", "壬", "癸"}
	//1900 庚子年
	//     6
	//     0
	offset := lunarYear - 1900
	gi := (offset%10 + 6) % 10
	return g[gi]
}
func GetYearZhi(lunarYear int) string {
	if lunarYear < 1900 {
		return ""
	}

	z := []string{"子", "丑", "寅", "卯", "辰", "巳", "午", "未", "申", "酉", "戌", "亥"}
	//1900 庚子年
	//     6
	//     0
	offset := lunarYear - 1900
	zi := offset % 12
	return z[zi]
}
