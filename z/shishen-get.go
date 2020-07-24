package z

import (
	"github.com/vhaoran/yi/cmn"
	. "github.com/vhaoran/yi/z/model"
)

//----------------------------------------------------
// auth: whr  date:2020/7/2216:28--------------------------
// ####请勿擅改此功能代码####
// 用途：天元论命，十神分析
//   偏官又名七杀
//   偏印又名枭
//---------------------------------------------
type (
	//十食获取
	ShiShenGet struct {
	}
)

var ShiShenGetX = new(ShiShenGet)

//传入四柱
func (r *ShiShenGet) Get(z *SiZhuModel) *ShiShenModel {
	if z == nil {
		return nil
	}
	riGan := z.RiGan
	bean := &ShiShenModel{
		NianGan: cmn.GetShiShen(riGan, z.NianGan),
		YueGan:  cmn.GetShiShen(riGan, z.YueGan),
		RiGan:   "日",
		ShiGan:  cmn.GetShiShen(riGan, z.ShiGan),
	}
	return bean
}
