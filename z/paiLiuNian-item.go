package z

import (
	"fmt"

	"github.com/vhaoran/yi/cmn"
	. "github.com/vhaoran/yi/z/model"
)

type (
	PaiLiuNianExec struct {
	}

	LiuNianItem struct {
		Nian int

		Gan string
		Zhi string

		//
		JiShen    []*GuiRenItem
		XiongShen []*XiongShenItem
		FamYin    bool
		FuYin     bool
		//運上的流年解析
		AnalyzeInfo []string
	}
)

func (r *LiuNianItem) ToString() string {
	str := fmt.Sprint(r.Zhu(), "(", r.Nian, ")")

	s3 := "吉神:"
	for _, v := range r.JiShen {
		if s3 == "吉神:" {
			s3 = v.ToString()
			continue
		}
		s3 += "," + v.ToString()
	}

	//-------- -----------------------------
	s4 := "凶神:"
	for _, v := range r.JiShen {
		if s4 == "凶神:" {
			s4 = v.ToString()
			continue
		}
		s4 += "," + v.ToString()
	}
	s5 := " 特另注意:： "
	if r.FamYin {
		s5 = "/反吟"
	}
	if r.FuYin {
		s5 += "/伏吟"
	}

	return str + s3 + s4 + s5
}

func (r *LiuNianItem) Zhu() string {
	return r.Gan + r.Zhi
}

func (r *PaiLiuNianExec) Exec(nian int, z *SiZhuModel, yun *DaYunInfo) *LiuNianItem {
	gan, zhi := cmn.GetNianGanZhi(nian)

	bean := &LiuNianItem{
		Nian:      nian,
		Gan:       gan,
		Zhi:       zhi,
		JiShen:    nil,
		XiongShen: nil,

		FamYin: false,
		FuYin:  false,
	}
	//
	r.setJiSheng(bean, z, yun)
	r.setXiongShen(bean, z, yun)
	bean.FamYin = cmn.FanYinSlice(bean.Zhu(), z.ZhuList()...)
	bean.FuYin = cmn.FuYinSlice(bean.Zhu(), z.ZhuList()...)
	//

	return bean
}

func (r *PaiLiuNianExec) setJiSheng(bean *LiuNianItem, z *SiZhuModel, yun *DaYunInfo) {
	lGuiRen := r.getGuiRen(z, yun, bean)
	bean.JiShen = lGuiRen
}

func (r *PaiLiuNianExec) setXiongShen(bean *LiuNianItem, z *SiZhuModel, yun *DaYunInfo) {
	l := r.getXiongShen(z, yun, bean)
	bean.XiongShen = l
}

