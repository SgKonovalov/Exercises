package main

import (
	"fmt"
	"math"
	"os"
	"sync"
)

var a, b, c, d, e = 2, 4, 6, 8, 10

/*
Последовательность действий:
1) Инициализируем саму последовательность чисел;
2) Указанную последовательность поместить в срез и возвести каждый элемент этого среза в квадрат, записать в канал;
3) Получить срез из канала и вычислить сумму элементов среза. Записать результат в канал;
4) Получить  значение суммы элементов среза в канале и вывести результат в консоль.
*/
var Wg sync.WaitGroup
var NumbersChan = make(chan []float64, len(Numbers)+1)
var ForSumm = make(chan int, 2)
var Numbers []float64

/*
Помещаем последовательность чисел в срез и вычисляем их квадраты.
*/
func SqaureNum(a, b, c, d, e int) []float64 {

	Numbers = append(Numbers, float64(a), float64(b), float64(c), float64(d), float64(e))

	for i, number := range Numbers {
		Numbers[i] = math.Pow(number, 2)
	}

	return Numbers
}

/*
Вычисляем сумму квадратов чисел из исходной последовательности.
*/
func Sum(Numbers []float64) int {

	result := 0
	for _, number := range Numbers {
		result += int(number)
	}
	return result
}

/*
Выводим на печать.
*/
func PrintResult(result int) {

	defer Wg.Done()

	fmt.Fprintf(os.Stdout, "Summ is %d", result)
}

/*func main() {

	Wg.Add(3)
	go func() {
		defer close(NumbersChan)
		defer Wg.Done()
		NumbersChan <- SqaureNum(a, b, c, d, e)
	}()

	go func() {
		defer close(ForSumm)
		defer Wg.Done()
		ForSumm <- Sum(<-NumbersChan)
	}()

	go PrintResult(<-ForSumm)

	Wg.Wait()
}*/
