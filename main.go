package main

import "fmt"

//array2
//func main() {
//	var n int
//	fmt.Println("Введите число: ")
//	fmt.Scan(&n)
//	numbers := []float64{}
//	for i := 1; i <= n; i++ {
//		num := math.Pow(float64(n), float64(i))
//		numbers = append(numbers, num)
//	}
//	fmt.Println(numbers)
//}

// array14
//func main() {
//	var n int
//	fmt.Println("Введите число: ")
//	fmt.Scan(&n)
//	numbers := []float64{}
//	for i := 1; i <= n; i++ {
//		if i%2 == 0 {
//			numbers = append(numbers, float64(i))
//		}
//	}
//	for j := 1; j <= n; j++ {
//		if j%2 != 0 {
//			numbers = append(numbers, float64(j))
//		}
//	}
//	fmt.Println(numbers)
//}

// 21
func main() {
	var n int
	fmt.Println("Введите размер массива: ")
	fmt.Scan(&n)
	var k int
	fmt.Println("Введите первый элемент: ")
	fmt.Scan(&k)
	var l int
	fmt.Println("Введите втарой элемент: ")
	fmt.Scan(&l)
	numbers := []int{}
	for i := 1; i <= n; i++ {
		numbers = append(numbers, i)
	}
	fmt.Println(numbers)
	count := 0
	sum := 0
	//for _, value := range numbers {
	//	if k >= value && l <= value {
	//		count++
	//		sum += value
	//	}
	//}
	//fmt.Println(sum / count)
	for j := k; j <= l; j++ {
		count++
		sum += j
	}
	average := float64(sum) / float64(count)
	fmt.Println(average)
}
