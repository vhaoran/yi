package z

import (
	"fmt"
	"strings"

	"github.com/vhaoran/yi/cmn"
	"github.com/vhaoran/yi/z/model"
)

//----------------------------------------------------
// auth: whr  date:2020/7/2314:54--------------------------
// ####请勿擅改此功能代码####
// 用途：地支藏幹的計算
//在八字学中,每个地支中都含有几个天干,
//它们被称为“藏干”在分析八字中扮演着重要
//的角色。在八字的术语中被称为“支藏人元”。
//
//要达到八字最大的效果,就必须知道每个地支
//中所包含的天干。一个地支的本性是由其内所
//藏的主要元素或主要的气来确定。在八字的古
//籍中有写道天干的气,是属于纯洁和清澈的；
//而地支的气则属于复杂和浓密的。地支的气之
//所以复杂是因为其内涵有隐藏的天干。
//
//当我们把天干、地支和隐藏的天干三样东
//西合在一起后,我们就会得到宇宙三位一体
//的效果:就是每柱中的天、地、人
//---------------------------------------------
type (
	DiZhiCanGanGet struct {
	}
	DiZhiCanGanData struct {
		Name    string
		GanList []string
	}
)

func (r *DiZhiCanGanData) ToString() string {
	str := strings.Join(r.GanList, ",")
	s := fmt.Sprint(r.Name, "(", str, ")")
	return s
}

var DiZhiCanGanGetX = new(DiZhiCanGanGet)

//获取四柱中十神藏干
func (r *DiZhiCanGanGet) Get(z *model.SiZhuModel) []*DiZhiCanGanData {
	ret := make([]*DiZhiCanGanData, 0)
	//--------nianZhi -----------------------------
	for _, zhu := range z.ZhiList() {
		if l := r.GetSingle(zhu); len(l) > 0 {
			ret = append(ret, &DiZhiCanGanData{
				Name:    zhu,
				GanList: l,
			})
		}
	}
	return ret
}

func (r *DiZhiCanGanGet) GetSingle(zhi string) (lstTanGan []string) {
	//子丑寅卯辰巳午未申酉戌亥
	//甲乙丙丁戊己庚辛壬癸
	m := cmn.KV{
		"子": "癸",
		"丑": "己癸辛",
		"寅": "甲丙戊", //
		"卯": "乙",
		"辰": "戊乙癸",
		"巳": "丙戊庚",
		"午": "丁己",
		"未": "己丁乙",
		"申": "庚壬戊",
		"酉": "辛",
		"戌": "戊辛丁",
		"亥": "壬甲",
	}

	s, ok := m[zhi]
	if !ok {
		return nil
	}
	//
	l := strings.Split(s, "")
	return l
}
