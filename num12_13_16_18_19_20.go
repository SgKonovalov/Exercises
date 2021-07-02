/*
12.	Что выводит данная программа и почему?

func update(p *int) {
  b := 2
  p = &b
}

func main() {
  var (
     a = 1
     p = &a
  )
  fmt.Println(*p)
  update(p)
  fmt.Println(*p)
}
___________________________________________________________________________________________________________
ОТВЕТ:
Программа выведет: 1 и 1.

Причина: в функции update (p = &b) – мы не выполняем «разыменование», а присваиваем переменной b адрес p.

Решение: в теле функции update выполнить «разыменование». Т.е. код в функции будет выглядеть так:
b := 2
*p = b
____________________________________________________________________________________________________________

13.	Чем завершится данная программа?

func main() {
	wg := &sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup, i int) {
			fmt.Println(i)
			wg.Done()
		}(wg, i)
	}
	wg.Wait()
	fmt.Println("exit")
}
___________________________________________________________________________________________________________
ОТВЕТ:
Программа выведет:
1
4
0
2
3
fatal error: all goroutines are asleep - deadlock!

Причина: анонимной go func wg sync.WaitGroup передан как параметр
и этот параметр не инициализирован -> WaitGroup работает.

Решение:
Вариант № 1 - в сигнатуре функции go func изменить состав параметров:
go func(i int) {
			fmt.Println(i)
			wg.Done()
		}(i)
Вариант №2 - передать этой функции указатель на адрес wg sync.WaitGroup:
wg := &sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup, i int) {
			fmt.Println(i)
			wg.Done()
		}(wg, i)
	}
_______________________________________________________________________________________________________

16.	Что выведет программа данная программа?

func main() {
	n := new(int)
	*n = 0
	if true {
		*n = 1
		*n++
	}
	fmt.Println(*n)
}
___________________________________________________________________________________________________________
ОТВЕТ:
Программа выведет: 0.

Причина: объявление переменной в if «перекрыватет» исходное значение и -> он не используется за рамками этого if
и равно первоначальному значению, присвоенному при инициализации.

Решение:
Вариант №1 - в теле if, вместо n := 1, указать n = 1:
if true {
		n = 1
		n++
	}
Вариант №2 - работать с указателем на n:
func main() {
	n := new(int)
	*n = 0
	if true {
		*n = 1
		*n++
	}
	fmt.Println(*n)
____________________________________________________________________________________________________________

18.	Что выведет данная программа и почему?


func someAction(v []int8, b int8) {
  v[0] = 100
  v = append(v, b)
}

func main() {
  var a = []int8{1, 2, 3, 4, 5}
  someAction(a, 6)
  fmt.Println(a)
}
___________________________________________________________________________________________________________
ОТВЕТ:
Программа выведет: массив [100 2 3 4 5].

Причина: в методе мы работаем не с самим срезом, а его копией -> не изменяем значения
в содержимом основного массива.

Решение: работать через указатель:
func someAction(v *[]int8, b int8) {
    (*v)[0] = 100
    (*v) = append((*v), b)
}

func main() {
    var a = &[]int8{1, 2, 3, 4, 5}
    someAction(a, 6)
    fmt.Println(a)
}
____________________________________________________________________________________________________________

19.	К каким негативным последствиям может привести данный кусок кода и как это исправить?

var justString string
func someFunc() {
  v := createHugeString(1 << 10)
  justString = v[:100]
}

func main() {
  someFunc()
}
___________________________________________________________________________________________________________
ОТВЕТ:
Возникнет 2 проблемы:
1) panic: runtime error: slice bounds out of range [:100] with length x;
2) justString останется nil или заданным изначально значением.

Причина:
1) Длина массива, может не совпасть с длиной строки после побитового смещения;
2) В функции мы работаем с копией строки justString и в итоге, она не изменится, нужно работать с указателями.

Решение:
1) Изменить длину массива в v[:100] на v[:len(justString)];
2) Работать с указателями:
var justString *string

func someFunc() {
    v := createHugeString(1 << 10)
    justString = & v[:100]
}

func main() {
    someFunc()
}
____________________________________________________________________________________________________________

20.	Какой результат выполнения данного кода и почему?


func main() {
  slice := []string{"a", "a"}

  func(slice []string) {
     slice = append(slice, "a")
     slice[0] = "b"
     slice[1] = "b"
     fmt.Print(slice)
  }(slice)
  fmt.Print(slice)
}
___________________________________________________________________________________________________________
ОТВЕТ:
Программа выведет: 2 разных массива [b b a][a a].

Причина: в функции м работаем с копией этого массива, а надо с указателем.

Решение: работать через указатель:
func main() {
    slice := &[]string{"a", "a"}

    func(slice *[]string) {
        *slice = append(*slice, "a")
        (*slice)[0] = "b"
        (*slice)[1] = "b"
        fmt.Print(*slice)
    }(slice)
    fmt.Print(*slice)
}
____________________________________________________________________________________________________________
*/
