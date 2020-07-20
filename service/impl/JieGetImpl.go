package impl

import (
	"strings"

	"github.com/6tail/lunar-go/calendar"

	"github.com/vhaoran/yi/service"
)

type (
	JieGetImpl struct {
	}
)

const jie_STR = "立春|惊蛰|清明|立夏|芒种|小暑|立秋|白露|寒露|立冬|大雪"

func (r *JieGetImpl) Exec(lunarYear int) []*service.JieData {
	ret := r.getExcludeXiaoHan(lunarYear)

	//小寒
	lunar := calendar.NewLunarFromYmd(lunarYear+1, 6, 6)
	m := lunar.GetJieQiTable()
	if solar, ok := m["小寒"]; ok {
		ret = append(ret,
			&service.JieData{
				Name: "小寒",
				Date: solar.GetCalendar(),
			})
	}

	return ret
}

func (r *JieGetImpl) getExcludeXiaoHan(lunarYear int) []*service.JieData {
	lunar := calendar.NewLunarFromYmd(lunarYear, 6, 6)
	m := lunar.GetJieQiTable()
	ret := make([]*service.JieData, 0)

	delete(m, "小寒")

	//
	names := strings.Split(jie_STR, "|")
	for _, v := range names {
		if solar, ok := m[v]; ok {
			ret = append(ret,
				&service.JieData{
					Name: v,
					Date: solar.GetCalendar(),
				})
		}
	}
	//
	return ret
}
