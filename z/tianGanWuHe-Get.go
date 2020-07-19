package z

import (
	"github.com/vhaoran/yi/cmn"
	. "github.com/vhaoran/yi/z/model"
)

type (
	TiGanWuHe struct {
	}
)

var TiGanWuHeX = new(TiGanWuHe)

var tianGanWuHe = cmn.KV{
	"甲己": "土",
	"乙庚": "金",
	"丙辛": "水",
	"丁壬": "木",
	"戊癸": "火",
}

func (r *TiGanWuHe) Get(z *SiZhuModel) []*TianGanWuHeModel {
	l := make([]*TianGanWuHeModel, 0)
	//
	a := []string{z.NianGan, z.YueGan, z.RiGan, z.ShiGan}
	for k1, v1 := range a {
		for k2, v2 := range a {
			if k1 == k2 {
				continue
			}
			heWuXin := r.he(v1, v2)
			if len(heWuXin) == 0 {
				continue
			}
			l = append(l, &TianGanWuHeModel{
				He:         v1 + v2,
				HeHuaWuXin: heWuXin,
			})
		}
	}
	return l
}

func (r *TiGanWuHe) he(a, b string) string {
	str := a + b
	s, ok := tianGanWuHe[str]
	if ok {
		return s
	}
	return ""
}
