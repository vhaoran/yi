package model

import (
	"time"

	"github.com/vhaoran/vchat/common/ytime"
)

type (
	SiZhuModel struct {
		//男性
		IsMale bool `json:"is_male"`
		//solar
		YYYY   int `json:"year"`
		MM     int `json:"month"`
		DD     int `json:"day"`
		HH     int `json:"hour"`
		Minute int `json:"minute"`

		//lunar
		Nian int    `json:"nian"`
		Yue  int    `json:"yue"`
		Ri   int    `json:"ri"`
		Shi  string `json:"shi"`

		//gan/zhi
		NianGan string `json:"nian_gan"`
		NianZhi string `json:"nian_zhi"`
		YueGan  string `json:"yue_gan"`

		YueZhi string `json:"yue_zhi"`
		RiGan  string `json:"ri_gan"`
		RiZhi  string `json:"ri_zhi"`
		ShiGan string `json:"shi_gan"`
		ShiZhi string `json:"shi_zhi"`
	}
)

func (r *SiZhuModel) NianZhu() string {
	return r.NianGan + r.NianZhi
}
func (r *SiZhuModel) YueZhu() string {
	return r.YueGan + r.YueZhi
}
func (r *SiZhuModel) RiZhu() string {
	return r.RiGan + r.RiZhi
}
func (r *SiZhuModel) ShiZhu() string {
	return r.ShiGan + r.ShiZhi
}

func (r *SiZhuModel) GanList() []string {
	return []string{r.NianGan, r.YueGan, r.RiGan, r.ShiGan}
}

func (r *SiZhuModel) ZhuList() []string {
	return []string{r.NianZhu(), r.YueZhu(), r.RiZhu(), r.ShiZhu()}
}

func (r *SiZhuModel) Solar() time.Time {
	return ytime.OfInt(r.YYYY, r.MM, r.DD, r.HH, r.Minute, 0).Time
}

func (r *SiZhuModel) ZhiList() []string {
	return []string{r.NianZhi, r.YueZhi, r.RiZhi, r.ShiZhi}
}
