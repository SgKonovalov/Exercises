package main

/*
1) Создаём структуру PointTO с инкапсулированными параметрами x и y.
2) Инициализируем конструктор.
3) Создаём метод  calcDistTO у структуры PointTO,
который принимает значения координат и вычитывает их по 2-м сценариям:
	1. Если значение x <= y, то отнимаем х от у;
	2. Если х > у, то отнимаем у от х.
Возвращаем результат.
4) В main, при помощи конструктора создаём новый объект типа PointTO,
передавая в качестве аргумента значения для х и у.
5) Запускаем метод calcDistTO у этого объекта.
*/

type PointTO struct {
	x, y int
}

func NewstTO(x, y int) *PointTO {
	return &PointTO{
		x: x,
		y: y,
	}
}

func (s *PointTO) calcDistTO(x, y int) (resultTO int) {
	s.x = x
	s.y = y
	if x <= y {
		resultTO = y - x
	} else {
		resultTO = x - y
	}
	return resultTO
}

/*func main() {
	oTO := NewstTO(0, 3)
	fmt.Printf("Distance between point is %d.", oTO.calcDistTO(oTO.x, oTO.y))
}*/
