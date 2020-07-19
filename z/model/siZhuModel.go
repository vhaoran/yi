package model

type (
	SiZhuModel struct {
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

func (r *SiZhuModel) ZhiList() []string {
	return []string{r.NianZhi, r.YueZhi, r.RiZhi, r.ShiZhi}
}
