package main

var arrTTe1 = []int{7, 13, 1, -2, 0}
var arrTTe2 = []float64{7.3, 13.5, 1.11, -2.20, 0.3}
var arrTTe3 = []string{"Vasiliy", "Petr", "Ivan", "Andrei"}

type personsTTe struct {
	Name string
	Age  int
}

var exeTTe = []personsTTe{
	{"Vasiliy", 33},
	{"Petr", 17},
	{"Ivan", 81},
	{"Andrei", 4},
}

/*
   Бинарный поиск возможен только в отсортированном массиве,
   поэтому первоначально производим сортировку встроенными методами языка.
*/
/*func main() {

	sort.Ints(arrTTe1)
	sort.Float64s(arrTTe2)
	sort.Strings(arrTTe3)
	//Бинарный поиск для int
	fmt.Printf("In Array: %v, number 13 have positioned is: %d\n", arrTTe1, sort.SearchInts(arrTTe1, 13))

	//Бинарный поиск для float64
	fmt.Printf("In Array: %v, number 7.3 have positioned is: %d\n", arrTTe2, sort.SearchFloat64s(arrTTe2, 7.3))

	//Бинарный поиск для string
	fmt.Printf("In Array: %v, string `Petr` position is: %d\n", arrTTe3, sort.SearchStrings(arrTTe3, "Petr"))

	//Бинарный поиск для любого типа
	searchedAgeTTe := 81
	sortedTTe := sort.Search(len(exeTTe), func(i int) bool {
		return searchedAgeTTe <= exeTTe[i].Age
	})
	if sortedTTe < len(exeTTe) && exeTTe[sortedTTe].Age == searchedAgeTTe {
		fmt.Printf("Searched person: name - %s, age - %d\n", exeTTe[sortedTTe].Name, exeTTe[sortedTTe].Age)
	} else {
		fmt.Printf("Can't found person by age - %d\n", exeTTe[sortedTTe].Age)
	}
}*/
