package main

/*
Варианты:
1) Функция reverseTSVarOne – в качестве аргумента принимает исходную строку и на «выходе» получаем «перевёрнутую».
	1. Т.к. string – это массив байтов -> мы можем перебрать всё её элементы (буквы);
	2. В процессе переборки будем записывать в переменную result все элементы исходной строки в обратном прядке,
	т.к. переборка начинается с последнего элемента.
2) Функция reverseTSVarTwo – в качестве аргумента принимает исходную строку и на «выходе» получаем «перевёрнутую».
	1. Переводим исходную строку в массив rune (присваивает кодовую точку Unicode) в цикле for;
	2. «Переворачиваем» строку в цикле for, присваивая первому элементу значение последнего;
	3. Переводим из rune обратно в string.
*/

var sourceTS = "Hello world"

//Реализация первого варианта
func reverseTSVarOne(sourceTS string) (result string) {
	for _, letter := range sourceTS {
		result = string(letter) + result
	}
	return result
}

//Реализация второго варианта
func reverseTSVarTwo(sourceTS string) (result string) {
	letterPosTS := 0

	runeTS := make([]rune, len(sourceTS))

	for _, letter := range sourceTS {
		runeTS[letterPosTS] = letter
		letterPosTS++
	}

	for i := 0; i < letterPosTS/2; i++ {
		runeTS[i], runeTS[letterPosTS-1-i] = runeTS[letterPosTS-1-i], runeTS[i]
	}

	result = string(runeTS)
	return result

}

/*func main() {

	fmt.Println(reverseTSVarOne(sourceTS))
	fmt.Println(reverseTSVarTwo(sourceTS))
}*/
