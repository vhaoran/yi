package z

import (
	"github.com/vhaoran/yi/cmn"
	. "github.com/vhaoran/yi/z/model"
)

type (
	DiZhiHeHuiGet struct {
	}
)

var DiZhiHeHuiGetX = new(DiZhiHeHuiGet)

var liuHe_KV = cmn.KV{
	/*
	   子丑合土
	   寅亥合木
	   卯戌合火
	   辰酉合金
	   巳申合水
	   午未合土
	*/

	"子丑": "土",
	"寅亥": "木",
	"卯戌": "火",
	"辰酉": "金",
	"巳申": "水",
	"午未": "土",
}

var sanHe_KV = cmn.KV{
	"申子辰": "水",
	"亥卯未": "木",
	"寅午戌": "火",
	"巳酉丑": "金",
}

//地支三会
var sanHui_KV = cmn.KV{
	/*寅卯辰三会东方木，
	巳午未三会南方火，
	申酉戌三会西方金，
	亥子丑三会北方水 */

	"亥子丑": "水",
	"寅卯辰": "木",
	"巳午未": "火",
	"申酉戌": "金",
}

var banHe_KV = cmn.KV{
	"申子": "水",
	"亥卯": "木",
	"寅午": "火",
	"巳酉": "金",
}

var muDiBanHe_KV = cmn.KV{
	"子辰": "水",
	"卯未": "木",
	"午戌": "火",
	"酉丑": "金",
}

//取地支三合
func (r *DiZhiHeHuiGet) GetSanHe(z *SiZhuModel) []*HeModel {
	ret := make([]*HeModel, 0)

	l := z.ZhiList()
	for k1, v1 := range l {
		for k2, v2 := range l {
			if k1 == k2 {
				continue
			}
			for k3, v3 := range l {
				if k2 == k3 {
					continue
				}

				if heWuXing := r.he(sanHe_KV, v1, v2, v3); len(heWuXing) > 0 {
					ret = append(ret, &HeModel{
						He:          v1 + v2,
						HeHuaWuXing: heWuXing,
					})
				}
			}
		}
	}

	return ret
}

//得到地支三会
func (r *DiZhiHeHuiGet) GetSanHui(z *SiZhuModel) []*HeModel {
	ret := make([]*HeModel, 0)

	l := z.ZhiList()
	for k1, v1 := range l {
		for k2, v2 := range l {
			if k1 == k2 {
				continue
			}
			for k3, v3 := range l {
				if k2 == k3 {
					continue
				}

				if heWuXing := r.he(sanHui_KV, v1, v2, v3); len(heWuXing) > 0 {
					ret = append(ret, &HeModel{
						He:          v1 + v2,
						HeHuaWuXing: heWuXing,
					})
				}
			}
		}
	}

	return ret
}

func (r *DiZhiHeHuiGet) GetLiuHe(z *SiZhuModel) []*HeModel {
	ret := make([]*HeModel, 0)

	l := z.ZhiList()
	for k, src := range l {
		for k1, dst := range l {
			if k == k1 {
				continue
			}
			if heWuXing := r.he(liuHe_KV, src, dst); len(heWuXing) > 0 {
				ret = append(ret, &HeModel{
					He:          src + dst,
					HeHuaWuXing: heWuXing,
				})
			}
		}
	}

	return ret
}

func (r *DiZhiHeHuiGet) GetShenDiBanHe(z *SiZhuModel) []*HeModel {
	ret := make([]*HeModel, 0)

	l := z.ZhiList()
	for k, src := range l {
		for k1, dst := range l {
			if k == k1 {
				continue
			}
			if heWuXing := r.he(banHe_KV, src, dst); len(heWuXing) > 0 {
				ret = append(ret, &HeModel{
					He:          src + dst,
					HeHuaWuXing: heWuXing,
				})
			}
		}
	}

	return ret
}

func (r *DiZhiHeHuiGet) GetMuDiBanHe(z *SiZhuModel) []*HeModel {
	ret := make([]*HeModel, 0)

	l := z.ZhiList()
	for k, src := range l {
		for k1, dst := range l {
			if k == k1 {
				continue
			}
			if heWuXing := r.he(muDiBanHe_KV, src, dst); len(heWuXing) > 0 {
				ret = append(ret, &HeModel{
					He:          src + dst,
					HeHuaWuXing: heWuXing,
				})
			}
		}
	}

	return ret
}

//l 最多有三个
func (r *DiZhiHeHuiGet) he(m cmn.KV, l ...string) (heWuXing string) {
	heWuXing = ""

	//--------dst -----------------------------
	dst := ""
	for _, str := range l {
		dst += str
	}

	//--------match -----------------------------
	s, ok := liuHe_KV[dst]
	if ok {
		return s
	}
	return ""
}
