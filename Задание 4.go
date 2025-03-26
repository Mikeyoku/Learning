package main

import (
	"fmt"
	"strings"
)

func main() {
	// шаблон цикла
	// for init; condition; post {
	// init - блок инициализации переменных цикла
	// condition - уловие (если условие верно - тело цикла выполняется,
	// если - нет, цикл завершает свою работу)
	// post - изменение переменной цикла (инкрементарное или декрементарное действие)
	// }

	for i :=0; i <= 5; i++ {
		fmt.Println("Currebt value of i:", i)
	}

	// важный момент. в качестве init можно использовать выражение ТОЛЬКО через :=

	// break - команда, которая прерывает текущее выполнение цикла и
	// передает управление следующим за циклом инструкциям

	for i := 11; i <=22; i++ {
		if i > 20 {
			break
		}
		fmt.Println("Current value of i:", i)
	}
	fmt.Println("After for loop with BREAK")

	// continue - команда, которая прерывает текущее выполнение цикла и
	// передает управление следующей итерации цикла

	for i := 100; i <=120; i++ {
		if i > 110 && i <= 115 {
			continue
		}
		fmt.Println("Current value of i:", i)
	}
	fmt.Println("After for loop with CONTINUE")

	// вложенные циклы и лейблы
	for i := 0; i < 10; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
	fmt.Println("Выше должен быть треугольник")

	// Бывают ситуации, когда нужно прервать сразу оба цикла
	// Здесь помогут лейблы. Лейблы - это синтаксический сахар

	outer:
	for i := 0; i <=2; i++ {
		for j := 1; j <=3; j++ {
			fmt.Printf("i:%d and j:%d and sum i+j:%d\n", i, j, i+j)
			if i == j {
				break outer // Хочу, чтобы все циклы остановились, даже внешние
			}
		}
	}
	fmt.Println("After outer stop")

	// Шаблон цикла for как while (модификация for)
	// 1. Классический цикл while do
	var loopVar int = 0
	// while (loopVar <10) {
	// ...
	// loopVar++
	//}
	for loopVar <10 {
		fmt.Println("In while like loop", loopVar)
		loopVar++
	}
	fmt.Println(loopVar) 

	// loopVar - это переменная цикла, а не main (считает снизу вверх, разворачивает цикл изнутри)
	// это пример, можно, но не нужно
	// for loopVar := 0;loopVar <10; {
	// 	fmt.Println("In while like loop", loopVar)
	// 	loopVar++
	// }
	// fmt.Println(loopVar) 

	// 2. Классический бесконечный цикл
	var password string
	outher2:
		for {
			fmt.Println("Insert password:")
			fmt.Scan(&password)
			if strings.Contains(password, "1234") {
				fmt.Println("Weak password. Try again")
			} else {
				fmt.Println("Password accepted")
				break outher2
			}
		}

		// 3. Цикл с множественными переменными цикла
		for x, y := 0, 1; x<5 && y<8; x, y = x+1, y+2 {
			fmt.Printf("%d + %d = %d\n", x, y, x+y)
		}

		x1 := 5
		for x1 = 10, y1 := 10, 11; x1<15 && y1<18; x1, y1 = x1+1, y1+2 {
			fmt.Printf("%d + %d = %d\n", x1, y1, x1+y1)
		}
		fmt.Println(x1)
}