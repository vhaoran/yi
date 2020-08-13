package y

//64填编号 ，名称
//64填的名称
type Gua64Name struct {
	//矢名称
	Short string
	//长名称
	Long string
}

//64填的构造
func NewGua64Name(s, l string) *Gua64Name {
	return &Gua64Name{
		Short: s,
		Long:  l,
	}
}
