package main

/*
Основная функция для работы – insertIntoTempDiff ().
Принимает сами числа в качестве аргумента, на «выходе» получаем map с ключом int и элементами типа срез float32.
	1) Инициализируем map differented – которую будем выдавать в return;
	2) Инициализируем и заполняем значениями, полученными из параметра функции map – preparatoryMap.
	В качестве ключа записываем саму температуру, а в качестве значения – 1-ую цифру в значении температуры (учитывая её знак);
	3) Инициализируем 4 среза для наполнения ими значений в основной map differented;
	4) Перебираем map preparatoryMap по значениям. Если значение удовлетворяет условию if-а,
	то записываем его ключ в соответствующий срез;
	5) Заполняем map differented срезами, полученными в п. 4. В качестве ключа выбираем температуру с шагом в 10.
*/

var t1, t2, t3, t4, t5, t6, t7, t8 float32 = -25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5

func insertIntoTempDiff(t1, t2, t3, t4, t5, t6, t7, t8 float32) map[int][]float32 {

	differented := make(map[int][]float32)

	preparatoryMap := make(map[float32]int)
	preparatoryMap[t1] = int(t1 / 10)
	preparatoryMap[t2] = int(t2 / 10)
	preparatoryMap[t3] = int(t3 / 10)
	preparatoryMap[t4] = int(t4 / 10)
	preparatoryMap[t5] = int(t5 / 10)
	preparatoryMap[t6] = int(t6 / 10)
	preparatoryMap[t7] = int(t7 / 10)
	preparatoryMap[t8] = int(t8 / 10)

	var minusTwenty []float32
	var plusTen []float32
	var plusTwenty []float32
	var plusThirty []float32

	for temperatureActual, firstNumber := range preparatoryMap {
		if firstNumber < 0 {
			minusTwenty = append(minusTwenty, temperatureActual)
		} else if firstNumber == 1 {
			plusTen = append(plusTen, temperatureActual)
		} else if firstNumber == 2 {
			plusTwenty = append(plusTwenty, temperatureActual)
		} else if firstNumber >= 3 {
			plusThirty = append(plusThirty, temperatureActual)
		}
	}

	differented[-20] = minusTwenty
	differented[10] = plusTen
	differented[20] = plusTwenty
	differented[30] = plusThirty

	return differented
}

/*func main() {
	fmt.Println(insertIntoTempDiff(t1, t2, t3, t4, t5, t6, t7, t8))
}*/
