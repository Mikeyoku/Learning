package main

import (
	"fmt"
	"strconv"
	
)

func main() {
	// пункт задачи 3.2
	var (
		a, b, tablUmnoj int
	)
	fmt.Printf("Ввод a: ")
	fmt.Scan(&a)
	fmt.Printf("Ввод b: ")
	fmt.Scan(&b)

	if b > a {	// пункт задачи 3.1
	for a:=a; a <= b; a++ {
		if a% 5 ==0 {	// п.3.2 вывод числа деленным на 5 без остатка
			fmt.Println(a)
		}	
	}
	}

	// пункт задачи 3.3
	fmt.Printf("Ввод числа для таблицы умножения: ") 
	fmt.Scan(&tablUmnoj)
	i := 1
	
	for tablUmnoj := tablUmnoj; i <= 10; i++ {
		j := tablUmnoj * i
		fmt.Println(j)
	}

	// пункт задачи 3.4

	var (
		exam int
		sum1 int
	)
	for i := 1; i <= 4; i++ {
		fmt.Println("Введите оценку: ", i)
		fmt.Scan(&exam)
		sum1 += exam
	}
	fmt.Println("Сумма набранных баллов: ", sum1)


	//пункт задачи 3.5

	var (
		sum2, infinityCycle int
	)
	fmt.Printf("Введите ещё число отличное от -1: ") 
	for infinityCycle != -1 {
		sum2 += infinityCycle
		fmt.Scan(&infinityCycle)
		if infinityCycle != -1 {
		fmt.Printf("Введите ещё число: ") 
		}
	}
	fmt.Printf("Сумма чисел: %d\n", sum2)

	// пункт задачи 3.6
	var (
		natureNumber int
		summa int
		stringNumber string
	)

	fmt.Printf("Введите натуральное число: ")
	fmt.Scan(&natureNumber)
	stringNumber = fmt.Sprintf("%d", natureNumber)
	for i := 0; i < len(stringNumber); i++ {
		k := fmt.Sprintf("%c", stringNumber[i])
		var conv, _= strconv.Atoi(k)
		summa += conv
	}
	fmt.Printf("Cумма цифр %d: %d.\n", natureNumber, summa)



	
}