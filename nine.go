package main

import (
	"fmt"
)

/*
Для работы нужны:
	1) nmbrsArr – сам массив чисел;
	2) nmbrsArrChan – буферизованный канал с ёмкостью = длине nmbrsArr.
	Будет использован для записи чисел из массива;
	3) multChan - – буферизованный канал с ёмкостью = длине nmbrsArr.
	Будет использован для записи элементов массива * 2.
	4) imprortFromSrcArr – функция, которая записывает в канал nmbrsArrChan элементы массива.
	5) fromArrWithMult – функция, которая записывает в канал multChan элементы массива * 2.
	6) printNumb – функция для печати чисел.
В main:
	1) Запускаем функции imprortFromSrcArr и fromArrWithMult;
	2) В цикле Select выбираем готовый канал и включаем метод printNumb при наличии готового канала.
	Если каналы пустые – выходим из метода.
*/

var nmbrsArr = [5]int{2, 3, 4, 5, 6}
var nmbrsArrChan = make(chan int, len(nmbrsArr))
var multChan = make(chan int, len(nmbrsArr))

func imprortFromSrcArr(nmbrsArr [5]int, nmbrsArrChan chan int) {
	for _, number := range nmbrsArr {
		nmbrsArrChan <- number
	}
}

func fromArrWithMult(nmbrsArr [5]int, multChan chan int) {
	for _, number := range nmbrsArr {
		multChan <- number * 2
	}
}

func printNumb(number int) {
	fmt.Println(number)
}

/*func main() {

	imprortFromSrcArr(nmbrsArr, nmbrsArrChan)
	fromArrWithMult(nmbrsArr, multChan)

	for {

		select {
		case x := <-nmbrsArrChan:
			printNumb(x)
		case y := <-multChan:
			printNumb(y)
		default:
			return
		}
	}
}*/
