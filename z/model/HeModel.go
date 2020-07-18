package model

type (
	HeModel struct {
		He          string
		HeHuaWuXing string
	}
)

func (r *HeModel) ToString() string {
	return r.He + "化" + r.HeHuaWuXing
}
