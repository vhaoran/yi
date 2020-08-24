package z

import (
	"fmt"
	"math"
	"time"

	"github.com/6tail/lunar-go/calendar"
	"github.com/vhaoran/vchat/common/ytime"

	"github.com/vhaoran/yi/cmn"
	"github.com/vhaoran/yi/service"
	. "github.com/vhaoran/yi/z/model"
)

//----------------------------------------------------
// auth: whr  date:2020/7/2216:26--------------------------
// ####请勿擅改此功能代码####
// 用途：排八这入口，主要完成八字四柱的设置
//---------------------------------------------

type (
	BaZiGet struct {
	}
)

var BaZiGetX = new(BaZiGet)

func (r *BaZiGet) FromSolar(Y, M, D, H, minutes int) *SiZhuModel {
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

func (r *BaZiGet) FromLunar(nian, yue, ri, shi, fen int) *SiZhuModel {

	//
	var d = calendar.NewLunarFromYmd(nian, yue, ri)
	var solar = d.GetSolar()

	lunar := calendar.NewLunar(nian, yue, ri, shi, fen, 0)

	z := &SiZhuModel{
		YYYY:   solar.GetYear(),
		MM:     solar.GetMonth(),
		DD:     solar.GetDay(),
		HH:     solar.GetHour(),
		Minute: fen,

		Nian: lunar.GetYear(),
		Yue:  lunar.GetMonth(),
		Ri:   lunar.GetDay(),
		Shi:  "",
	}
	if r.tran(z) {
		return z
	}

	return nil
}

func (r *BaZiGet) tran(z *SiZhuModel) bool {
	// set NianGaoZi
	r.SetNianGanZhi(z)
	r.SetYueGanZhi(z)
	r.SetRiGanZhi(z)
	r.SetShiGanZhi(z)
	return true
}

//看时间是否在月的分界点上
func (r *BaZiGet) SetYueGanZhi(z *SiZhuModel) {
	//主要判断在节前还是节后的日期
	l := service.GetJie(z.Nian)
	//
	t := z.Solar()
	yue := r.locateYue(l, t)

	fmt.Println("设置月干支:yue", yue)

	//用上一年12月的干支，
	if yue == -1 {
		z.YueGan, z.YueZhi = service.GetYueGanZhi(z.Nian-1, 12)
		return
	}
	z.YueGan, z.YueZhi = service.GetYueGanZhi(z.Nian, yue)
}

func (r *BaZiGet) locateYue(l []*service.JieData, t time.Time) (month int) {
	//排好序的24节气
	for i, v := range l {
		if i < len(l)-1 {
			if t.After(v.Date) && t.Before(l[i+1].Date) {
				return i + 1
			}
		}
	}

	//------if not found------------
	if t.After(l[len(l)-1].Date) {
		return 12
	}

	//
	return -1
}

func (r *BaZiGet) SetNianGanZhi(z *SiZhuModel) {
	// 主要判断时间在立春前还是后
	lichunSolar := service.GetJieLiChun(z.Nian)

	//日期早于立春，则取上年的年十支
	if z.Solar().Before(lichunSolar.Date) {
		z.NianGan, z.NianZhi = cmn.GetNianGanZhi(z.Nian - 1)
		return
	}

	z.NianGan, z.NianZhi = cmn.GetNianGanZhi(z.Nian)
}

func (r *BaZiGet) SetRiGanZhi(z *SiZhuModel) {
	//主要判断日十支是否在节气上，如果在节的日上，则要看早于节还是晚于节
	//早于节取上日，晚于节的时，取当日
	//主要判断在节前还是节后的日期
	l := service.GetJie(z.Nian)
	//
	t := z.Solar()
	data := r.matchDay(l, t)

	//是否在节令转换的当日
	if data != nil {
		//在转换时辰之前，取上一日
		if t.Before(data.Date) {
			//取前一日
			next := z.Solar().Add(time.Hour * 24 * (-1))
			z.RiGan, z.RiZhi = cmn.GetRiGanZhi(next.Year(), int(next.Month()), next.Day())
			return
		}
	}

	//否则
	//如里时间>=23点，取下一日
	if z.HH >= 23 {
		//取後一日
		next := z.Solar().Add(time.Hour * 24 * (-1))
		z.RiGan, z.RiZhi = cmn.GetRiGanZhi(next.Year(), int(next.Month()), next.Day())
		return
	}

	//--------取当日-----------------------------
	z.RiGan, z.RiZhi = cmn.GetRiGanZhi(z.YYYY, z.MM, z.DD)
	return
}

func (r *BaZiGet) SetShiGanZhi(z *SiZhuModel) {
	//
	zhi := r.shiChenZhi(z.HH)
	//
	gan := cmn.GetShiGan(z.RiGan, zhi)
	z.ShiGan = gan
	z.ShiZhi = zhi
}

func (r *BaZiGet) shiChenZhi(hour int) string {
	h := hour
	//                 23-1   1-3   3-5  5-7   7-9   9-11 11-13 13-15 15-37
	Zhi_list := []string{"子", "丑", "寅", "卯", "辰", "巳", "午", "未", "申", "酉", "戌", "亥"}

	i := 0
	switch {
	case h >= 23 || h < 1:
		i = 0
	case h >= 1 && h < 3:
		i = 1
	case h >= 3 && h < 5:
		i = 2
	case h >= 5 && h < 7:
		i = 3
	case h >= 7 && h < 9:
		i = 4
	case h >= 9 && h < 11:
		i = 5
	case h >= 11 && h < 13:
		i = 6
	case h >= 13 && h < 15:
		i = 7
	case h >= 15 && h < 17:
		i = 8
	case h >= 17 && h < 19:
		i = 9
	case h >= 19 && h < 21:
		i = 10
	case h >= 21 && h < 23:
		i = 11
	default:
		i = 0
	}
	return Zhi_list[i]
}

func (r *BaZiGet) matchDay(l []*service.JieData, t time.Time) *service.JieData {
	for _, v := range l {
		offset := t.Sub(v.Date)

		if math.Abs(offset.Hours()) < 24 {
			return v
		}
	}

	return nil
}
