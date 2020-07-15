package li

import (
	"strconv"
	"strings"
	"time"
	"github.com/vchaoran/vchat/lib/ytime"
)



const LEAP_STR = "ezc|esg|wog|gr9|15k0|16xc|1yl0|h40|ukw|gya|esg|wqe|wk0|15jk|2k45|zsw|16e8|yaq|tkg|1t2v|ei8|wj4|zp1|l00|lkw|2ces|8kg|tio|gdu|ei8|k12|1600|1aa8|lud|hxs|8kg|257n|t0g|2i8n|13rk|1600|2ld2|ztc|h40|2bas|7gw|t00|15ma|xg0|ztj|lgg|ztc|1v11|fc0|wr4|1sab|gcw|xig|1a34|l28|yhy|xu8|ew0|xr8|wog|g9s|1bvn|16xc|i1j|h40|tsg|fdh|es0|wk0|161g|15jk|1654|zsw|zvk|284m|tkg|ek0|xh0|wj4|z96|l00|lkw|yme|xuo|tio|et1|ei8|jw0|n1f|1aa8|l7c|gxs|xuo|tsl|t0g|13s0|16xg|1600|174g|n6a|h40|xx3|7gw|t00|141h|xg0|zog|10v8|y8g|gyh|exs|wq8|1unq|gc0|xf4|nys|l28|y8g|i1e|ew0|wyu|wkg|15k0|1aat|1640|hwg|nfn|tsg|ezb|es0|wk0|2jsm|15jk|163k|17ph|zvk|h5c|gxe|ek0|won|wj4|xn4|2dsl|lk0|yao"

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
func getDaysByMonth(y, m int64) int64 {
	y1, m1 := y, m+1
	if m == 12 {
       y1 = y +1
       m = 1
	}

	t0 := ytime.

}
