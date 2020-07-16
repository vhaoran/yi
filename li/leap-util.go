package li

import (
	"strconv"
	"strings"
	"time"

	"github.com/vhaoran/vchat/common/ytime"
)

const LEAP_STR = "ezc|esg|wog|gr9|15k0|16xc|1yl0|h40|ukw|gya|esg|wqe|wk0|15jk|2k45|zsw|16e8|yaq|tkg|1t2v|ei8|wj4|zp1|l00|lkw|2ces|8kg|tio|gdu|ei8|k12|1600|1aa8|lud|hxs|8kg|257n|t0g|2i8n|13rk|1600|2ld2|ztc|h40|2bas|7gw|t00|15ma|xg0|ztj|lgg|ztc|1v11|fc0|wr4|1sab|gcw|xig|1a34|l28|yhy|xu8|ew0|xr8|wog|g9s|1bvn|16xc|i1j|h40|tsg|fdh|es0|wk0|161g|15jk|1654|zsw|zvk|284m|tkg|ek0|xh0|wj4|z96|l00|lkw|yme|xuo|tio|et1|ei8|jw0|n1f|1aa8|l7c|gxs|xuo|tsl|t0g|13s0|16xg|1600|174g|n6a|h40|xx3|7gw|t00|141h|xg0|zog|10v8|y8g|gyh|exs|wq8|1unq|gc0|xf4|nys|l28|y8g|i1e|ew0|wyu|wkg|15k0|1aat|1640|hwg|nfn|tsg|ezb|es0|wk0|2jsm|15jk|163k|17ph|zvk|h5c|gxe|ek0|won|wj4|xn4|2dsl|lk0|yao"
const JQMAP_STR = "0|gd4|wrn|1d98|1tuh|2akm|2rfn|38g9|3plp|46vz|4o9k|55px|5n73|64o5|6m37|73fd|7kna|81qe|8io7|8zgq|9g4b|9wnk|ad3g|ath2"
const JQNAMES_STR = "小寒|大寒|立春|雨水|惊蛰|春分|清明|谷雨|立夏|小满|芒种|夏至|小暑|大暑|立秋|处暑|白露|秋分|寒露|霜降|立冬|小雪|大雪|冬至"

var (
	tg = []string{"甲", "乙", "丙", "丁", "戊", "己", "庚", "辛", "壬", "癸"}
	dz = []string{"子", "丑", "寅", "卯", "辰", "巳", "午", "未", "申", "酉", "戌", "亥"}

	//节气
)

func jqMap() []int64 {
	l := strings.Split(JQMAP_STR, "|")
	if l == nil {
		return nil
	}

	list := make([]int64, len(l))
	for i, v := range l {
		j, err := strconv.ParseInt(v, 36, 64)
		if err != nil {
			panic("" + err.Error())
		}
		list[i] = j
	}

	return list
}

func leapList() []int64 {
	l := strings.Split(LEAP_STR, "|")
	if l == nil {
		return nil
	}

	list := make([]int64, len(l))
	for i, v := range l {
		j, err := strconv.ParseInt(v, 36, 64)
		if err != nil {
			panic("" + err.Error())
		}
		list[i] = j
	}

	return list
}

//返回公历年份的闰月月份
func GetLeapMonth(ySolar int) int {
	l := leapList()
	return int(l[ySolar-1900] & 0xf)
}

//返回公历年份的闰月天数
func GetLeapDays(ySolar int) int {
	l := leapList()

	if GetLeapMonth(ySolar) > 0 {
		if l[ySolar-1900]&0x10000 > 0 {
			return 30
		}
		return 29
	}
	return 0
}

// 返回农历月份天数
func GetDaysByLunarMonth(yLunar, m int) int {
	l := leapList()
	if l[yLunar-1900] > 0 && (0x10000>>m) > 0 {
		return 30
	}
	return 29
}

// 返回公历月份天数
func GetDaysByMonth(ySolar, m int) int {
	y1, m1 := ySolar, m+1
	if m == 12 {
		y1, m1 = ySolar+1, 1
	}

	t0 := ytime.OfInt(ySolar, m, 1).Time
	t1 := ytime.OfInt(y1, m1, 1).Time

	d := t1.Sub(t0)
	return int(d.Hours() / 24)
}

// 返回公历年份天数
func GetDaysByYear(ySolar int) int {
	y1 := ySolar + 1

	t0 := ytime.OfInt(ySolar, 1, 1).Time
	t1 := ytime.OfInt(y1, 1, 1).Time

	d := t1.Sub(t0)
	return int(d.Hours() / 24)
}

// 根据序号返回干支组合名
func Cyclical(n int) string {
	return tg[n%10] + dz[n%12]
}

//返回公历年份的第n个节气日期
func GetDateBySolar(y, jieN int) (year, month, day int) {
	offset := int64(31556925974.7*float64(y-1900)) + jqMap()[jieN]*60000
	t0 := ytime.OfInt(1900, 1, 6, 2, 5).Time
	d := t0.Add(time.Duration(offset) * time.Millisecond)

	return d.Year(), int(d.Month()), d.Day()
}

//根据农历月份，返回公历年、月、日->即24节所中规定的第一天，真正意义上的某一月
func GetDateByLunar(y, m int) (year, month, day int) {
	jieN := m * 2
	if m == 12 {
		y, jieN = y+1, 0
	}
	return GetDateBySolar(y, jieN)
}
