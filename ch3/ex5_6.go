/*Используя пример программы выше напишите программу,
переводящую температуру из градусов Фаренгейта в градусы Цельсия. (C = (F - 32) * 5/9) */
package main

import "fmt"

func main() {
	var farenheit float64
	fmt.Println("vvedite farenheity")
	fmt.Scanln(&farenheit)
	C := ((farenheit - 32) * 5 / 9)
	fmt.Println(C, "gradusov")
	izmerenie()
}

//Напишите другую программу для перевода футов в метры (1 фут = 0.3048 метр).

func izmerenie() {
	var fut float64
	fmt.Println("vvedite kol-vo futov")
	fmt.Scanln(&fut)
	metr := (fut * 0.3048)
	fmt.Println(metr, "metrov")
}
