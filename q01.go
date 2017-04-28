package main

import (
	"fmt"
	"strconv"
)

// 10進数、2進数、8進数のいずれで表現しても回文数となる数のうち、10進数の10以上で最小の値を求めてください。

func main() {
	i := 11

	for {
		binary := fmt.Sprintf("%b", i)
		octal := fmt.Sprintf("%o", i)

		if binary == reverse(binary) &&
			octal == reverse(octal) &&
			strconv.Itoa(i) == reverse(strconv.Itoa(i)) {
			break
		}

		i += 2
	}

	fmt.Println(i)
}

func reverse(s string) string {
	data := []rune(s)
	res := []rune{}

	for i := len(data) - 1; i >= 0; i-- {
		res = append(res, data[i])
	}
	return string(res)
}
