package z

import (
	"strings"

	"github.com/vhaoran/vchat/common/g"

	"github.com/vhaoran/yi/cmn"
	. "github.com/vhaoran/yi/z/model"
)

//----------------------------------------------------
// auth: whr  date:2020/7/2216:29--------------------------
// ####请勿擅改此功能代码####
// 用途：凶神分析
//---------------------------------------------

type (
	XiongShenGet struct {
	}
	XiongShenData struct {
		KVRoot
	}
)

func (r *XiongShenGet) Call(z *SiZhuModel) []*XiongShenData {
	//-------- -----------------------------
	ret := make([]*XiongShenData, 0)
	if l := r.YanRen(z.RiGan, z.NianZhi, z.YueZhi, z.RiZhi, z.ShiZhi); len(l) > 0 {
		ret = append(ret, l...)
	}
	//-------- -----------------------------
	if l := r.LieJiaKongWang(z.RiGan+z.RiZhi, z.NianZhi, z.YueZhi, z.ShiZhi); len(l) > 0 {
		ret = append(ret, l...)
	}
	if l := r.LieJiaKongWang(z.NianGan+z.NianZhi, z.YueZhi, z.RiZhi, z.ShiZhi); len(l) > 0 {
		ret = append(ret, l...)
	}
	//-------- -----------------------------
	if l := r.TaoHua(z.RiZhi, z.NianZhi, z.YueZhi, z.ShiZhi); len(l) > 0 {
		ret = append(ret, l...)
	}
	if l := r.TaoHua(z.NianZhi, z.YueZhi, z.RiZhi, z.ShiZhi); len(l) > 0 {
		ret = append(ret, l...)
	}
	//-------- -----------------------------
	if l := r.YinYangChaCuo(z.RiGan + z.RiZhi); len(l) > 0 {
		ret = append(ret, l...)
	}
	if l := r.YinYangChaCuo(z.NianGan + z.NianZhi); len(l) > 0 {
		ret = append(ret, l...)
	}
	//-------- -----------------------------
	if l := r.TianLuoDiWang(z.RiZhi, z.NianZhi, z.YueZhi, z.ShiZhi); len(l) > 0 {
		ret = append(ret, l...)
	}
	if l := r.TianLuoDiWang(z.NianZhi, z.YueZhi, z.RiZhi, z.ShiZhi); len(l) > 0 {
		ret = append(ret, l...)
	}
	//-------- -----------------------------

	return nil
}

//羊刃
func (r *XiongShenGet) YanRen(riGan string, lZhi ...string) []*XiongShenData {
	//子丑寅卯辰巳午未申酉戌亥
	//甲乙丙丁戊己庚辛壬癸
	m := cmn.KV{
		"甲": "卯",
		"乙": "寅",
		"丙": "午",
		"丁": "巳",
		"戊": "午",
		"己": "巳",
		"庚": "酉",
		"辛": "申",
		"壬": "子",
		"癸": "亥",
	}
	return r.match(m, "羊刃煞", riGan, lZhi...)
}

//六甲空亡
func (r *XiongShenGet) LieJiaKongWang(nianOrRiGanZhi string, otherZhi ...string) []*XiongShenData {
	//子丑寅卯辰巳午未申酉戌亥
	//甲乙丙丁戊己庚辛壬癸
	m := cmn.KV{
		"甲子": "戌亥",
		"甲戌": "申酉",
		"甲申": "午未",
		"甲午": "辰巳",
		"甲辰": "寅卯",
		"甲寅": "子丑",
	}
	return r.match(m, "羊刃煞", nianOrRiGanZhi, otherZhi...)
}

//桃花
func (r *XiongShenGet) TaoHua(nianOrRiZhi string, otherZhi ...string) []*XiongShenData {
	//子丑寅卯辰巳午未申酉戌亥
	//甲乙丙丁戊己庚辛壬癸
	m := cmn.KV{
		"寅": "卯",
		"午": "卯",
		"戌": "卯",

		"申": "酉",
		"子": "酉",
		"辰": "酉",

		"亥": "子",
		"卯": "子",
		"未": "子",

		"巳": "午",
		"酉": "午",
		"丑": "午",
	}
	return r.match(m, "桃花煞", nianOrRiZhi, otherZhi...)
}

//阴阳差错
func (r *XiongShenGet) YinYangChaCuo(nianOrRiGanZhi string) []*XiongShenData {
	//子丑寅卯辰巳午未申酉戌亥
	//甲乙丙丁戊己庚辛壬癸
	l := []string{"丙子", "丁丑", "戊寅", "辛卯",
		"壬辰", "癸巳", "丙午", "丁未",
		"戊申", "辛酉", "壬戌", "癸亥"}

	ret := make([]*XiongShenData, 0)
	if g.InSlice(nianOrRiGanZhi, l) {
		ret = append(ret, &XiongShenData{
			KVRoot: KVRoot{
				Name:    "阴阳差错",
				Comment: nianOrRiGanZhi,
			},
		})
	}
	return ret
}

//天罗地网
func (r *XiongShenGet) TianLuoDiWang(nianOrRiZhi string, otherZhi ...string) []*XiongShenData {
	//子丑寅卯辰巳午未申酉戌亥
	//甲乙丙丁戊己庚辛壬癸
	m := cmn.KV{
		"戌": "亥",
		"亥": "戌",
		"辰": "巳",
		"巳": "辰",
	}
	return r.match(m, "天罗地网", nianOrRiZhi, otherZhi...)
}

//动煞
func (r *XiongShenGet) JieSha(nianOrRiZhi string, otherZhi ...string) []*XiongShenData {
	//子丑寅卯辰巳午未申酉戌亥
	//甲乙丙丁戊己庚辛壬癸
	m := cmn.KV{
		"寅": "亥",
		"午": "亥",
		"戌": "亥",

		"申": "巳",
		"子": "巳",
		"辰": "巳",

		"亥": "申",
		"卯": "申",
		"未": "申",

		"巳": "寅",
		"酉": "寅",
		"丑": "寅",
	}
	return r.match(m, "动煞", nianOrRiZhi, otherZhi...)
}

func (r *XiongShenGet) match(m cmn.KV, title, key string, l ...string) []*XiongShenData {
	s, ok := m[key]
	if !ok {
		return nil
	}

	ret := make([]*XiongShenData, 0)
	for _, sub := range l {
		if strings.Contains(s, sub) {
			ret = append(ret, &XiongShenData{
				KVRoot: KVRoot{
					Name:    title,
					Comment: key + s,
				},
			})
		}
	}

	//
	return ret
}
