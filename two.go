package main

import (
	"fmt"
	"math"
	"os"
	"sync"
)

/*
Инициализируем необходимые переменные для работы:
	1. Инициализируем сам массив фиксированной длины и записываем в него нужные сообщения;
	2. Инициализируем буферизованный канал для «общения» горутин;
	3. Инициализируем WaitGroup, чтобы работать с горутинами в main.
*/
var wg sync.WaitGroup
var numbersChan = make(chan [len(numbers)]float64, len(numbers)+1)
var numbers = [5]float64{2, 4, 6, 8, 10}

/*
Функция sqaureNum ().
Принимает в качестве аргумента массив чисел.
В теле функции производим переборку полученного массива,
в процессе возводя в квадрат все его значения и перезаписывая их в этот же массив.

В результате работы функция возвращает «обновлённый» массив
со возведёнными в квадрат значениями элементов исходного массива.
*/
func sqaureNum(numbers [len(numbers)]float64) [len(numbers)]float64 {

	for i, number := range numbers {
		numbers[i] = math.Pow(number, 2)
	}

	return numbers
}

/*
Функция printResult()
Принимает в качестве аргумента массив чисел.
В теле функции производим переборку полученного массива и выводим в консоль порядковый номер элемента + 1
и само значение этого элемента.
*/
func printResult(numbers [len(numbers)]float64) {

	defer wg.Done()

	for _, number := range numbers {
		fmt.Fprintf(os.Stdout, "Square of %d is %d\n", int(math.Sqrt(number)), int(number))
	}
}

/*
В функции main():
1) Добавляем в WaitGroup «ожидание завершения» 2 горутин;
2) Функцию sqaureNum() оборачиваем в анонимную функцию и запускаем в горутине, чтобы избежать DeadLock,
параллельно уменьшая счётчик горутин на 1 и закрывая канал
(т.к. запись в него осуществляется именно в этой функции).
3) Запускаем в отдельной горутине функцию printResult(), в качестве параметра передаём запись из канала numbersChan,
полученную в результате отработки предыдущей горутины.
4) При помощи wg.Wait() уведомляем функцию main о необходимости ждать выполнения установленного в wg.Add()
количества горутин.
*/
/*func main() {

	wg.Add(2)
	go func() {
		defer close(numbersChan)
		defer wg.Done()
		numbersChan <- sqaureNum(numbers)
	}()

	go printResult(<-numbersChan)
	wg.Wait()
}*/
