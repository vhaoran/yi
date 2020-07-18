package z

import (
	"github.com/vhaoran/vchat/common/g"

	. "github.com/vhaoran/yi/z/model"
)

type (
	DiZiXingChongHaiGet struct {
	}
)

var DiZiXingChongHaiGetX = new(DiZiXingChongHaiGet)

var liuChong_List = []string{
	//相冲，相冲，相冲，巳亥相冲
	"子午",
	"丑未",
	"寅申",
	"卯酉",
	"辰戌",
	"巳亥",
}

//六冲
func (r *DiZiXingChongHaiGet) GetLiChong(z SiZhuModel) []string {
	ret := make([]string, 0)

	l := z.ZhiList()
	for k1, v1 := range l {
		for k2, v2 := range l {
			if k1 == k2 {
				continue
			}
			if g.InSlice(v1+v2, liuChong_List) {
				ret = append(ret, v1+v2)
			}
		}
	}

	return l
}
