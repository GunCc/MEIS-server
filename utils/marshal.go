package utils

import (
	"fmt"
	"strconv"
)

// 接收数据 string 转 int
type JsonStringToInt string

func (j *JsonStringToInt) MarshalJSON() ([]byte, error) {
	fmt.Println("j", j)
	if *j != "" {
		tInt, _ := strconv.Atoi(string(*j))

		return []byte(fmt.Sprintf(`%d`, tInt)), nil
	}
	return []byte(fmt.Sprintf(`%d`, 0)), nil
}
