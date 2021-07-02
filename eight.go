package main

import (
	"math"
)

/*
Чтобы обнулить какой-либо бит числа, нужно его логически умножить на 0.
Для этого само число логически умножаем результат из п. 1 на разность числа и степни двойки
(степень зависит от положения бита в числе).
*/
func bitToZero(bitNumber, actualNumber int64) (result int64) {

	powOfTwo := bitNumber - 1
	result = actualNumber & (actualNumber - int64(math.Pow(2, float64(powOfTwo))))

	return result

}

/*
Для установки битов в единицу используется побитовая логическая операция |
Решение: само число логически складываем с числом = степени 2 разряда.
*/

func bitToOne(bitNumber, actualNumber int64) (result int64) {

	powOfTwo := bitNumber - 1
	result = actualNumber | int64(math.Pow(2, float64(powOfTwo)))
	return result
}

/*func main() {

	fmt.Println(bitToZero(4, 15))
	fmt.Println(bitToOne(2, 1))

}*/
