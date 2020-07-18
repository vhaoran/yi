package model

type (
	SiZhuModel struct {
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

func (r *SiZhuModel) ZhiList() []string {
	return []string{r.NianZhi, r.YueZhi, r.RiZhi, r.ShiZhi}
}
