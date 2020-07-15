package li

import (
	"strconv"
	"strings"

	"github.com/vhaoran/vchat/common/ytime"
)

const LEAP_STR = "ezc|esg|wog|gr9|15k0|16xc|1yl0|h40|ukw|gya|esg|wqe|wk0|15jk|2k45|zsw|16e8|yaq|tkg|1t2v|ei8|wj4|zp1|l00|lkw|2ces|8kg|tio|gdu|ei8|k12|1600|1aa8|lud|hxs|8kg|257n|t0g|2i8n|13rk|1600|2ld2|ztc|h40|2bas|7gw|t00|15ma|xg0|ztj|lgg|ztc|1v11|fc0|wr4|1sab|gcw|xig|1a34|l28|yhy|xu8|ew0|xr8|wog|g9s|1bvn|16xc|i1j|h40|tsg|fdh|es0|wk0|161g|15jk|1654|zsw|zvk|284m|tkg|ek0|xh0|wj4|z96|l00|lkw|yme|xuo|tio|et1|ei8|jw0|n1f|1aa8|l7c|gxs|xuo|tsl|t0g|13s0|16xg|1600|174g|n6a|h40|xx3|7gw|t00|141h|xg0|zog|10v8|y8g|gyh|exs|wq8|1unq|gc0|xf4|nys|l28|y8g|i1e|ew0|wyu|wkg|15k0|1aat|1640|hwg|nfn|tsg|ezb|es0|wk0|2jsm|15jk|163k|17ph|zvk|h5c|gxe|ek0|won|wj4|xn4|2dsl|lk0|yao"

var (
	tg = []string{"甲", "乙", "丙", "丁", "戊", "己", "庚", "辛", "壬", "癸"}
	dz = []string{"子", "丑", "寅", "卯", "辰", "巳", "午", "未", "申", "酉", "戌", "亥"}
)

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
func GetLeapMonth(y int) int64 {
	l := leapList()
	return l[y-1900] & 0xf
}

//返回公历年份的闰月天数
func GetLeapDays(y int) int64 {
	l := leapList()

	if GetLeapMonth(y) > 0 {
		if l[y-1900]&0x10000 > 0 {
			return 30
		}
		return 29
	}
	return 0
}

// 返回农历月份天数
func GetDaysByLunarMonth(y, m int64) int64 {
	l := leapList()
	if l[y-1900] > 0 && (0x10000>>m) > 0 {
		return 30
	}
	return 29
}

// 返回公历月份天数
func GetDaysByMonth(y, m int) int {
	y1, m1 := y, m+1
	if m == 12 {
		y1, m1 = y+1, 1
	}

	t0 := ytime.OfInt(y, m, 1).Time
	t1 := ytime.OfInt(y1, m1, 1).Time

	d := t1.Sub(t0)
	return int(d.Hours() / 24)
}

// 返回公历年份天数
func GetDaysByYear(y int) int {
	y1 := y + 1

	t0 := ytime.OfInt(y, 1, 1).Time
	t1 := ytime.OfInt(y1, 1, 1).Time

	d := t1.Sub(t0)
	return int(d.Hours() / 24)
}

// 根据序号返回干支组合名
func Cyclical(n int) string {
	return tg[n%10] + dz[n%12]
}

// 根据公历日期返回农历日期
func toLunar(Y, M, D int) {
m := 1900 //起始年份
n := 0
d := (new Date(Y, M - 1, D) - new Date(1900, 0, 31)) / 86400000,
//起始date
leap := getLeapMonth(Y)
//当年闰月
isleap := false
//标记闰月
_y;

for(; m < 2050 && d > 0; m++) {
n = getDaysByYear(m)
d -= n
}

if(d < 0) {
d += n
m--
}

_y = m

for(m = 1; m < 13 && d > 0; m++) {

if(leap > 0 && m == leap + 1 && isleap === false){
--m
isleap = true
n = getLeapDays(_y)
} else {
n = getDaysByLunarMonth(_y, m)
}

if(isleap == true && m == (leap + 1)) isleap = false

d -= n
}

if(d == 0 && leap > 0 && m == leap + 1 && !isleap) --m

if(d < 0){
d += n
--m
}

//修正闰月下一月第一天为非闰月
if(d == 0) isleap = m == leap

//转换日期格式为1开始
d = d + 1

var _fixDate = fixResult(LUNAR.fixDate,
Y, M, D,
// BUG?
Y - ( M < m ? 1 : 0),   //如果公历月份小于农历就是跨年期，农历年份比公历-1
m, d);

return {
cy: _fixDate.y,
cm: _fixDate.m,
cd: _fixDate.d,
CM: (isleap ? "闰" : "") + ((_fixDate.m > 9 ? '十' : '') + LUNAR.c1[_fixDate.m%10]).replace('十二','腊').replace(/^一/,'正') + '月',
CD: {'10': '初十', '20': '二十', '30': '三十'}[_fixDate.d] || (LUNAR.c2[Math.floor(_fixDate.d/10)] + LUNAR.c1[~~_fixDate.d%10]),
isleap: isleap
}
}

