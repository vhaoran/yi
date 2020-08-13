package y

type Gua64 struct {
	//所属宫
	Gong string `json:"gong,omitempty"`
	//宫内顺序
	GongOrder int `json:"gong_order"`
	//所属五行
	WuXing string `json:"wuxing"`
	//长名称
	Name string `json:"name"`
	//短名称
	NameS string `json:"name_short"`
	//六爻(包含了世应)
	L []string `json:"l,omitempty"`
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
