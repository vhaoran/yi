package z

import (
	"strings"

	"github.com/vhaoran/yi/cmn"
	"github.com/vhaoran/yi/z/model"
)

//----------------------------------------------------
// auth: whr  date:2020/7/2216:31--------------------------
// ####请勿擅改此功能代码####
// 用途：吉神分析
//---------------------------------------------

type (
	GuiRenGet struct {
	}

	GuiRenItem struct {
		model.KVRoot
	}
)

func (r *GuiRenGet) Call(z *model.SiZhuModel) []*GuiRenItem {
	ret := make([]*GuiRenItem, 0)

	//-------- -----------------------------
	l := r.TianDe(z.YueZhi, z.NianZhi, z.RiZhi, z.ShiZhi)
	if len(l) > 0 {
		ret = append(ret, l...)
	}

	//-------- -----------------------------
	l = r.YueDe(z.YueZhi, z.NianGan, z.YueGan, z.RiGan, z.ShiGan)
	if len(l) > 0 {
		ret = append(ret, l...)
	}

	//-------- -----------------------------
	l = r.TianYi(z.RiGan, z.NianZhi, z.YueZhi, z.RiZhi, z.ShiZhi)
	if len(l) > 0 {
		ret = append(ret, l...)
	}
	l = r.TianYi(z.NianGan, z.NianZhi, z.YueZhi, z.RiZhi, z.ShiZhi)
	if len(l) > 0 {
		ret = append(ret, l...)
	}

	//-------- -----------------------------
	l = r.ShiGanLu(z.RiGan, z.NianZhi, z.YueZhi, z.RiZhi, z.ShiZhi)
	if len(l) > 0 {
		ret = append(ret, l...)
	}
	//-------- -----------------------------
	l = r.WenChang(z.NianGan, z.NianZhi, z.YueZhi, z.RiZhi, z.ShiZhi)
	if len(l) > 0 {
		ret = append(ret, l...)
	}
	l = r.WenChang(z.RiGan, z.NianZhi, z.YueZhi, z.RiZhi, z.ShiZhi)
	if len(l) > 0 {
		ret = append(ret, l...)
	}
	//-------- -----------------------------
	l = r.HuaGai(z.NianZhi, z.YueZhi, z.RiZhi, z.ShiZhi)
	if len(l) > 0 {
		ret = append(ret, l...)
	}
	l = r.HuaGai(z.RiZhi, z.NianZhi, z.YueZhi, z.ShiZhi)
	if len(l) > 0 {
		ret = append(ret, l...)
	}
	//-------- -----------------------------
	l = r.HuaGai(z.RiZhi, z.NianZhi, z.YueZhi, z.ShiZhi)
	if len(l) > 0 {
		ret = append(ret, l...)
	}
	//-------- -----------------------------
	l = r.JiangXing(z.NianZhi, z.YueZhi, z.RiZhi, z.ShiZhi)
	if len(l) > 0 {
		ret = append(ret, l...)
	}
	l = r.JiangXing(z.RiZhi, z.NianZhi, z.YueZhi, z.ShiZhi)
	if len(l) > 0 {
		ret = append(ret, l...)
	}
	//-------- -----------------------------
	//-------- -----------------------------
	l = r.YiMa(z.NianZhi, z.YueZhi, z.RiZhi, z.ShiZhi)
	if len(l) > 0 {
		ret = append(ret, l...)
	}
	l = r.YiMa(z.RiZhi, z.NianZhi, z.YueZhi, z.ShiZhi)
	if len(l) > 0 {
		ret = append(ret, l...)
	}
	//-------- -----------------------------
	l = r.JinYu(z.RiGan, z.NianZhi, z.YueZhi, z.RiZhi, z.ShiZhi)
	if len(l) > 0 {
		ret = append(ret, l...)
	}
	//-------- -----------------------------
	l = r.TianYiGuiRen(z.YueZhi, z.NianZhi, z.YueZhi, z.RiZhi, z.ShiZhi)
	if len(l) > 0 {
		ret = append(ret, l...)
	}
	//-------- -----------------------------
	return ret
}

//
//天德贵人,月支上与其它二支的组合
//l为其它柱的天十和地址，包括月上的地支
func (r *GuiRenGet) TianDe(yueZhi string, lZhi ...string) []*GuiRenItem {
	//子丑寅卯辰巳午未申酉戌亥
	//甲乙丙丁戊己庚辛壬癸
	m := cmn.KV{
		"子": "巳",
		"丑": "庚",
		"寅": "丁",
		"卯": "申",
		"辰": "壬",
		"巳": "辛",
		"午": "亥",
		"未": "甲",
		"申": "癸",
		"酉": "寅",
		"戌": "丙",
		"亥": "乙",
	}

	return r.match(m, "天德贵人", yueZhi, lZhi...)
}

//月德貴人,查其它支的干，包括月干
func (r *GuiRenGet) YueDe(yueZhi string, lGan ...string) []*GuiRenItem {
	//子丑寅卯辰巳午未申酉戌亥
	//甲乙丙丁戊己庚辛壬癸
	m := cmn.KV{
		"寅": "丙",
		"午": "丙",
		"戌": "丙",

		"申": "壬",
		"子": "壬",
		"辰": "壬",

		"亥": "甲",
		"卯": "甲",
		"未": "甲",

		"巳": "庚",
		"酉": "庚",
		"丑": "庚",
	}

	return r.match(m, "月德贵人", yueZhi, lGan...)
}

