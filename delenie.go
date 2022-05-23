package main

import (
	"fmt"
	"math/rand"
	"time"
)

func divide(numbers []int, divisible1 int, divisible2 int) (max1 int, max2 int) {
	max1 = 0
	max2 = 0
	for i := 0; i < 50000000; i++ {
		if numbers[i]%divisible1 == 0 {
			max1 += 1
		}
	}
	for i := 0; i < 50000000; i++ {
		if numbers[i]%divisible2 == 0 {
			max2 += 1
		}
	}
	return
}
func main() {
	numbers := make([]int, 50000000)
	rand.Seed(int64(time.Now().Second()))
	for i := 0; i < 50000000; i++ {
		numbers[i] = rand.Int()
	}
	divisible1 := 0
	divisible2 := 0
	fmt.Println("Введите число 1")
	fmt.Scanf("%d\n", &divisible1)
	fmt.Println("Введите число 2")
	fmt.Scanf("%d\n", &divisible2)
	max1, max2 := divide(numbers, divisible1, divisible2)
	switch {
	case max1 > max2:
		fmt.Println("Количество чисел делимых на 1 число больше и состовляет", max1)
	case max2 > max1:
		fmt.Println("Количество чисел делимых на 2 число больше и состовляет", max2)
	default:
		fmt.Println("Количество чисел делимых на 1 и 2 число равно", max1, max2)
	}
}
