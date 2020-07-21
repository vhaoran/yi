package model

import (
	"time"

	"github.com/vhaoran/vchat/common/ytime"
)

type (
	SiZhuModel struct {
		//男性
		IsMale bool
		//solar
		YYYY   int
		MM     int
		DD     int
		HH     int
		Minute int

		//lunar
		Nian int
		Yue  int
		Ri   int
		Shi  string

		//gan/zhi
		NianGan string
		NianZhi string
		YueGan  string
		YueZhi  string
		RiGan   string
		RiZhi   string
		ShiGan  string
		ShiZhi  string
	}
)

func (r *SiZhuModel) GanList() []string {
	return []string{r.NianGan, r.YueGan, r.RiGan, r.ShiGan}
}

func (r *SiZhuModel) Solar() time.Time {
	return ytime.OfInt(r.YYYY, r.MM, r.DD, r.HH, r.Minute, 0).Time
}

func (r *SiZhuModel) ZhiList() []string {
	return []string{r.NianZhi, r.YueZhi, r.RiZhi, r.ShiZhi}
}
