package z

import (
	"github.com/vhaoran/vchat/common/g"
)

//----------------------------------------------------
// auth: whr  date:2020/7/2216:26--------------------------
// ####请勿擅改此功能代码####
// 用途： 地支弄效忠上克害
//---------------------------------------------

type (
	DiZiXingChongHaiGet struct {
	}
)

var DiZiXingChongHaiGetX = new(DiZiXingChongHaiGet)

var liuChong_list = []string{
	//相冲，相冲，相冲，巳亥相冲
	"子午",
	"丑未",
	"寅申",
	"卯酉",
	"辰戌",
	"巳亥",
}
var xiangXing_list = []string{
	//丑刑未、未刑戌、戌刑丑、为无恩之刑；
	//寅刑巳、巳刑申、申刑寅，为持势之刑；
	//子刑卯、卯刑子为无礼之刑；
	"丑未",
	"未戌",
	"戌丑",
	"寅巳",
	"巳申",
	"申寅",
	"子卯",
	// 辰午酉亥 自刑
	"辰辰",
	"午午",
	"酉酉",
	"酉酉",
	"亥亥",
}
var xiangHai_list = []string{
	//子未相害、丑午相害、寅巳相害、
	//卯辰相害、申亥相害、酉戌相害
	"子未",
	"丑午",
	"寅巳",
	"卯辰",
	"申亥",
	"酉戌",
}

//相害
func (r *DiZiXingChongHaiGet) GetXiangHai(lstDiZhi ...string) []string {
	return r.match(xiangHai_list, lstDiZhi...)
}

//相刑
func (r *DiZiXingChongHaiGet) GetXiangXing(lstDiZhi ...string) []string {
	return r.match(xiangXing_list, lstDiZhi...)
}

//相冲
func (r *DiZiXingChongHaiGet) GetLiuChong(lstDiZhi ...string) []string {
	return r.match(liuChong_list, lstDiZhi...)
}

func (r *DiZiXingChongHaiGet) match(patMap []string, lstDiZhi ...string) []string {
	ret := make([]string, 0)
	for k1, v1 := range lstDiZhi {
		for k2, v2 := range lstDiZhi {
			if k1 == k2 {
				continue
			}
			if g.InSlice(v1+v2, liuChong_list) {
				ret = append(ret, v1+v2)
			}
		}
	}

	return ret
}
