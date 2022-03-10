package main

import (
	"fmt"
	"math/rand"
	"time"
)

const a = 10

func bubbleSort(array [a]int) [a]int {

	for j := 1; j < len(array); j++ {
		f := 0

		for i := 0; i < len(array)-j; i++ {
			if array[i] < array[i+1] {
				array[i], array[i+1] = array[i+1], array[i]
				f = 1
			}
		}
		if f == 0 {
			break
		}
	}
	return array
}

func main() {

	fmt.Println("24.5 Практическая работа")
	fmt.Println("")
	fmt.Println("Задание 2. Анонимные функции.")
	fmt.Println("------------")

	rand.Seed(time.Now().UnixNano())

	var array [a]int

	for i, _ := range array {
		array[i] = rand.Intn(a) + 1
	}

	f := bubbleSort(array) // согласно тз =) *Напишите анонимную функцию, которая на вход получает массив типа integer, сортирует его пузырьком и переворачивает (либо сразу сортирует в обратном порядке, как посчитаете нужным).*
	fmt.Println("Cгенерировать массив:", array)
	fmt.Println("Отсортированный массив:", f)
}
