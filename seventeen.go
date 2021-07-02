package main

var value interface{} = 17

/*func main() {

	// Вариант №1 - использователь переключатель типа.
	switch search := value.(type) {
	case int:
		fmt.Println("This is int. With Value:", search)
	case string:
		fmt.Println("This is string. With Value:", search)
	case bool:
		fmt.Println("This is boolean. With Value:", search)
	case chan int:
		fmt.Println("This is channel int. With Value:", search)
	case chan string:
		fmt.Println("This is channel string. With Value:", search)
	case chan bool:
		fmt.Println("This is channel bool. With Value:", search)
	default:
		fmt.Println("Can't found actual type.")
	}

	//Вариант №2 - пакет reflction, который даёт полную информацию о типе.
	valueTypeVar2 := reflect.TypeOf(value)
	fmt.Println(valueTypeVar2)

	//Вариант №3 - использовать флаг %T.
	valueTypeVar3 := fmt.Sprintf("%T", value)
	fmt.Println(valueTypeVar3)
}*/
