package z

import (
	"fmt"
	"strings"

	. "github.com/vhaoran/yi/z/model"
)

//根据八字排出来大运和流年
type (
	PaiDaYunItemExec struct {
		//起运虚岁
		QiYuSuiShu int
		//顺排
		IsAsc bool
	}

	DaYunItem struct {
		//大运基本分析信彷
		DaYu *DaYunInfo
		//对应大运中的10个汉年
		LiuNian []*LiuNianItem
		//大運解析
		AnalyzeInfo []string
	}
)

func (r *DaYunItem) ToString() string {
	//--------liuNian -----------------------------
	l := ""
	for _, v := range r.LiuNian {
		if len(l) == 0 {
			l = v.ToString()
			continue
		}
		l += "," + v.ToString()
	}

	s := fmt.Sprintf("%s \n\r   %s \n\n  %s \n\n ",
		r.DaYu.ToString(),
		l,
		strings.Join(r.AnalyzeInfo, "\n\r"))

	return s
}

func (r *PaiDaYunItemExec) Exec(z *SiZhuModel) []*DaYunItem {
	ret := make([]*DaYunItem, 0)

	//-------- -----------------------------
	r.qiYun(z)

	//-------- -----------------------------
	l := new(DaYunInfoGet).Exec(z, r.QiYuSuiShu)
	if len(l) == 0 {
		return nil
	}

	//-------- -----------------------------
	for _, v := range l {
		//linian,10个流年
		liuNian := r.getLiuNianItems(z, v)
		bean := &DaYunItem{
			DaYu:        v,
			LiuNian:     liuNian,
			AnalyzeInfo: nil,
		}
		ret = append(ret, bean)

		//设置 AnalyzeInfo
		r.analyze(z, bean)
	}

	return ret
}

func (r *PaiDaYunItemExec) getLiuNianItems(z *SiZhuModel, yun *DaYunInfo) []*LiuNianItem {
	ret := make([]*LiuNianItem, 0)
	for n := yun.NianStart; n < yun.NianEnd; n++ {
		item := new(PaiLiuNianExec).Exec(n, z, yun)
		ret = append(ret, item)
	}
	return ret
}

func (r *PaiDaYunItemExec) qiYun(z *SiZhuModel) {
	//
	n, y := new(QiYunGet).Call(z)
	r.IsAsc = new(QiYunGet).IsAsc(z)
	//
	if y > 6 {
		r.QiYuSuiShu = n + 1
		return
	}

	r.QiYuSuiShu = n + 1
}

func (r *PaiDaYunItemExec) analyze(z *SiZhuModel, bean *DaYunItem) {

}
