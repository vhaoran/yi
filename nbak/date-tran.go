package nbak

import (
	"github.com/vhaoran/vchat/common/ytime"
)

type (
	ToSolar struct {
	}
)

func (r *ToSolar) GetSolarDate(Y, M, D int) (y, m, d int) {
	y, m, d = 0, 0, 0
	//
	// 阴历日期=14Q+10.6(R+1)+年内日期序数-29.5n
	i := Y - 1901
	Q, R := i/4, i%4
	//
	n := r.DayNOfYear(Y, M, D)
	//
	z := float64(14*Q) + 10.6*float64(R+1) + float64(n)

	//
	m = int(z / 29.5)
	//
	d = int(z - float64(m)*29.5)
	return
}

//返回
func (r *ToSolar) DayNOfYear(Y, M, D int) int {
	t := ytime.OfInt(Y, M, D).Time
	return t.YearDay()
}
