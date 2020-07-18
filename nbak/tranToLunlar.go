package nbak

import (
	"github.com/vhaoran/vchat/common/ytime"
)

type (
	Tran2Lunar struct {
	}
)

var (
	FIXDATE_LIST = []string{"2013-1-1~2013-1-11=0|-1|1",
		"2013-1-12~2013-2-9=0|-1|0"}
)

func (r *Tran2Lunar) ToLunar(Y, M, D int) {
	m := 1900 //起始年份
	n := 0
	//d := (newDate(Y, M-1, D) - new Date(1900, 0, 31)) / 86400000
	d := r.days(Y, M, D)

	//起始date

	//当年闰月的月份
	leap := GetLeapMonth(Y)
	//标记闰月
	isleap := false
	//yy := false

	//-------- -----------------------------
	for m = 1900; m < 2050 && d > 0; m++ {
		n = GetDaysByYear(m)
		d -= n
	}
	if d < 0 {
		d += n
		m--
	}
	yy := m
	//-------- -------------------------
	for m = 1; m < 13 && d > 0; m++ {
		if leap > 0 && m == leap+1 && !isleap {
			m = m - 1
			isleap = true
			n = GetLeapDays(yy)
		} else {
			n = GetDaysByLunarMonth(yy, m)
		}

		if isleap && m == (leap+1) {
			isleap = false
		}
		d -= n
	}

	//-------- -----------------------------
	if d == 0 && leap > 0 && m == leap+1 && !isleap {
		m = m - 1
	}

	if d < 0 {
		d += n
		m = m - 1
	}

	//修正闰月下一月第一天为非闰月
	if d == 0 {
		isleap = m == leap
	}
	//转换日期格式为1开始
	d = d + 1
	//

}

func (r *Tran2Lunar) days(Y, M, D int) int {
	M = M - 1
	if M == 1 {
		Y, M = Y-1, 12
	}

	//
	t1 := ytime.OfInt(Y, M, D).Time
	t0 := ytime.OfInt(1900, 12, 31).Time

	d := t1.Sub(t0)
	return int(d.Hours() / 24)
}

//func fixResult(data []string, Y, M, D, y, m, d int) {
//	if data == nil && len(data) == 0 {
//		return
//	}
//
//	l := len(data)
//	match := func(y, m, d int, str []string) {
//		strList := strings.Split(str,"~")
//		strList[1] = strList[1] || strList[0]
//		pre := str[0].split("-")
//		suf := str[1].split("-")
//		return new
//		Date(y, m, d) >= new
//		Date(pre[0], pre[1], pre[2]) && new
//		Date(y, m, d) <= new
//		Date(suf[0], suf[1], suf[2])
//	}
//
//	val := 0
//	nbak := 0
//	while(l--) {
//		nbak = data[l].split("=")
//		val = nbak[1].split("|")
//		match(Y, M, D, nbak[0]) && (y = y + ~~(val[0]), m = m + ~~(val[1]), d = d + ~~(val[2]))
//	}
//
//	return
//	{
//	y:
//		~~y,
//			m: ~~m,
//		d: ~~d
//	}
//}
