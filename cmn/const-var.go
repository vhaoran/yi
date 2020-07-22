package cmn


var (
	Gan_list    = []string{"甲", "乙", "丙", "丁", "戊", "己", "庚", "辛", "壬", "癸"}
	Zhi_list    = []string{"子", "丑", "寅", "卯", "辰", "巳", "午", "未", "申", "酉", "戌", "亥"}
	WuXing_list = []string{"金", "木", "水", "火", "土"}
)

type KV map[string]string

func GanIndex(gan string) int {
	for i, v := range Gan_list {
		if v == gan {
			return i
		}
	}
	return -1
}

func ZhiIndex(zhi string) int {
	for i, v := range Zhi_list {
		if v == zhi {
			return i
		}
	}

	return -1
}
