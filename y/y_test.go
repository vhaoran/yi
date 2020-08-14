package y

import (
	"fmt"
	"strings"
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func Test_GetLiuShen(t *testing.T) {
	//
	l := GetLiuShen("甲")
	fmt.Println("-----------------")
	spew.Dump(l)
}

func Test_NoLiuQin(t *testing.T) {
	c := 0
	for i := 0; i < 64; i++ {
		code := fmt.Sprintf("%06b", i)
		g := GetG64(code)
		if g != nil {
			l := g.NoLiuQin()
			if len(l) > 0 {
				c++
				fmt.Println("----", c, "----", g.Name, "-------------", code)
				fmt.Println(strings.Join(l, ","))
				fmt.Println("-----")
				fmt.Println(strings.Join(g.L, "\n\r"))
				fu := g.GetFuShen()
				if len(fu) > 0 {
					fmt.Println(strings.Join(fu, "/"))
				}
			}
		}
	}
}
func Test_BianGua(t *testing.T) {
	for i := 0; i < 64/2; i++ {
		code := fmt.Sprintf("%06b", i)
		g := GetG64(code)

		codeB := fmt.Sprintf("%06b", 64-i-1)
		b := GetG64(codeB)

		lb := g.TranLiuqinOfBiangua(b.L)
		//-------------------------------
		fmt.Println("-----------------")
		fmt.Println(fmt.Sprintf(" %s(%s)   %s(%s)", g.Name, g.Gong, b.Name, b.Gong))
		for i, v := range g.L {
			fmt.Println(v, "  ---  ", lb[i], "--", b.L[i])
		}
	}
}
