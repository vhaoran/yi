package y

import (
	"errors"
	"strconv"
	"strings"
)

func VerifyLiuYaoCode(code string) error {
	l := strings.Split(code, "")
	//must be len = 6
	if len(l) < 6 || len(l) > 6 {
		return errors.New("必须是六个数字(0-1)")
	}

	// muse be 0-1
	for _, v := range l {
		i, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return err
		}
		if !(i >= 0 && i <= 3) {
			return errors.New("输入的代码必须由0-1构成")
		}
	}

	return nil
}
