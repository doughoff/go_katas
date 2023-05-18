package main

import (
	"fmt"
	"strings"
)

func Accum(s string) string {
	letters := strings.Split(s, "")
	result := []string{}

	for pos, letter := range letters {
		temp := ""
		for i := 0; i <= pos; i++ {
			if i == 0 {
				temp += strings.ToUpper(letter)
			} else {
				temp += strings.ToLower(letter)
			}
		}

		result = append(result, temp)
	}

	return strings.Join(result, "-")
}

func main() {
	fmt.Println(Accum("ZpglnRxqenU"), "||", "Z-Pp-Ggg-Llll-Nnnnn-Rrrrrr-Xxxxxxx-Qqqqqqqq-Eeeeeeeee-Nnnnnnnnnn-Uuuuuuuuuuu")
	fmt.Println(Accum("NyffsGeyylB"), "||", "N-Yy-Fff-Ffff-Sssss-Gggggg-Eeeeeee-Yyyyyyyy-Yyyyyyyyy-Llllllllll-Bbbbbbbbbbb")
	fmt.Println(Accum("MjtkuBovqrU"), "||", "M-Jj-Ttt-Kkkk-Uuuuu-Bbbbbb-Ooooooo-Vvvvvvvv-Qqqqqqqqq-Rrrrrrrrrr-Uuuuuuuuuuu")
	fmt.Println(Accum("EvidjUnokmM"), "||", "E-Vv-Iii-Dddd-Jjjjj-Uuuuuu-Nnnnnnn-Oooooooo-Kkkkkkkkk-Mmmmmmmmmm-Mmmmmmmmmmm")
	fmt.Println(Accum("HbideVbxncC"), "||", "H-Bb-Iii-Dddd-Eeeee-Vvvvvv-Bbbbbbb-Xxxxxxxx-Nnnnnnnnn-Cccccccccc-Ccccccccccc")
}
