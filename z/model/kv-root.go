package model

import (
	"fmt"
)

type (
	KVRoot struct {
		Name    string `json:"name"`
		Comment string `json:"comment"`
	}
)

func (r *KVRoot) ToString() string {
	return fmt.Sprint(r.Name, "(", r.Comment, ")")
}
