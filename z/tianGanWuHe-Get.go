package z

import (
	. "github.com/vhaoran/yi/z/model"
)

type (
	TiGanWuHe struct {
	}
)

var TiGanWuHeX = new(TiGanWuHe)

var tianGanWuHe = KV{
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
	for k, src := range a {
		for k1, dst := range a {
			if k == k1 {
				continue
			}
			heWuXin := r.he(src, dst)
			if len(heWuXin) == 0 {
				continue
			}
			bean := &TianGanWuHeModel{
				He:         src + dst,
				HeHuaWuXin: heWuXin,
			}
			l = append(l, bean)
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
