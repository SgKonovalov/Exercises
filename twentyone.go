package main

import "fmt"

/*
Для работы:
	1) Создаём сам массив чисел;
	2) Канал, типизированный эти массивом;
	3) Функция readFromArrTO – читает из массива и пишет данные в канал;
	4) Функция printTO – берёт число из канала и выводит в консоль.
*/

var arrayForTO = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
var chanTO = make(chan int, len(arrayForTO)+1)

func readFromArrTO(arrayForTO [10]int, chanTO chan int) {
	for _, value := range arrayForTO {
		chanTO <- value
	}
}

func printTO(number int) {
	fmt.Println(number)
}

/*func main() {

	readFromArrTO(arrayForTO, chanTO)
	for {
		select {
		case x := <-chanTO:
			printTO(x)
		default:
			return
		}
	}
}*/