//天乙貴人
//lZhi,其它支，其4个支
func (r *GuiRenGet) TianYi(nianOrRiGan string, lZhi ...string) []*GuiRenItem {
	//甲、戊年日干见支中丑、未；
	//乙、己年日干见支中子、申；
	//丙、丁年日干见支中亥、酉；
	//庚、辛年日干见支中寅、午；
	//壬、癸年日干见支中卯、巳
	m := cmn.KV{
		"甲": "丑未",
		"戊": "丑未",
		"乙": "子申",
		"己": "子申",
		"丙": "亥酉",
		"丁": "亥酉",
		"壬": "卯巳",
		"癸": "卯巳",
	}
	return r.match(m, "天乙貴人", nianOrRiGan, lZhi...)
}

//十干禄
func (r *GuiRenGet) ShiGanLu(riGan string, lZhi ...string) []*GuiRenItem {
	//子丑寅卯辰巳午未申酉戌亥
	//甲乙丙丁戊己庚辛壬癸
	m := cmn.KV{
		//丑辰未戌
		"甲": "寅",
		"乙": "卯",

		"丙": "巳",
		"丁": "午",

		"戊": "巳",
		"己": "午",

		"庚": "申",
		"辛": "酉",
		"壬": "亥",
		"癸": "子",
	}

	return r.match(m, "十干禄", riGan, lZhi...)
}

//文昌
func (r *GuiRenGet) WenChang(nianOrRiGan string, lZhi ...string) []*GuiRenItem {
	//甲乙 见巳或午
	//丙戊 见 申\
	//丁已 见 酉
	//庚见 亥
	//辛见 子
	//壬见 寅
	//癸见 卯
	m := cmn.KV{
		"甲": "巳午",
		"乙": "巳午",
		"丙": "申",
		"戊": "申",
		"丁": "酉",
		"已": "酉",
		"庚": "亥",
		"辛": "子",
		"壬": "寅",
		"癸": "卯",
	}

	return r.match(m, "文昌星", nianOrRiGan, lZhi...)
}

//华盖,
func (r *GuiRenGet) HuaGai(nianOrRiZhi string, otherZhi ...string) []*GuiRenItem {
	//子丑寅卯辰巳午未申酉戌亥
	m := cmn.KV{
		"寅": "戌",
		"午": "戌",
		"戌": "戌",

		"申": "辰",
		"子": "辰",
		"辰": "辰",

		"巳": "丑",
		"酉": "子",
		"丑": "寅",

		"亥": "未",
		"卯": "未",
		"未": "未",
	}

	return r.match(m, "华盖星", nianOrRiZhi, otherZhi...)
}

//将星,
func (r *GuiRenGet) JiangXing(nianOrRiZhi string, otherZhi ...string) []*GuiRenItem {
	//子丑寅卯辰巳午未申酉戌亥
	m := cmn.KV{
		"寅": "午",
		"午": "午",
		"戌": "午",

		"申": "子",
		"子": "子",
		"辰": "子",

		"巳": "酉",
		"酉": "酉",
		"丑": "酉",

		"亥": "卯",
		"卯": "卯",
		"未": "卯",
	}

	return r.match(m, "将星", nianOrRiZhi, otherZhi...)
}

//驿马,
func (r *GuiRenGet) YiMa(nianOrRiZhi string, otherZhi ...string) []*GuiRenItem {
	//子丑寅卯辰巳午未申酉戌亥
	m := cmn.KV{
		"寅": "申",
		"午": "申",
		"戌": "申",

		"申": "寅",
		"子": "寅",
		"辰": "寅",

		"巳": "亥",
		"酉": "亥",
		"丑": "亥",

		"亥": "巳",
		"卯": "巳",
		"未": "巳",
	}

	return r.match(m, "驿马星", nianOrRiZhi, otherZhi...)
}

//金舆
func (r *GuiRenGet) JinYu(riGan string, l4Zhi ...string) []*GuiRenItem {
	//子丑寅卯辰巳午未申酉戌亥
	//甲乙丙丁戊己庚辛壬癸
	m := cmn.KV{
		"甲": "辰",
		"乙": "巳",
		"丙": "未",
		"丁": "申",
		"戊": "未",
		"己": "申",
		"庚": "戌", //子卯午酉
		"辛": "亥",
		"壬": "丑",
		"癸": "寅",
	}

	return r.match(m, "金舆贵人", riGan, l4Zhi...)
}

//天医路人
func (r *GuiRenGet) TianYiGuiRen(yueZhi string, otherZhi ...string) []*GuiRenItem {
	//子丑寅卯辰巳午未申酉戌亥
	//甲乙丙丁戊己庚辛壬癸
	m := cmn.KV{
		"子": "亥",
		"丑": "子",
		"寅": "丑",
		"卯": "寅",
		"辰": "卯",
		"巳": "辰",
		"午": "巳", //
		"未": "午",
		"申": "未",
		"酉": "申",
		"戌": "酉",
		"亥": "戌",
	}

	return r.match(m, "天医贵人", yueZhi, otherZhi...)
}

func (r *GuiRenGet) match(m cmn.KV, guirenName, key string, l ...string) []*GuiRenItem {
	s, ok := m[key]
	if !ok {
		return nil
	}

	ret := make([]*GuiRenItem, 0)
	for _, sub := range l {
		if strings.Contains(s, sub) {
			ret = append(ret, &GuiRenItem{
				KVRoot: model.KVRoot{
					Name:    guirenName,
					Comment: key + s,
				},
			})
		}
	}

	//
	return ret
}
