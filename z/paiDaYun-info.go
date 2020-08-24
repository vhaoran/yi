package z

import (
	"fmt"

	"github.com/vhaoran/yi/cmn"
	. "github.com/vhaoran/yi/z/model"
)

type (
	DaYunInfoGet struct {
	}

	//主要用于显示
	DaYunInfo struct {
		//大支干支
		Gan string `json:"gan"`
		Zhi string `json:"zhi"`

		//大运名称
		Gong12 string `json:"gong12"`
		//大运简称
		Gong12Short string `json:"gong12_short"`

		//该运开始时年龄
		YearsOld int `json:"years_old"`

		//开始年
		NianStart int `json:"nian_start"`
		//结束年
		NianEnd int `json:"nian_end"`

		//
		JiShen    []*GuiRenItem    `json:"ji_shen"`
		XiongShen []*XiongShenItem `json:"xiong_shen"`
		FanYin    bool             `json:"fan_yin"`
		FuYin     bool             `json:"fu_yin"`
	}
)

func (r *DaYunInfo) Zhu() string {
	return r.Gan + r.Zhi
}

func (r *DaYunInfo) ToString() string {
	s := "-------- " + r.Zhu() + "(" + r.Gong12 + ") ---------"
	s1 := fmt.Sprint("(", r.NianStart, "-", r.NianEnd, ")")
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

	//-------- -----------------------------
	s5 := " 特另注意:： "
	if r.FanYin {
		s5 = "/反吟"
	}
	if r.FuYin {
		s5 += "/伏吟"
	}

	return s + s1 + s3 + s4
}

func (r *DaYunInfoGet) Exec(z *SiZhuModel, qiYunSuiShu int) []*DaYunInfo {
	isAsc := new(QiYunGet).IsAsc(z)
	l := r.calc(z.YueGan, z.YueZhi, qiYunSuiShu, z.Nian, isAsc)
	//子丑寅卯辰巳午未申酉戌亥
	//甲乙丙丁戊己庚辛壬癸
	//----------设置支中藏干------------------
	for _, v := range l {
		//设置12宫
		r.setGong12(z.RiGan, v)
	}

	//---------设置流年-----------------------

	return l
}

func (r *DaYunInfoGet) calc(yueGan, yueZhi string, qiYunYearsOld, bornNian int, isAsc bool) []*DaYunInfo {
	if isAsc {
		return r.calcShu(yueGan, yueZhi, qiYunYearsOld, bornNian)
	}

	return r.calcNi(yueGan, yueZhi, qiYunYearsOld, bornNian)
}

func (r *DaYunInfoGet) calcShu(yueGan, yueZhi string, qiYunYearsOld, bornNian int) []*DaYunInfo {
	ret := make([]*DaYunInfo, 0)

	//zhiStr := "子丑寅卯辰巳午未申酉戌亥"
	//ganStr := "甲乙丙丁戊己庚辛壬癸"
	ganIdx := cmn.GanIndex(yueGan)
	zhiIdx := cmn.ZhiIndex(yueZhi)

	//顺排 排100年运
	offset := qiYunYearsOld
	for c := 1; c < 11; c++ {
		i, j := ganIdx+c, zhiIdx+c
		gan, zhi := cmn.Gan_list[i%10], cmn.Zhi_list[j%12]
		bean := &DaYunInfo{
			Gan:         gan,
			Zhi:         zhi,
			Gong12:      "",
			Gong12Short: "",
			YearsOld:    offset,
			NianStart:   bornNian + offset,
			NianEnd:     bornNian + offset + 10,
		}

		offset += 10
		ret = append(ret, bean)
	}

	return ret
}

// 逆排
func (r *DaYunInfoGet) calcNi(yueGan, yueZhi string, qiYunYearsOld, bornNian int) []*DaYunInfo {
	ret := make([]*DaYunInfo, 0)

	//zhiStr := "子丑寅卯辰巳午未申酉戌亥"
	//ganStr := "甲乙丙丁戊己庚辛壬癸"
	ganIdx := cmn.GanIndex(yueGan)
	zhiIdx := cmn.ZhiIndex(yueZhi)

	//顺排 排100年运
	offset := qiYunYearsOld
	for c := 1; c < 11; c++ {
		i, j := ganIdx-c+10, zhiIdx-c+12
		gan, zhi := cmn.Gan_list[i%10], cmn.Zhi_list[j%12]
		bean := &DaYunInfo{
			Gan:         gan,
			Zhi:         zhi,
			Gong12:      "",
			Gong12Short: "",
			YearsOld:    offset,
			NianStart:   bornNian + offset,
			NianEnd:     bornNian + offset + 10,
		}

		offset += 10
		ret = append(ret, bean)
	}

	return ret
}

func (r *DaYunInfoGet) setShiShen(riGan string, v *DaYunInfo) {
	shiShen := cmn.GetShiShen(riGan, v.Gan)
	//
	v.Gong12 = shiShen
	v.Gong12Short = cmn.GetShiShenShort(v.Gong12)
}

func (r *DaYunInfoGet) setGong12(riGan string, v *DaYunInfo) {
	_, gong12 := cmn.GetGong12OfGan(riGan, v.Zhi)
	//
	v.Gong12 = cmn.Gong12FullStr(gong12)
	v.Gong12Short = gong12
}

func (r *DaYunInfoGet) setJiSheng(bean *DaYunInfo, z SiZhuModel) {
	lGuiRen := r.getGuiRen(z, bean)
	bean.JiShen = lGuiRen
}

