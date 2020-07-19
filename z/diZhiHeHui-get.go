package z

import (
	"github.com/vhaoran/yi/cmn"
	. "github.com/vhaoran/yi/z/model"
)

type (
	//地支合会
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

var shengDiBanHe_KV = cmn.KV{
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
func (r *DiZhiHeHuiGet) GetSanHe(lstDiZhi ...string) []*HeModel {
	return r.match(sanHe_KV, lstDiZhi...)
}

//得到地支三会
func (r *DiZhiHeHuiGet) GetSanHui(lstDiZhi ...string) []*HeModel {
	return r.match(sanHui_KV, lstDiZhi...)
}

func (r *DiZhiHeHuiGet) GetLiuHe(lstDiZhi ...string) []*HeModel {
	return r.match(liuHe_KV, lstDiZhi...)
}

func (r *DiZhiHeHuiGet) GetShenDiBanHe(lstDiZhi ...string) []*HeModel {
	return r.match(shengDiBanHe_KV, lstDiZhi...)
}

func (r *DiZhiHeHuiGet) GetMuDiBanHe(lstDiZhi ...string) []*HeModel {
	return r.match(muDiBanHe_KV, lstDiZhi...)
}

//二支合的情况
func (r *DiZhiHeHuiGet) match(patKV cmn.KV, lstDiZhi ...string) []*HeModel {
	ret := make([]*HeModel, 0)

	for key, wuxing := range patKV {
		ok := cmn.StrEachInSlice(key, lstDiZhi...)
		if ok {
			ret = append(ret, &HeModel{
				He:          key,
				HeHuaWuXing: wuxing,
			})
		}
	}

	return ret
}
