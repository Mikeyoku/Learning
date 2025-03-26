/*

Задача № 3. Вывести на экран в порядке возрастания три введенных числа

Пример:

Вход: 1 9 2

Выход: 1 2 9



Про sort мы пока ничего не знаем!

Это задача на условный оператор

*/

package main



// Ваш код



import "fmt"



func main () {

	var ( 

		in1, in2, in3, k int

	)



	fmt.Println(" Число 1: ")

	fmt.Scan(&in1)

	fmt.Println(" Число 2: ")

	fmt.Scan(&in2)

	fmt.Println(" Число 3: ")

	fmt.Scan(&in3)



	// проверка 1

	if in2 > in3 {

		k = in2

		in2 = in3

		in3 = k

	}

	if in1 > in3 {

		k = in1

		in1 = in3

		in3 = k

	}

	if in1 > in2 {

		k = in1

		in1 = in2

		in2 = k

	}



	fmt.Printf("Вывод: %v %v %v", in1, in2, in3)

}