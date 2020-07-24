package cmn

//ganOrZhi干或支
//other 其它幹或支
func GetShiShen(ganOrZhi, other string) string {
	rx, same := GetRX(ganOrZhi, other)
	if rx == RX_WoSheng {
		if same {
			return "食神"
		}
		return "伤官"
	}
	if rx == RX_ShengWo {
		if same {
			return "偏印"
		}
		return "正印"
	}

	//-------- -----------------------------
	if rx == RX_KeWo {
		if same {
			return "七杀"
		}
		return "正官"
	}
	if rx == RX_WoKe {
		if same {
			return "偏财"
		}
		return "正财"
	}
	//
	if rx == RX_TongWo {
		if same {
			return "比肩"
		}
		return "劫财"
	}
	return ""
}

func GetShiShenShort(shiShen string) string {
	m := KV{
		"比肩": "比",
		"劫财": "劫",
		"食神": "食",
		"伤官": "伤",
		"偏财": "才",
		"正财": "财",

		"偏官": "杀",
		"七杀": "杀",
		"正官": "官",

		"偏印": "枭",
		"正印": "印",
	}
	s, ok := m[shiShen]
	if ok {
		return s
	}
	return ""
}
