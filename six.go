package main

//Вариант № 2 – использовать 1 канал для передачи данных, передать close.
func varTwoExT() chan int {

	varTTchan := make(chan int)

	go func() {
		numbEx := 1
		for {
			select {
			case varTTchan <- numbEx:
				numbEx++
			case <-varTTchan:
				return
			}
		}
	}()
	return varTTchan
}

/*func main() {

	//Вариант №1 – создать специальный канал для «закрытия» горутины.
	stopT := make(chan struct{})
	go func() {
		for {
			select {
			case <-stopT:
				close(stopT)
			}
		}
	}()
	//Запуск варианта №2
	numberForVatTwoT := varTwoExT()
	fmt.Println(<-numberForVatTwoT)
	fmt.Println(<-numberForVatTwoT)
	fmt.Println(<-numberForVatTwoT)
	close(numberForVatTwoT)

	//Вариант №3 – закрыть горутину по time out.
	intervalT := time.After(1 * time.Second)
	mainChannelT := make(chan int, 15)
	timeIsUpT := make(chan bool, 1)

	for i := 0; i <= 10; i++ {
		mainChannelT <- i
	}
	for {
		select {
		case num := <-mainChannelT:
			fmt.Println(num)
		case <-intervalT:
			timeIsUpT <- true
			return
		}
	}
}*/
