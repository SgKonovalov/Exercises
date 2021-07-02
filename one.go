package main

import "fmt"

/*
При выборе варианта конструирования композиции, важно исходить из цели её создания, которыми могут быть:
1)	Конструирование композиции для простого встраивания методов и полей «родительской» структуры,
без использования возможностей полиморфизма. В данном случае возможны 2 варианта композиции:
	1.	Встроить именованное поле «родительского» типа (пример: HumanFirst/ActionFirst);
	2.	Встроить анонимное поле «родительского» типа (пример: HumanSecond/ActionSecond).
	Выбор одного из вышеперечисленных вариантов, обусловлен необходимостью использования в наследуемом типе
	нескольких полей типа «родителя», т.к. в Golang невозможно создать 2 и более анонимных полей в структуре.
2)	Конструирование композиции с возможностью дальнейшего использования полиморфизма
(пример: HumanThird/ActionThird).
*/

// Создаём структуру «родительского» типа Human. Определяем поля, которые будут доступны из структуры Action.
type HumanFirst struct {
	Name string
	Age  string
}

//Создаём структуру типа Action и встраиваем в неё именованное поле типа Human.
type ActionFirst struct {
	ActionsName string
	EmbHuman    HumanFirst
}

// Создаём метод для структуры Human, который будет доступен при работе с объектом типа Action.
func (h *HumanFirst) TestHumanFirst() {
	fmt.Printf("Hello, %s from HumanFirst by ActionFirst", h.Name)
}

// Создаём структуру «родительского» типа Human. Определяем поля, которые будут доступны из структуры Action.
type HumanSecond struct {
	Name string
	Age  string
}

//Создаём структуру типа Action и встраиваем в неё анонимное поле типа Human.
type ActionSecond struct {
	ActionsName string
	HumanSecond
}

//Создаём метод для структуры Human, который будет доступен при работе с объектом типа Action.
func (h *HumanSecond) TestHumanSecond() {
	fmt.Printf("Hello, %s from HumanSecond by ActionSecond", h.Name)
}

/*
Для создания композиции типов с использованием полиморфизма:
	1) Создаём интерфейс GroupAll, который объявляет метод TestHumanThird();
	2) Реализуем этот интерфейс в структуре HumanThird;
	3) Встраиваем анонимное поле типа HumanThird в структуру ActionThird.
В итоге получаем соответствие типов HumanThird и ActionThird интерфейсу GroupAll,
т.к. в типе ActionThird использовано анонимное поле HumanThird, реализующее интерфейс GroupAll.
*/

// Создаём сам интерфейс и описываем сигнатуру метода.
type GroupAll interface {
	TestHumanThird()
}

// Создаём структуру «родительского» типа Human. Определяем поля, которые будут доступны из структуры Action.
type HumanThird struct {
	Name string
	Age  string
}

//Создаём структуру типа Action и встраиваем в неё анонимное поле типа Human.
type ActionThird struct {
	ActionsName string
	HumanThird
}

/*
Создаём метод для структуры Human, который будет доступен при работе с объектом типа Action.
ПО сигнатуре этого метода и будет определён факт соответствия структуры HumanThird интерфейсу GroupAll.
*/
func (h *HumanThird) TestHumanThird() {
	fmt.Printf("Hello, %s from GroupAll", h.Name)
}

/*
Создаём функцию для проверки работы полимофризма.
Эта функция вызывает метод TestHumanFirst() структуры HumanThird,
принимая в качестве аргумента объект типа интерфейса GroupAll.
*/
func SayHello(g GroupAll) {
	g.TestHumanThird()
}

/*func main() {

	// Работа HumanFirst/ActionFirst
	af := ActionFirst{}
	af.EmbHuman.Name = "first user"
	af.EmbHuman.TestHumanFirst() // Вызываем метод TestHumanFirst() при обращении к имени поля структуры Action.
	// Работа HumanSecond/ActionSecond
	as := ActionSecond{}
	as.HumanSecond.Name = "second user"
	fmt.Println()
	as.HumanSecond.TestHumanSecond() // Вызываем метод TestHumanFirst() при обращении к типу поля структуры Action.

	// Работа HumanThird/ActionThird
	at := &ActionThird{}              // Создаём указатель на объект типа ActionThird
	ht := &HumanThird{}               // Создаём указатель на объект типа HumanThird
	at.HumanThird.Name = "HumanThird" // Устанавливаем значение поля Name для HumanThird
	ht.Name = "ActionThird"           // Устанавливаем значение поля Name для ActionThird
	fmt.Println()
	SayHello(at) // Вызываем функцию SayHello и в качетве параметра передаём объект типа ActionThird
	fmt.Println()
	SayHello(ht) // Вызываем функцию SayHello и в качетве параметра передаём объект типа HumanThird
}*/
