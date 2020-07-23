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
		NianGan: r.GetSingle(riGan, z.NianGan),
		YueGan:  r.GetSingle(riGan, z.YueGan),
		RiGan:   "日",
		ShiGan:  r.GetSingle(riGan, z.ShiGan),
	}
	return bean
}

//ganOrZhi干或支
//other 其它幹或支
func (r *ShiShenGet) GetSingle(ganOrZhi, other string) string {
	rx, same := cmn.GetRX(ganOrZhi, other)
	if rx == cmn.RX_WoSheng {
		if same {
			return "食神"
		}
		return "伤官"
	}
	if rx == cmn.RX_ShengWo {
		if same {
			return "偏印"
		}
		return "正印"
	}

	//-------- -----------------------------
	if rx == cmn.RX_KeWo {
		if same {
			return "偏官"
		}
		return "正官"
	}
	if rx == cmn.RX_WoKe {
		if same {
			return "偏财"
		}
		return "正财"
	}
	//
	if rx == cmn.RX_TongWo {
		if same {
			return "比肩"
		}
		return "劫财"
	}
	return ""
}
