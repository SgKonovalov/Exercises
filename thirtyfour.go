package main

import (
	"fmt"
	"strings"
)

var sourseStrTF = "Abc d"

/*
Работаем в функции checkTf:
1) На вход передаём строку, которую хоти преобразовать;
2) Все буквы этой строки переводим в «строчные»;
3) Переводим получившуюся строку в п. 2 в срез байт;
4) Передаём каждый байт строки в map в качестве ключа;
5) Если длина среза (п. 3) равна длине map (п. 4) -> все символы в строке уникальны, если нет, то есть повторы.
*/

func checkTf(sourseStrTF string) {

	stringToByteTF := []byte(strings.ToLower(sourseStrTF))

	uniqueTF := make(map[byte]bool)

	for _, letterTF := range stringToByteTF {
		uniqueTF[letterTF] = true
	}

	if len(stringToByteTF) == len(uniqueTF) {
		fmt.Printf("All symbols in %s string are unique!", sourseStrTF)
	} else {
		fmt.Printf("Found repeated symbols in string '%s'.", sourseStrTF)
	}
}

/*func main() {
	checkTf(sourseStrTF)
}*/
