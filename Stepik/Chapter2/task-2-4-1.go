// В рамках этого урока мы постарались представить себе уже привычные нам переменные и
// функции, как объекты из реальной жизни. Чтобы закрепить результат мы предлагаем
// вам небольшую творческую задачу.

// Вам необходимо реализовать структуру со свойствами-полями On, Ammo и Power,
// с типами bool, int, int соответственно.
// У этой структуры должны быть методы: Shoot и RideBike, которые не принимают аргументов,
// но возвращают значение bool.

// Если значение On == false, то оба метода вернут false.

// Делать Shoot можно только при наличии Ammo (тогда Ammo уменьшается на единицу,
// 	а метод возвращает true), если его нет, то метод вернет false.
// 	Метод RideBike работает также, но только зависит от свойства Power.

// Чтобы проверить, что вы все сделали правильно,
// вы должны создать указатель на экземпляр этой структуры с
// именем testStruct в функции main, в дальнейшем программа проверит результат.

package main

import "fmt"

type Biker struct {
	On    bool
	Ammo  int
	Power int
}

func (s *Biker) Shoot() bool {
	var shoot bool

	switch {
	case s.Ammo != 0 && s.On != false:
		s.Ammo--
		fmt.Printf("Стреляю! Осталось выстрелов: %d.\n", s.Ammo)
		shoot = true
	case s.Ammo == 0 && s.On == true:
		fmt.Print("Патронов нет!\n")
		shoot = false
	case s.On == false:
		fmt.Print("Стрелять нельзя!\n")
	}

	return shoot
}

func (r *Biker) RideBike() bool {
	var ride bool

	switch {
	case r.Power != 0 && r.On != false:
		r.Power--
		fmt.Printf("Еду! Осталось поездок: %d.\n", r.Power)
		ride = true
	case r.Power == 0 && r.On == true:
		fmt.Print("Поездок не осталось!\n")
		ride = false
	case r.On == false:
		fmt.Print("Ехать нельзя!\n")
	}

	return ride
}

func main() {

	testStruct := &Biker{true, 3, 2}

	for testStruct.Shoot() == true {
		if testStruct.Power != 0 {
			testStruct.RideBike()
		}
	}

	for testStruct.RideBike() == true {
		if testStruct.Ammo != 0 {
			testStruct.Shoot()
		}
	}

}
