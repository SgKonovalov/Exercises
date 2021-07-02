package main

/*
Для работы с большими числами, используем пакет math/big. Чтобы инициализировать переменную типа big.Int,
создаём её с помощью оператора new и переводим её из строки в.
Далее, для каждой математической операции инициализирум переменную и с помощью функций:
	1. Mul – умножаем;
	2. Div – делим;
	3. Add – складываем;
	4. Sub – вычитаем.
*/

/*func main() {

	var a, _ = new(big.Int).SetString("100", 10)
	var b, _ = new(big.Int).SetString("200", 10)

	multiply := new(big.Int)
	multiply.Mul(a, b)

	division := new(big.Int)
	division.Div(b, a)

	sum := new(big.Int)
	sum.Add(a, b)

	subtraction := new(big.Int)
	subtraction.Sub(b, a)

	fmt.Printf("a = %v\nb = %v\na * b = %v\nb / a = %v\na + b = %v\nb - a = %v\n", a, b, multiply, division, sum, subtraction)
}*/
