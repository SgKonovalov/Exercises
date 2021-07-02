package main

import (
	"fmt"
	"os"
	"time"
)

/*
Для работы программы:
1) Инициализируем:
	1. Константу numberOfData – где задаём количетсво передаваемых данных в канале;
	2. Переменную interval – где задаём время работы программы;
	3. Буферизированный канал  mainChannel –размером numberOfData+1, в который будем записывать сами данные;
	4. Буферизированный канал timeIsUp – на 1 значение. В нём будем передавать true, если interval истёк.
2) Создаём функцию readFromChannel (), принимающую в качетсве аргумента переменную типа данных (в нашем сучае int).
Сама функция просто выводит в консоль полученный результат.
*/
const numberOfData = 10000

var interval = time.After(3 * time.Second)
var mainChannel = make(chan int, numberOfData+1)
var timeIsUp = make(chan bool, 1)

func readFromChannel(num int) {

	fmt.Fprintln(os.Stdout, num)
}

/*
В функции main():
	1) В цикле for записываем данные в канал;
	2) Далее в бесконечном цикле for, организуем select по каналам.
	Если появляются данные в mainChannel, считываем из него значение в перменную num,
	которую в теле case передаём на выполнение в функции readFromChannel();
	Т.к. цикл бесконечный, то проверяем наличие true в канале timeIsUp.
	Если значение true оказывается в канале, то выходим из метода main.
*/
/*func main() {

	for i := 0; i <= numberOfData; i++ {
		mainChannel <- i
	}
	for {
		select {
		case num := <-mainChannel:
			readFromChannel(num)
		case <-interval:
			timeIsUp <- true
			return
		}
	}
}*/
