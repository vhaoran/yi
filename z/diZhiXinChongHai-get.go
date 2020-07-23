package z

import (
	"fmt"

	"github.com/vhaoran/vchat/common/g"

	"github.com/vhaoran/yi/z/model"
)

//----------------------------------------------------
// auth: whr  date:2020/7/2216:26--------------------------
// ####请勿擅改此功能代码####
// 用途： 地支弄冲克害
//---------------------------------------------

type (
	DiZiXingChongHaiGet struct {
	}
)

var DiZiXingChongHaiGetX = new(DiZiXingChongHaiGet)

func (r *DiZiXingChongHaiGet) Get(z *model.SiZhuModel) []string {
	ret := make([]string, 0)
	if l := r.GetXiangHai(z.NianZhi, z.YueZhi, z.RiZhi, z.ShiZhi); len(l) > 0 {
		ret = append(ret, l...)
	}
	if l := r.GetXiangXing(z.NianZhi, z.YueZhi, z.RiZhi, z.ShiZhi); len(l) > 0 {
		ret = append(ret, l...)
	}
	if l := r.GetLiuChong(z.NianZhi, z.YueZhi, z.RiZhi, z.ShiZhi); len(l) > 0 {
		ret = append(ret, l...)
	}
	return ret
}

//相害
func (r *DiZiXingChongHaiGet) GetXiangHai(lZhi ...string) []string {
	l := []string{
		//子未相害、丑午相害、寅巳相害、
		//卯辰相害、申亥相害、酉戌相害
		"子未", "丑午", "寅巳", "卯辰", "申亥", "酉戌",
	}

	return r.match("害", l, lZhi...)
}

//相刑
func (r *DiZiXingChongHaiGet) GetXiangXing(lZhi ...string) []string {
	l := []string{
		//丑刑未、未刑戌、戌刑丑、为无恩之刑；
		//寅刑巳、巳刑申、申刑寅，为持势之刑；
		//子刑卯、卯刑子为无礼之刑；
		"丑未", "未戌", "戌丑", "寅巳", "巳申", "申寅", "子卯",
		// 辰午酉亥 自刑
		"辰辰", "午午", "酉酉", "酉酉", "亥亥",
	}

	return r.match("刑",l, lZhi...)
}

//相冲
func (r *DiZiXingChongHaiGet) GetLiuChong(lZhi ...string) []string {
	l := []string{
		//相冲，相冲，相冲，巳亥相冲
		"子午", "丑未", "寅申", "卯酉", "辰戌", "巳亥",
	}

	return r.match("冲", l, lZhi...)
}

func (r *DiZiXingChongHaiGet) match(title string, patMap []string, lZhi ...string) []string {
	ret := make([]string, 0)
	for k1, v1 := range lZhi {
		for k2, v2 := range lZhi {
			if k1 == k2 {
				continue
			}
			if g.InSlice(v1+v2, lZhi) {
				ret = append(ret, fmt.Sprint(v1+v2, "(", title, ")"))
			}
		}
	}

	return ret
}
