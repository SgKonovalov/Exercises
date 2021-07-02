package main

import (
	"fmt"
	"sync"
)

/*
Возможны 2 варианта:
1) Конкурентная запись, используя Mutex.
	1. Инициализирует саму map и переменную типа sync.Mutex{}.
	2. При каждом обращении к элементам map, оборачиваем это обращение в Lock и Unlock.
2) Конкурентная запись, используя анонимную структуру.
	1. Создаём анонимную структуру у которой инициализируем анонимное поле sync.RWMutex для работы с Mutex,
	а так же map, с последующей инициализацией.
	2. Создаём метод writeInto() – который будет записывать данные в map. Во время записи, ставим Lock на map.
	3. Создаём метод  readFrom() – который будет считывать и выводить данные из map в консоль.
	Во время работы с map из этой функции, ставим блокировки RLock().
*/

var wgForSevenExe sync.WaitGroup
var channelForSevenExe = make(chan map[int]int)

var counterSecongVar = struct {
	sync.RWMutex
	mapSecondVar map[int]int
}{mapSecondVar: make(map[int]int)}

func writeInto(mapSecondVar map[int]int) map[int]int {
	counterSecongVar.Lock()

	for i := 0; i < 5; i++ {
		counterSecongVar.mapSecondVar[i] = i * 5
	}
	counterSecongVar.Unlock()

	return mapSecondVar
}

func readFrom(mapSecondVar map[int]int) {
	defer wgForSevenExe.Done()
	counterSecongVar.RLock()

	for i, k := range mapSecondVar {
		fmt.Printf("Key is: %d, value is: %d\n", i, k)
	}
	counterSecongVar.RUnlock()
}

//func main() {

/*
	// Вариант №1
		wgForSevenExe.Add(1)
		var counters = map[int]int{}
		mu := &sync.Mutex{}
		for i := 0; i < 5; i++ {
			go func(counters map[int]int, th int) {
				for j := 0; j < 5; j++ {
					mu.Lock()
					counters[th*10+j]++
					mu.Unlock()
				}
			}(counters, i)
		}
		fmt.Scanln()
		mu.Lock()
		fmt.Println("counters result", counters)
		mu.Unlock()*/

// Вариант №2
/*wgForSevenExe.Add(2)
	go func() {
		defer wgForSevenExe.Done()
		channelForSevenExe <- writeInto(counterSecongVar.mapSecondVar)
	}()
	go readFrom(<-channelForSevenExe)
	wgForSevenExe.Wait()
}*/
