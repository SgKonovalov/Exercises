package main

import (
	"fmt"
	"regexp"
	"strings"
)

const sourceStrTS = "snow dog sun"
const targetStrTS = "sun dog snow"

/*
Вариант 1 – простая замена «нужных» слов в строке, при помощи strings.NewReplacer.
В качестве аргументов этой функции передаём пару «слово-замена»
*/

func firstVarTS(sourceStrTS string) {
	replaceTF := strings.NewReplacer("snow", "sun", "sun", "snow")
	resultStrTSOne := replaceTF.Replace(sourceStrTS)
	fmt.Printf("Before is: '%s', after is: '%s'.(By 1 variant).\n", sourceStrTS, resultStrTSOne)
}

/*
Вариант №2 – замена с использованием регулярного выражения:
1) В regexp.MustCompile передаём строку, которую нужно изменить;
2) ReplaceAllString – функция, которая заменяет текст всех совпадений регулярного выражения строкой замены.
*/

func secondVarTS(sourceStrTS, targetStrTS string) {
	regularExTS := regexp.MustCompile(sourceStrTS)
	resultStrTSTwo := regularExTS.ReplaceAllString(sourceStrTS, targetStrTS)
	fmt.Printf("Before is: '%s', after is: '%s'.(By 2 variant).\n", sourceStrTS, resultStrTSTwo)
}

/*func main() {

	firstVarTS(sourceStrTS)
	secondVarTS(sourceStrTS, targetStrTS)

}*/
