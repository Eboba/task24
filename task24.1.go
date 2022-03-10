package main

import (
	"fmt"
	"math/rand"
	"time"
)

const a = 10

func insertionSort(array [a]int) (output [a]int) {

	for i := 1; i < len(array); i++ {
		for j := i; j > 0; j-- {
			if array[j-1] > array[j] {
				array[j-1], array[j] = array[j], array[j-1]
			} else {
				break
			}
		}
	}

	return array
}

func main() {

	fmt.Println("24.5 Практическая работа")
	fmt.Println("")
	fmt.Println("Задание 1. Сортировка вставками.")
	fmt.Println("------------")

	rand.Seed(time.Now().UnixNano())

	var array [a]int

	for i, _ := range array {
		array[i] = rand.Intn(a) + 1
	}

	fmt.Println("Cгенерировать массив:", array)
	fmt.Println("Отсортированный массив:", insertionSort(array))
}
