package nbak

type (
	GanZi struct {
		//年十
		NGang string
		//年支
		NZhi string
		//月干
		YGang string
		//月支
		YZhi string
		//日干
		RGang string
		//日支
		RZhi string
	}
)

func GetGanZi() (GanZi, error) {
	bean := GanZi{

	}

	//1900-07-01
	//庚子年
	//壬午月
	//乙亥日
	return bean, nil
}
