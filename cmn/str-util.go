package cmn

import (
	"strings"

	"github.com/vhaoran/vchat/common/g"
)

func StrEachInSlice(s string, l ...string) bool {
	lst := strings.Split(s, "")
	ok := true
	for _, str := range lst {
		if !g.InSlice(str, l) {
			return false
		}
	}
	return ok
}

func StrIndexOfSlice(elem string, l ...string) int {
	for i, v := range l {
		if elem == v {
			return i
		}
	}
	return -1
}
