/*

Задача № 2. Получить реверсную запись трехзначного числа

Пример: 

вход: 346, выход: 643

вход: 120, выход: 021

вход: 100, выход: 001

*/

package main



// Ваш код



import "fmt"



var in1 = 346

var in2 = 120

var in3 = 100



func main() {

fmt.Println(in1%10,in1/10%10,in1/100)

fmt.Println(in2%10,in2/10%10,in2/100)

fmt.Println(in3%10,in3/10%10,in3/100)

}