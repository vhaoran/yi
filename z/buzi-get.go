package z

import (
	"time"

	"github.com/6tail/lunar-go/calendar"
	"github.com/vhaoran/vchat/common/ytime"

	"github.com/vhaoran/yi/cmn"
	"github.com/vhaoran/yi/service"
	. "github.com/vhaoran/yi/z/model"
)

type (
	BaZiGet struct {
	}
)

var BaZiGetX = new(BaZiGet)

func (r *BaZiGet) GetBySolar(Y, M, D, H, minutes int) *SiZhuModel {
	// solor := calendar.NewSolar(Y, M, D, H, minutes, 0)
	t0 := ytime.OfInt(Y, M, D, H, minutes, 0).Time
	lunar := calendar.NewLunarFromDate(t0)

	//
	z := &SiZhuModel{
		YYYY:   Y,
		MM:     M,
		DD:     D,
		HH:     H,
		Minute: minutes,
		Nian:   lunar.GetYear(),
		Yue:    lunar.GetMonth(),
		Ri:     lunar.GetDay(),
		Shi:    "",
	}
	if r.tran(z) {
		return z
	}

	return nil
}

func (r *BaZiGet) GetByLunar() {

}

func (r *BaZiGet) tran(z *SiZhuModel) bool {
	// set NianGaoZi
	r.SetNianGanZhi(z)
	r.SetYueGanZhi(z)
	SetRiGanZhi(z)
	SetShiGanZhi(z)
}

//看时间是否在月的分界点上
func (r *BaZiGet) SetYueGanZhi(z *SiZhuModel) {
	//主要判断在节前还是节后的日期
	l := service.GetJie(z.Nian)
	//
	t := z.Solar()
	yue := r.locateYue(l, t)
	//

}

func (r *BaZiGet) locateYue(l []*service.JieData, t time.Time) (month int) {
	//排好序的24节气
	for i, v := range l {
		if i < len(l)-1 {
			if t.Before(v.Date) && t.After(l[i+1].Date) {
				return i
			}
		}
	}

	//------if not found------------

	//
}

func (r *BaZiGet) SetNianGanZhi(z *SiZhuModel) {
	// 主要判断时间在立春前还是后
	lichunSolar := service.GetJieLiChun(z.Nian)

	//日期早于立春，则取上年的年十支
	if z.Solar().Before(lichunSolar.Date) {
		z.NianGan = cmn.GetYearGan(z.Nian - 1)
		z.NianZhi = cmn.GetYearZhi(z.Nian - 1)
		return
	}
	//
	z.NianGan = cmn.GetYearGan(z.Nian)
	z.NianZhi = cmn.GetYearZhi(z.Nian)
}
