package main

import (
	"fmt"
	"math/rand"
)

/*
Создаём:
1) Функцию randomThT – генерирует случайные числа. Для этого принимает 2 значения – начало и конец отрезка,
в рамках которого будут находиться сгенерированные числа.
2) Канал randChanThT – в него будем писать числа, сгенерированные randomThT.
3) Функцию checkEvenNumThT – проверяет числа на чётность. Если число чётное – возвращает само число,
если не чётное – возвращает 0.
4) Канал evenChanThT - в него будем писать числа, сгенерированные checkEvenNumThT.
5) Функцию printThT – выводит на печать полученные по каналу evenChanThT от функции checkEvenNumThT числа.
Если число != 0 – выводим на печать, если == 0, то выходим из функции.
*/

func randomThT(start, finish int) int {
	return rand.Intn(finish-start) + start
}

var randChanThT = make(chan int, 100)

func checkEvenNumThT(ranNum int) int {
	if ranNum%2 == 0 {
		return ranNum
	}
	return 0
}

var evenChanThT = make(chan int, 100)

func printThT(numberThT int) {
	if numberThT != 0 {
		fmt.Println(numberThT)
	} else {
		return
	}
}

/*func main() {

for i := 0; i < 100; i++ {
	randChanThT <- randomThT(1, 200000)
}
for {
	select {
	case number := <-randChanThT:
		evenChanThT <- checkEvenNumThT(number)
	case evenNumber := <-evenChanThT:
		printThT(evenNumber)
	default:
		return
	}
}
*/
