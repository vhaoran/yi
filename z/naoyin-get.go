package z

import (
	"strings"

	"github.com/vhaoran/yi/cmn"
	. "github.com/vhaoran/yi/z/model"
)

type (
	//---纳音------
	NaYinGet struct {
	}
)

var NaYinGetX = new(NaYinGet)
var mNaYin = cmn.KV{
	"甲子": "海中金",
	"甲午": "沙中金",
	"丙寅": "炉中火",
	"丙申": "山下火",
	"戊辰": "大林木",
	"戊戌": "平地木",
	"庚午": "路旁土",
	"庚子": "壁上土",
	"壬申": "剑锋金",
	"壬寅": "金箔金",
	"甲戌": "山头火",
	"甲辰": "覆灯火",
	"丙子": "涧下水",
	"丙午": "天河水",
	"戊寅": "城头土",
	"戊申": "大驿土",
	"庚辰": "白蜡金",
	"庚戌": "钗钏金",
	"壬午": "杨柳木",
	"壬子": "桑柘木",
	"甲申": "泉中水",
	"甲寅": "大溪水",
	"丙戌": "屋上土",
	"丙辰": "沙中土",
	"戊子": "霹雳火",
	"戊午": "天上火",
	"庚寅": "松柏木",
	"庚申": "石榴木",
	"壬辰": "长流水",
	"壬戌": "大海水",
	"乙丑": "海中金",
	"乙未": "沙中金",
	"丁卯": "炉中火",
	"丁酉": "山下火",
	"己巳": "大林木",
	"己亥": "平地木",
	"辛未": "路旁土",
	"辛丑": "壁上土",
	"癸酉": "剑锋金",
	"癸卯": "金箔金",
	"乙亥": "山头火",
	"乙巳": "覆灯火",
	"丁丑": "涧下水",
	"丁未": "天河水",
	"己卯": "城头土",
	"己酉": "大驿土",
	"辛巳": "白蜡金",
	"辛亥": "钗钏金",
	"癸未": "杨柳木",
	"癸丑": "桑柘木",
	"乙酉": "泉中水",
	"乙卯": "大溪水",
	"丁亥": "屋上土",
	"丁巳": "沙中土",
	"己丑": "霹雳火",
	"己未": "天上火",
	"辛卯": "松柏木",
	"辛酉": "石榴木",
	"癸巳": "长流水",
	"癸亥": "大海水",
}

func (r *NaYinGet) Get(z *SiZhuModel) []*KVRoot {
	ret := make([]*KVRoot, 0)
	//--------nian -----------------------------
	if str := r.GetSingle(z.NianZhu()); len(str) > 0 {
		ret = append(ret, &KVRoot{
			Name:    z.NianZhu(),
			Comment: str,
		})
	}
	//--------yue -----------------------------
	if str := r.GetSingle(z.YueZhu()); len(str) > 0 {
		ret = append(ret, &KVRoot{
			Name:    z.YueZhu(),
			Comment: str,
		})
	}
	//--------ri -----------------------------
	if str := r.GetSingle(z.RiZhu()); len(str) > 0 {
		ret = append(ret, &KVRoot{
			Name:    z.RiZhu(),
			Comment: str,
		})
	}
	//--------shi -----------------------------
	if str := r.GetSingle(z.ShiZhu()); len(str) > 0 {
		ret = append(ret, &KVRoot{
			Name:    z.ShiZhu(),
			Comment: str,
		})
	}

	return ret
}

func (r *NaYinGet) GetSingle(ganZi string) string {
	s, ok := mNaYin[ganZi]
	if ok {
		return s
	}
	return ""
}

func (r *NaYinGet) GetWuXin(nayin string) string {
	l := []string{"金", "木", "水", "火", "土"}
	for _, v := range l {
		if strings.Contains(nayin, v) {
			return v
		}
	}
	return ""
}
