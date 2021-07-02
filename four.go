package main

import (
	"fmt"
)

var forJobsT = make(chan int, 100)
var forResultsT = make(chan int, 100)

/*
1) Функция workerF – считывает данные из канала для получения «работы» и выводит их на печать
так же записывает их в канал для результатов;
2) Функция workF:
	1. Вызывает установленное количество воркеров в цикле for
	(MAX кол-во воркеров устанавливается в аргументе функции: quanOfWorkers);
	2. Во 2-м цикле for создаёт необходимое количество «работы» для воркеров
	(MAX кол-во работы устанавливается в аргументе функции: quanOfWork)
	и отправляет всё это в канал forJobsT + закрывает этот канал, когда всё «запишет»,
	что и позволит завершить работу всех воркеров;
	3. В 3-м цикле for читает данные из канала forResultsTб по количеству установленной работы.
*/

func workerF(workerId int, forJobsT <-chan int, forResultsT chan<- int) {

	for work := range forJobsT {

		fmt.Printf("Worker № %d, processing work № %d\n", workerId, work)
		forResultsT <- work
	}
}

func workF(quanOfWorkers, quanOfWork int, forJobsT, forResultsT chan int) {

	for wq := 1; wq <= quanOfWorkers; wq++ {
		go workerF(wq, forJobsT, forResultsT)
	}

	for qw := 1; qw <= quanOfWork; qw++ {
		forJobsT <- qw
	}
	close(forJobsT)

	for qwr := 1; qwr <= quanOfWork; qwr++ {
		<-forResultsT
	}

}

/*func main() {
	workF(10, 50, forJobsT, forResultsT)
}*/
