/*

Задача № 4. Проверить, является ли четырехзначное число палиндромом

Пример:

Вход: 1221  Выход: 1221 - палиндром

Вход: 1234  Выход: 1234 - не палиндром

*/

package main



// Ваш код

import "fmt"



func main() {

	var in1 int

	

	fmt.Println("Введите число: ")

	fmt.Scan(&in1)



	out1 := in1/1000

	out2 := in1%1000/100

	out3 := in1%100/10

	out4 := in1%10



	if out1 == out4 && out2 == out3 {

		fmt.Println("Число палиндром")

	} else {

		fmt.Println("Число не палиндром")

	}

}