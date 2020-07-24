package cmn

import (
	"strings"
)

/*  天十地支的阴阳属性  */
type ShuXing int

const (
	YIN  = 1
	YANG = 2
)

//计处天十、地支的阴阳属性
func GetShuXing(ganOrZhi string) ShuXing {
	//子丑寅卯辰巳午未申酉戌亥
	//甲乙丙丁戊己庚辛壬癸
	const yan = "子寅辰午申戌甲丙戊庚壬"
	const yin = "丑卯巳未酉亥乙丁己辛癸"
	if strings.Contains(yan, ganOrZhi) {
		return YANG
	}
	if strings.Contains(yin, ganOrZhi) {
		return YIN
	}
	return 0
}
