package y

import (
	"strings"
)

var Gong8WuXing2LiuQin = map[string]map[string]string{
	"金": map[string]string{
		"金": "兄弟",
		"水": "子孙",
		"木": "妻财",
		"火": "官鬼",
		"土": "父母",
	},
	"木": map[string]string{
		"木": "兄弟",
		"火": "子孙",
		"土": "妻财",
		"金": "官鬼",
		"水": "父母",
	},
	"水": map[string]string{
		"水": "兄弟",
		"木": "子孙",
		"火": "妻财",
		"土": "官鬼",
		"金": "父母",
	},
	"火": map[string]string{
		"火": "兄弟",
		"土": "子孙",
		"金": "妻财",
		"水": "官鬼",
		"木": "父母",
	},
	"土": map[string]string{
		"土": "兄弟",
		"金": "子孙",
		"水": "妻财",
		"木": "官鬼",
		"火": "父母",
	},
}

type Gua64 struct {
	//所属宫
	Gong string `json:"gong"`
	//宫内顺序
	GongOrder int `json:"gong_order"`
	//所属五行
	WuXing string `json:"wuxing"`
	//长名称
	Name string `json:"name"`
	//短名称
	NameS string `json:"name_short"`
	//六爻(包含了世应)
	L []string `json:"l"`
}

func NewGua64(gong string, gongOrder int, wuxing, name, names string, l ...string) *Gua64 {
	if len(l) < 6 {
		panic("不够六爻")
	}

	return &Gua64{
		Gong:      gong,
		GongOrder: gongOrder,
		WuXing:    wuxing,
		Name:      name,
		NameS:     names,
		L:         l,
	}
}

//查找填中缺失的六亲
func (r *Gua64) NoLiuQin() []string {
	pat := []string{
		"父母", "兄弟", "官鬼", "妻财", "子孙",
	}

	l := make([]string, 0)
	//
	for _, dst := range pat {
		ok := false
		for _, s := range r.L {
			if strings.Contains(s, dst) {
				ok = true
				break
			}
		}

		//没有找到，则添加
		if !ok {
			l = append(l, dst)
		}
	}

	return l
}

//得到各爻位上对应的伏神
func (r *Gua64) GetFuShen() []string {
	g := Get8ChunGua(r.Gong)

	no := r.NoLiuQin()

	//
	l := make([]string, 0)
	for _, s := range g.L {
		ok := false
		for _, v := range no {
			if strings.Contains(s, v) {
				ok = true
				break
			}
		}
		//找到缺失的六亲，则添加
		if ok {
			l = append(l, s)
			continue
		}
		l = append(l, "")
	}

	return l
}

//改变变卦 的六亲以使与原卦同
func (r *Gua64) TranLiuqinOfBiangua(src []string) []string {
	gongWuxing := GetGong8WuXing(r.Gong)
	liuQin := Gong8WuXing2LiuQin[gongWuxing]

	pat := []string{"兄弟", "子孙", "妻财", "官鬼", "父母"}
	ret := make([]string, 0)
	for _, yao := range src {
		s := yao
		for xin, dst := range liuQin {
			if strings.Contains(s, xin) {
				for _, q := range pat {
					s = strings.Replace(s, q, dst, -1)
				}
				break
			}
		}
		ret = append(ret, s)
	}

	return ret
}
