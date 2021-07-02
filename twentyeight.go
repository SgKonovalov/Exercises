package main

import "fmt"

/*
1. Описываем структуру builder, которая будет запускает обрабатывает запросы;
2. В структуре builder определяем метод getOrder, который вызывает метод build
из интерефейса house.
3. Интерфейс house  целевой интерфейс к которому мы приводим наши структуры.
4. Описываем структуры oneFloorHouse и twoFloorHouse – это струткруы, которы мы хотим привести к house.
5. Для приведения к типу house, реализовываем у структур oneFloorHouse и twoFloorHouse метод buid.
6. С помощью вспомогательной структуры secondFloorAdapter, приводим структуру twoFloorHouse к типу,
который ожидает метод getOrder.
*/

type builder struct {
}

type house interface {
	build()
}

type oneFloorHouse struct{}
type twoFloorHouse struct{}

func (b *builder) getOrder(house house) {
	fmt.Println("Builder get order.")
	house.build()
}

func (ofh *oneFloorHouse) build() {
	fmt.Println("Oops! We builds two floor...")
}

func (tfh *twoFloorHouse) build() {
	fmt.Println("That's ok. It's two floor.")
}

type secondFloorAdapter struct {
	secondFloor *twoFloorHouse
}

func (sfa *secondFloorAdapter) build() {
	fmt.Println("We are converted our one-floor house to two-floor.")
	sfa.secondFloor.build()
}

/*func main() {

	letsBuld := &builder{}
	oneFloorHouse := &oneFloorHouse{}

	letsBuld.getOrder(oneFloorHouse)

	twoFloorHouse := &twoFloorHouse{}
	twoFloorHouseAdapter := &secondFloorAdapter{
		secondFloor: twoFloorHouse,
	}

	letsBuld.getOrder(twoFloorHouseAdapter)
}*/