func (r *DaYunInfoGet) setXiongShen(z SiZhuModel, bean *DaYunInfo) {
	l := r.getXiongShen(z, bean)
	bean.XiongShen = l
}

func (r *DaYunInfoGet) getGuiRen(z SiZhuModel, yun *DaYunInfo) []*GuiRenItem {
	ret := make([]*GuiRenItem, 0)

	obj := new(GuiRenGet)

	//-------- -----------------------------

	//-------- -----------------------------
	l := obj.TianYi(z.RiGan, z.NianZhi, z.YueZhi, z.RiZhi, z.ShiZhi, yun.Zhi)
	if len(l) > 0 {
		ret = append(ret, l...)
	}
	l = obj.TianYi(z.NianGan, z.NianZhi, z.YueZhi, z.RiZhi, z.ShiZhi, yun.Zhi)
	if len(l) > 0 {
		ret = append(ret, l...)
	}

	//-------- -----------------------------
	l = obj.ShiGanLu(z.RiGan, z.NianZhi, z.YueZhi, z.RiZhi, z.ShiZhi, yun.Zhi)
	if len(l) > 0 {
		ret = append(ret, l...)
	}
	//-------- -----------------------------
	l = obj.WenChang(z.NianGan, z.NianZhi, z.YueZhi, z.RiZhi, z.ShiZhi, yun.Zhi)
	if len(l) > 0 {
		ret = append(ret, l...)
	}
	l = obj.WenChang(z.RiGan, z.NianZhi, z.YueZhi, z.RiZhi, z.ShiZhi, yun.Zhi)
	if len(l) > 0 {
		ret = append(ret, l...)
	}
	//-------- -----------------------------
	l = obj.HuaGai(z.NianZhi, z.YueZhi, z.RiZhi, z.ShiZhi, yun.Zhi)
	if len(l) > 0 {
		ret = append(ret, l...)
	}
	l = obj.HuaGai(z.RiZhi, z.NianZhi, z.YueZhi, z.ShiZhi, yun.Zhi)
	if len(l) > 0 {
		ret = append(ret, l...)
	}
	//-------- -----------------------------
	l = obj.HuaGai(z.RiZhi, z.NianZhi, z.YueZhi, z.ShiZhi, yun.Zhi)
	if len(l) > 0 {
		ret = append(ret, l...)
	}
	//-------- -----------------------------
	l = obj.JiangXing(z.NianZhi, z.YueZhi, z.RiZhi, z.ShiZhi, yun.Zhi)
	if len(l) > 0 {
		ret = append(ret, l...)
	}
	l = obj.JiangXing(z.RiZhi, z.NianZhi, z.YueZhi, z.ShiZhi, yun.Zhi)
	if len(l) > 0 {
		ret = append(ret, l...)
	}
	//-------- -----------------------------
	//-------- -----------------------------
	l = obj.YiMa(z.NianZhi, z.YueZhi, z.RiZhi, z.ShiZhi, yun.Zhi)
	if len(l) > 0 {
		ret = append(ret, l...)
	}
	l = obj.YiMa(z.RiZhi, z.NianZhi, z.YueZhi, z.ShiZhi, yun.Zhi)
	if len(l) > 0 {
		ret = append(ret, l...)
	}
	//-------- -----------------------------
	l = obj.JinYu(z.RiGan, z.NianZhi, z.YueZhi, z.RiZhi, z.ShiZhi, yun.Zhi)
	if len(l) > 0 {
		ret = append(ret, l...)
	}
	//-------- -----------------------------
	l = obj.TianYiGuiRen(z.YueZhi, z.NianZhi, z.YueZhi, z.RiZhi, z.ShiZhi, yun.Zhi)
	if len(l) > 0 {
		ret = append(ret, l...)
	}
	//-------- -----------------------------
	return ret
}

func (r *DaYunInfoGet) getXiongShen(z SiZhuModel, bean *DaYunInfo) []*XiongShenItem {
	obj := new(XiongShenGet)
	//-------- -----------------------------
	ret := make([]*XiongShenItem, 0)
	if l := obj.YanRen(z.RiGan, z.NianZhi, z.YueZhi, z.RiZhi, z.ShiZhi, bean.Zhi); len(l) > 0 {
		ret = append(ret, l...)
	}
	//-------- -----------------------------
	if l := obj.LiuJiaKongWang(z.RiGan+z.RiZhi, z.NianZhi, z.YueZhi, z.ShiZhi, bean.Zhi); len(l) > 0 {
		ret = append(ret, l...)
	}
	if l := obj.LiuJiaKongWang(z.NianGan+z.NianZhi, z.YueZhi, z.RiZhi, z.ShiZhi, bean.Zhi); len(l) > 0 {
		ret = append(ret, l...)
	}
	//-------- -----------------------------
	if l := obj.TaoHua(z.RiZhi, z.NianZhi, z.YueZhi, z.ShiZhi, bean.Zhi); len(l) > 0 {
		ret = append(ret, l...)
	}
	if l := obj.TaoHua(z.NianZhi, z.YueZhi, z.RiZhi, z.ShiZhi, bean.Zhi); len(l) > 0 {
		ret = append(ret, l...)
	}

	//-------- -----------------------------
	if l := obj.TianLuoDiWang(z.RiZhi, z.NianZhi, z.YueZhi, z.ShiZhi, bean.Zhi); len(l) > 0 {
		ret = append(ret, l...)
	}
	if l := obj.TianLuoDiWang(z.NianZhi, z.YueZhi, z.RiZhi, z.ShiZhi, bean.Zhi); len(l) > 0 {
		ret = append(ret, l...)
	}
	//-------- -----------------------------

	return ret
}
