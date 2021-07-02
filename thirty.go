package main

var sliseT = []string{"Ivan", "Petr", "Vasiliy"}

/*
Варианты:
	1) Переносим нужный элемент в конец списка и вместо него записываем нулевое значение, а потом «усекаем» срез.
	Использование: если порядок элементов не важен.
	2) Выполняем сдвиг лево на 1 индекс, удаляем последний элемент и усекаем срез.
	Использование: если важен порядок элементов.
*/

func variantOneT(sliseT []string, elementNumber int) []string {

	sliseT[elementNumber] = sliseT[len(sliseT)-1]
	sliseT[len(sliseT)-1] = ""
	sliseT = sliseT[:len(sliseT)-1]
	return sliseT
}

func variantTwoT(sliseT []string, elementNumber int) []string {

	copy(sliseT[elementNumber:], sliseT[elementNumber+1:])
	sliseT[len(sliseT)-1] = ""
	sliseT = sliseT[:len(sliseT)-1]
	return sliseT
}

/*func main() {
	//fmt.Println(variantOneT(sliseT, 2))
	fmt.Println(variantTwoT(sliseT, 1))
}*/