func (r *PaiLiuNianExec) getGuiRen(z *SiZhuModel, yun *DaYunInfo, n *LiuNianItem) []*GuiRenItem {
	ret := make([]*GuiRenItem, 0)

	obj := new(GuiRenGet)

	//-------- -----------------------------

	//-------- -----------------------------
	l := obj.TianYi(z.RiGan, z.NianZhi, z.YueZhi, z.RiZhi, z.ShiZhi, yun.Zhi, n.Zhi)
	if len(l) > 0 {
		ret = append(ret, l...)
	}
	l = obj.TianYi(z.NianGan, z.NianZhi, z.YueZhi, z.RiZhi, z.ShiZhi, yun.Zhi, n.Zhi)
	if len(l) > 0 {
		ret = append(ret, l...)
	}

	//-------- -----------------------------
	l = obj.ShiGanLu(z.RiGan, z.NianZhi, z.YueZhi, z.RiZhi, z.ShiZhi, yun.Zhi, n.Zhi)
	if len(l) > 0 {
		ret = append(ret, l...)
	}
	//-------- -----------------------------
	l = obj.WenChang(z.NianGan, z.NianZhi, z.YueZhi, z.RiZhi, z.ShiZhi, yun.Zhi, n.Zhi)
	if len(l) > 0 {
		ret = append(ret, l...)
	}
	l = obj.WenChang(z.RiGan, z.NianZhi, z.YueZhi, z.RiZhi, z.ShiZhi, yun.Zhi, n.Zhi)
	if len(l) > 0 {
		ret = append(ret, l...)
	}
	//-------- -----------------------------
	l = obj.HuaGai(z.NianZhi, z.YueZhi, z.RiZhi, z.ShiZhi, yun.Zhi, n.Zhi)
	if len(l) > 0 {
		ret = append(ret, l...)
	}
	l = obj.HuaGai(z.RiZhi, z.NianZhi, z.YueZhi, z.ShiZhi, yun.Zhi, n.Zhi)
	if len(l) > 0 {
		ret = append(ret, l...)
	}
	//-------- -----------------------------
	l = obj.HuaGai(z.RiZhi, z.NianZhi, z.YueZhi, z.ShiZhi, yun.Zhi, n.Zhi)
	if len(l) > 0 {
		ret = append(ret, l...)
	}
	//-------- -----------------------------
	l = obj.JiangXing(z.NianZhi, z.YueZhi, z.RiZhi, z.ShiZhi, yun.Zhi, n.Zhi)
	if len(l) > 0 {
		ret = append(ret, l...)
	}
	l = obj.JiangXing(z.RiZhi, z.NianZhi, z.YueZhi, z.ShiZhi, yun.Zhi, n.Zhi)
	if len(l) > 0 {
		ret = append(ret, l...)
	}
	//-------- -----------------------------
	//-------- -----------------------------
	l = obj.YiMa(z.NianZhi, z.YueZhi, z.RiZhi, z.ShiZhi, yun.Zhi, n.Zhi)
	if len(l) > 0 {
		ret = append(ret, l...)
	}
	l = obj.YiMa(z.RiZhi, z.NianZhi, z.YueZhi, z.ShiZhi, yun.Zhi, n.Zhi)
	if len(l) > 0 {
		ret = append(ret, l...)
	}
	//-------- -----------------------------
	l = obj.JinYu(z.RiGan, z.NianZhi, z.YueZhi, z.RiZhi, z.ShiZhi, yun.Zhi, n.Zhi)
	if len(l) > 0 {
		ret = append(ret, l...)
	}
	//-------- -----------------------------
	l = obj.TianYiGuiRen(z.YueZhi, z.NianZhi, z.YueZhi, z.RiZhi, z.ShiZhi, yun.Zhi, n.Zhi)
	if len(l) > 0 {
		ret = append(ret, l...)
	}
	//-------- -----------------------------
	return ret
}

func (r *PaiLiuNianExec) getXiongShen(z *SiZhuModel, yun *DaYunInfo, n *LiuNianItem) []*XiongShenItem {
	obj := new(XiongShenGet)
	//-------- -----------------------------
	ret := make([]*XiongShenItem, 0)
	if l := obj.YanRen(z.RiGan, z.NianZhi, z.YueZhi, z.RiZhi, z.ShiZhi, yun.Zhi, n.Zhi); len(l) > 0 {
		ret = append(ret, l...)
	}
	//-------- -----------------------------
	if l := obj.LiuJiaKongWang(z.RiGan+z.RiZhi, z.NianZhi, z.YueZhi, z.ShiZhi, yun.Zhi, n.Zhi); len(l) > 0 {
		ret = append(ret, l...)
	}
	if l := obj.LiuJiaKongWang(z.NianGan+z.NianZhi, z.YueZhi, z.RiZhi, z.ShiZhi, yun.Zhi, n.Zhi); len(l) > 0 {
		ret = append(ret, l...)
	}
	//-------- -----------------------------
	if l := obj.TaoHua(z.RiZhi, z.NianZhi, z.YueZhi, z.ShiZhi, yun.Zhi, n.Zhi); len(l) > 0 {
		ret = append(ret, l...)
	}
	if l := obj.TaoHua(z.NianZhi, z.YueZhi, z.RiZhi, z.ShiZhi, yun.Zhi, n.Zhi); len(l) > 0 {
		ret = append(ret, l...)
	}

	//-------- -----------------------------
	if l := obj.TianLuoDiWang(z.RiZhi, z.NianZhi, z.YueZhi, z.ShiZhi, yun.Zhi, n.Zhi); len(l) > 0 {
		ret = append(ret, l...)
	}
	if l := obj.TianLuoDiWang(z.NianZhi, z.YueZhi, z.RiZhi, z.ShiZhi, yun.Zhi, n.Zhi); len(l) > 0 {
		ret = append(ret, l...)
	}
	//-------- -----------------------------

	return ret
}
