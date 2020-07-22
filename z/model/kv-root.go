package model

import (
	"fmt"
)

type (
	KVRoot struct {
		Name    string
		Comment string
	}
)

func (r *KVRoot) ToString() string {
	return fmt.Sprint(r.Name, "(", r.Comment, ")")
}
