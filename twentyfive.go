package main

import (
	"sync"
)

/*
1) Создаём структуру со встроенным RWMutex и переменной для самого счётчика;
2) Метод структуры incremTF – принимает 1 (как начало отсчёта) и изменяет значение в поле value структуры счётчика.
Для корректной работы, используем Lock.
3) Метод showTF – возвращает значение счётчика.
Для этого, получает доступ к полю value структуры счётчика. Для корректной работы, используем RLock.
4) В main:
	1. Создаём объект типа счётчика;
	2. В цикле (100 запусков) – запускаем анонимную горутину, которая уже в цикле (100 запусков) запускает
	incremTF для увеличения значения на 1, т.к. в параметры incremTF передаём именное единицу.
	3. Чтобы увидеть корректный результат работы программы, «засыпаем» на 1 секунду.
*/

type countTF struct {
	muTF  sync.RWMutex
	value int64
}

func (ctf *countTF) incremTF(value int64) {
	ctf.muTF.Lock()
	defer ctf.muTF.Unlock()
	ctf.value += value
}

func (ctf *countTF) showTF() (value int64) {
	ctf.muTF.RLock()
	defer ctf.muTF.RUnlock()
	value = ctf.value
	return value
}

/*func main() {

	startTF := countTF{}

	for i := 0; i < 100; i++ {
		go func() {
			for i := 0; i < 1000; i++ {
				startTF.incremTF(1)
			}
		}()
	}
	time.Sleep(time.Second)
	fmt.Println(startTF.showTF())
}*/
