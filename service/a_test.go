package service

import (
	"fmt"
	"testing"
)

func Test_jie_get(t *testing.T) {
	obj := new(JieGetImpl)
	l := obj.Exec(2050)
	for _, v := range l {
		fmt.Println(v.ToString())
	}
}
