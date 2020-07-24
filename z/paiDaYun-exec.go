package z

import (
	"fmt"

	"github.com/vhaoran/yi/cmn"
	. "github.com/vhaoran/yi/z/model"
)

type (
	PaiDaYunExec struct {
	}

	//主要用于显示
	DaYuItem struct {
		//大支干支
		Gan string
		Zhi string

		//大运名称
		Gong12 string
		//大运简称
		Gong12Short string

		//该运开始时年龄
		YearsOld int

		//开始年
		NianStart int
		//结束年
		NianEnd int
	}
)

func (r *DaYuItem) Zhu() string {
	return r.Gan + r.Zhi
}

func (r *DaYuItem) ToString() string {
	s := r.Zhu() + "(" + r.Gong12 + ")"
	s1 := fmt.Sprint("(", r.NianStart, "-", r.NianEnd, ")")
	return s + s1
}

func (r *PaiDaYunExec) Exec(z *SiZhuModel, qiYunSuiShu int) []*DaYuItem {
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

func (r *PaiDaYunExec) calc(yueGan, yueZhi string, qiYunYearsOld, bornNian int, isAsc bool) []*DaYuItem {
	ret := make([]*DaYuItem, 0)
	//zhiStr := "子丑寅卯辰巳午未申酉戌亥"
	//ganStr := "甲乙丙丁戊己庚辛壬癸"
	ganIdx := cmn.GanIndex(yueGan)
	zhiIdx := cmn.ZhiIndex(yueZhi)

	//排100年运
	offset := qiYunYearsOld
	for c := 0; c < 11; c++ {
		i, j := ganIdx+1+c, zhiIdx+c+1
		gan, zhi := cmn.Gan_list[i%10], cmn.Zhi_list[j%12]
		bean := &DaYuItem{
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

func (r *PaiDaYunExec) setShiShen(riGan string, v *DaYuItem) {
	shiShen := cmn.GetShiShen(riGan, v.Gan)
	//
	v.Gong12 = shiShen
	v.Gong12Short = cmn.GetShiShenShort(v.Gong12)
}

func (r *PaiDaYunExec) setGong12(riGan string, v *DaYuItem) {
	_, gong12 := cmn.GetGong12OfGan(riGan, v.Zhi)
	//
	v.Gong12 = cmn.Gong12FullStr(gong12)
	v.Gong12Short = gong12
}
