package model

import (
	"fmt"
)

type (
	ShiShenModel struct {
		NianGan string
		YueGan  string
		RiGan   string
		ShiGan  string
	}
)

func (r *ShiShenModel) ToString() string {
	s := fmt.Sprint(r.NianGan, ",", r.YueGan, ",", r.RiGan, ",", r.ShiGan)
	return s
}
