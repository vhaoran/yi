package service

import (
	"fmt"
	"time"

	"github.com/vhaoran/yi/service/impl"
)

type (
	JieGetS interface {
		Exec(lunarYear int) []*JieData
	}

	JieData struct {
		Name string
		Date time.Time
	}
)

//小han调整过的节
func GetJie(solarYear int) []*JieData {
	return new(impl.JieGetImpl).Exec(solarYear)
}

//得到立春的时间
func GetJieLiChun(solarYear int) *JieData {
	l := GetJie(solarYear)
	for _, v := range l {
		if v.Name == "立春" {
			return v
		}
	}
	return nil
}

//返回的节气为
func GetJieForLunarMonthN(solarYear int, solarMonth int) *JieData {
	l := GetJie(solarYear)
	if solarMonth >= 0 && solarMonth <= 12 {
		return l[solarMonth-1]
	}

	return nil
}

func (r *JieData) ToString() string {
	t := "2006-01-02 15:04:05"
	str := r.Date.Format(t)
	s := fmt.Sprintf("%s:%s", r.Name, str)
	return s
}
