package y

var GONG8 = map[string]string{
	"111": "乾",
	"110": "兑",
	"101": "离",
	"100": "震",
	"011": "巽",
	"010": "坎",
	"001": "艮",
	"000": "坤",
}

//8宫8纯填的代码，用于查找伏神
var GONG8_8ChunGuaCode = map[string]string{
	"乾": "111111",
	"兑": "110110",
	"离": "101101",
	"震": "100100",
	"巽": "011011",
	"坎": "010010",
	"艮": "001001",
	"坤": "000000",
}

var GONG8_WUXING = map[string]string{
	"乾": "金",
	"兑": "金",
	"离": "火",
	"震": "木",
	"巽": "木",
	"坎": "水",
	"艮": "土",
	"坤": "土",
}

func Get8ChunGua(gong string) *Gua64 {
	code, ok := GONG8_8ChunGuaCode[gong]
	if ok {
		g, ok := G64[code]
		if ok {
			return g
		}
	}

	return nil
}

func GetGong8WuXing(gong string) string {
	s, ok := GONG8_WUXING[gong]
	if ok {
		return s
	}
	return ""
}
