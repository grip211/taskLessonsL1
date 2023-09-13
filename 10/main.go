package main

import "fmt"

func main() {
	temperatures := []float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5} // слайс с дробными числами температуры
	m := map[int][]float32{}                                                     // объявляем пустой мапу для дробей

	// прогоняем слайс с дробями через функцию вычисления шага
	for i := -5; i <= 5; i++ {
		calculateStep(temperatures, m, i)
	}

	fmt.Println(m) // вывод результата
}

func calculateStep(temperatures []float32, m map[int][]float32, pace int) {
	var numbers []float32 //  пустой слайс с дробями
	// в цикле перебираем внешний слайс с дробями
	for _, t := range temperatures {
		l := t / 10         // вычисляем количество десятков в целой части дроби
		if int(l) == pace { // и проверяем, равна ли десятая часть шагу
			numbers = append(numbers, t) // если равна, то добавляем число в локальный слайс с дробями

		}
	}

	if len(numbers) > 0 { // если в локальном слайсе с дробями что-нибудь есть
		m[pace*10] = numbers // то добавляем это в мапу, в качестве ключа указывая шаг умноженный на 10
	}
}
