package main

/*
Для работы:
	1) Из последовательности формируем срез;
	2) Инициализируем карту, где ключ – это само значение строки, а значение – bool.
	3) Перебираем массив и устанавливаем в качестве ключа – значение из элементов массива,
	а в качестве значения элемента map – true.
В итоге, получаем простое множество, где значения не повторяются.
*/

var cat1, cat2, dog, cat3, tree = "cat", "cat", "dog", "cat", "tree"
var allValues = []string{cat1, cat2, dog, cat3, tree}

/*func main() {

	setOfValues := make(map[string]bool)
	for _, value := range allValues {
		setOfValues[value] = true
	}
	fmt.Println(setOfValues)
}*/
