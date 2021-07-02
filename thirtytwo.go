package main

import (
	"time"
)

/*
В функции sleepTtT:
1) Принимаем в качестве аргумента время, на которое хотим остановить горутину;
2) Устанавливаем таймер на этот промежуток;
3) Как только таймер закончил отсчёт, передаём информацию в канал горутины.
*/

func sleepTtT(forTime time.Duration) time.Time {

	timer := time.NewTimer(forTime)
	return <-timer.C
}

/*func main() {

	sleepTtT(5 * time.Second)

	fmt.Println("Hello world!")

}*/
