package li

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func Test_36_parse(t *testing.T) {
	s := "ezc|esg|wog|gr9|15k0|16xc|1yl0|h40|ukw|gya|esg|wqe|wk0|15jk|2k45|zsw|16e8|yaq|tkg|1t2v|ei8|wj4|zp1|l00|lkw|2ces|8kg|tio|gdu|ei8|k12|1600|1aa8|lud|hxs|8kg|257n|t0g|2i8n|13rk|1600|2ld2|ztc|h40|2bas|7gw|t00|15ma|xg0|ztj|lgg|ztc|1v11|fc0|wr4|1sab|gcw|xig|1a34|l28|yhy|xu8|ew0|xr8|wog|g9s|1bvn|16xc|i1j|h40|tsg|fdh|es0|wk0|161g|15jk|1654|zsw|zvk|284m|tkg|ek0|xh0|wj4|z96|l00|lkw|yme|xuo|tio|et1|ei8|jw0|n1f|1aa8|l7c|gxs|xuo|tsl|t0g|13s0|16xg|1600|174g|n6a|h40|xx3|7gw|t00|141h|xg0|zog|10v8|y8g|gyh|exs|wq8|1unq|gc0|xf4|nys|l28|y8g|i1e|ew0|wyu|wkg|15k0|1aat|1640|hwg|nfn|tsg|ezb|es0|wk0|2jsm|15jk|163k|17ph|zvk|h5c|gxe|ek0|won|wj4|xn4|2dsl|lk0|yao"
	l := strings.Split(s, "|")
	fmt.Println("-----------------")
	for j, v := range l {
		year := j + 1900
		if year < 1974 {
			continue
		}

		fmt.Println("-----", year, "->", v)
		i, _ := strconv.ParseInt(v, 36, 64)
		fmt.Println(i)
		fmt.Println("   ", year, "->闰", i&0xf, " : ",
			i&0x10000)
	}
	fmt.Println("-----------------")
}

func Test_get_leap_month(t *testing.T) {
	for i := 1974; i < 1990; i++ {
		j := GetLeapMonth(i)
		days := GetLeapDays(i)
		fmt.Println("------", i, "-----------")
		if j > 0 {
			fmt.Println(" run: ", j, " :", days)
		}
	}
}

func Test_GetDaysByMonth(t *testing.T) {
	for y := 2016; y < 2020; y++ {
		fmt.Println("-----------", y, "----------")
		for m := 1; m <= 12; m++ {
			fmt.Print(m, ":", GetDaysByMonth(y, m), "/")
		}
		fmt.Println("")
	}
}

func Test_jqMap(t *testing.T) {
	l := jqMap()
	//
	for k, v := range l {
		fmt.Println("----", k, ":", v*60000)
	}
}

func Test_GetDateBySolar(t *testing.T) {
	for year := 2000; year < 2020; year++ {
		fmt.Println("-----------", year, "----------")
		for m := 0; m < 12; m++ {
			yy, mm, dd := GetDateBySolar(year, m*2)
			s := fmt.Sprintf("(%d:%d-%d-%d) ", m, yy, mm, dd)
			fmt.Print(s)
		}
		fmt.Println("")
	}
}

func Test_GetDateByLunar(t *testing.T) {
	for year := 2000; year < 2020; year++ {
		fmt.Println("-----------", year, "----------")
		for m := 1; m < 13; m++ {
			yy, mm, dd := GetDateByLunar(year, m)
			s := fmt.Sprintf("(%d:%d-%d-%d) ", m, yy, mm, dd)
			fmt.Print(s)
		}
		fmt.Println("")
	}
}

func Test_ToSolar(t *testing.T) {
	y, m, d := new(ToSolar).GetSolarDate(2020, 7, 16)
	fmt.Println("-----------------")
	fmt.Println("------->---", y, " ", m, " ", d)

}
