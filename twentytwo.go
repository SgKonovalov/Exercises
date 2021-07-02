package main

var arrTT1 = []int{7, 13, 1, -2, 0}
var arrTT2 = []float64{7.3, 13.5, 1.11, -2.20, 0.3}
var arrTT3 = []string{"Vasiliy", "Petr", "Ivan", "Andrei"}

type personsTT struct {
	Name string
	Age  int
}

var exeTT = []personsTT{
	{"Vasiliy", 33},
	{"Petr", 17},
	{"Ivan", 81},
	{"Andrei", 4},
}

/*func main() {
	// Сортировка int
	sort.Ints(arrTT1)
	fmt.Println(arrTT1)

	// Сортировка float64
	sort.Float64s(arrTT2)
	fmt.Println(arrTT1)

	// Сортировка string
	sort.Strings(arrTT3)
	fmt.Println(arrTT3)

	//Сортировка объектов стркутур
	sort.SliceStable(exeTT, func(i, j int) bool {
		return exeTT[i].Age < exeTT[j].Age
	})
	fmt.Println(exeTT)
}*/
