package main

/*
За переборку отвечает функция arrayPart(), в качестве аргументов, передаём массивы, которые нужно перебрать;
на выходе – «выдаём» готовый массив с набором пересечений (если они есть).
	1) Изначально определяем какой из массивов длиннее, для этого просто сравниваем len.
	Цель – начать переборку с меньшего массива, чем ускорить процесс сравнения.
	2) Перебираем массива в цикле, если находим одинаковые элементы, - помещаем в «выходной» массив,
	параллельно присваивая индексы;
	3) Проверяем добавленные элементы в if.Если количество элементов, добавленных в п. 2,
	меньше исходных значений – продолжаем работать с циклом. Если больше или равное - выходим из цикла и метода,
	используя bool значение ok.
ПРИМЕЧАНИЕ: проверять равенство значений в строке 20 нельзя,
т.к. «выйдем» за пределы массива: index out of range [1] with length 1
*/

func arrayPart(first []string, second []string) (result []string) {

	smaller, bigger := first, second
	if len(first) > len(second) {
		smaller = second
		bigger = first
	}

	ok := false
	for fen, eofa := range smaller {
		for sen, eosa := range bigger {
			fen1 := fen + 1
			sen2 := sen + 1
			if eofa == eosa {
				result = append(result, eosa)
				if fen1 < len(smaller) && sen2 < len(bigger) {
					if smaller[fen1] != bigger[sen2] {
						ok = true
					}
				}
				break
			}
		}
		if ok {
			break
		}
	}
	return
}

/*func main() {
	nameOne := []string{"Ivan Ivanov", "Petr Petrov"}
	namesTwo := []string{"Petr Petrov"}
	fmt.Printf("%+v\n", arrayPart(nameOne, namesTwo))
}*/
