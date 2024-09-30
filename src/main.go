package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func task1() {
	fmt.Println("TASK 1")

	var num int

	_, err := fmt.Scanf("%d", &num)
	if err != nil { // подобные условия, полагаю, не будут браться в расчёт, т.к. они "не функциональные"
		fmt.Println(err.Error())
		return
	}

	// int max = 2147483647 (10 digits)
	n0 := num % 10
	num /= 10
	n1 := num % 10
	num /= 10
	n2 := num % 10
	num /= 10
	n3 := num % 10
	num /= 10
	n4 := num % 10
	num /= 10
	n5 := num % 10
	num /= 10
	n6 := num % 10
	num /= 10
	n7 := num % 10
	num /= 10
	n8 := num % 10
	num /= 10
	n9 := num % 10
	num /= 10

	sum := n0 + n1 + n2 + n3 + n4 + n5 + n6 + n7 + n8 + n9

	fmt.Println("RESULT:", sum)
}

func task2() {
	fmt.Println("TASK 2")

	var inp string

	_, err := fmt.Scanf("%s", &inp)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var res float64

	t := inp[len(inp)-1]
	n, err := strconv.ParseFloat(inp[:len(inp)-1], 64)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if t == 'F' || t == 'f' { // эти условия сделаны для удобства ввода. Логика преобразования безусловная
		res = (n - 32) * 5 / 9
	} else if t == 'C' || t == 'c' {
		res = (n * 9 / 5) + 32
	} else {
		fmt.Println("Unknown type")
		return
	}

	fmt.Println("RESULT:", res)
}

func task3() {
	fmt.Println("TASK 3")

	var inp string
	_, err := fmt.Scanf("%s", &inp)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	nums := strings.Split(inp[1:len(inp)-1], ",")
	i := 0
HACK:
	num, _ := strconv.Atoi(nums[i])
	nums[i] = strconv.Itoa(num * 2)
	i++
	if i < len(nums) {
		goto HACK
	}

	fmt.Println("RESULT:", nums)
}

func task4() {
	fmt.Println("TASK 4")

	var inp string
	_, err := fmt.Scanf("%s", &inp)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	strs := strings.Split(inp[1:len(inp)-1], ",")

	var res string
	i := 0
HACK:
	if i < len(strs) {
		strs[i] = strings.ReplaceAll(strs[i], "\"", "")
		res += " " + strs[i]
		i++
		goto HACK
	}

	fmt.Println("RESULT:", res)
}

func task5() {
	fmt.Println("TASK 5")

	var x1, x2, y1, y2 float64
	fmt.Scanf("%f", &x1)
	fmt.Scanf("%f", &y1)
	fmt.Scanf("%f", &x2)
	fmt.Scanf("%f", &y2)

	res := math.Sqrt((x2-x1)*(x2-x1) + (y2-y1)*(y2-y1))

	fmt.Println("RESULT:", res)
}

func task6() {
	fmt.Println("TASK 6")

	var num int
	fmt.Scanf("%d", &num)

	var res string
	if num%2 == 0 {
		res = "Четное"
	} else {
		res = "Нечетное"

	}

	fmt.Println("RESULT:", res)
}

func task7() {
	fmt.Println("TASK 7")

	var year int
	fmt.Scanf("%d", &year)

	var res string
	if year%4 == 0 && year%100 != 0 || year%400 == 0 {
		res = "Високосный"
	} else {
		res = "Не високосный"

	}

	fmt.Println("RESULT:", res)
}

func task8() {
	fmt.Println("TASK 8")

	var x1, x2, x3 float64
	fmt.Scanf("%f", &x1)
	fmt.Scanf("%f", &x2)
	fmt.Scanf("%f", &x3)

	res := math.Max(math.Max(x1, x2), x3)

	fmt.Println("RESULT:", res)
}

func task9() {
	fmt.Println("TASK 9")

	var age int
	fmt.Scanf("%d", &age)

	// Ребенок <12
	// Подросток <18
	// Взрослый <70
	// Пожилой >=70
	var res string
	if age < 12 {
		res = "Ребенок"
	} else if age < 18 {
		res = "Подросток"
	} else if age < 70 {
		res = "Взрослый"
	} else {
		res = "Пожилой"
	}

	fmt.Println("RESULT:", res)
}

func task10() {
	fmt.Println("TASK 10")

	var num int
	fmt.Scanf("%d", &num)

	var res string
	if num%3 == 0 && num%5 == 0 {
		res = "Делится"
	} else {
		res = "Не делится"

	}

	fmt.Println("RESULT:", res)
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func task11() {
	fmt.Println("TASK 11")

	var num int
	fmt.Scanf("%d", &num)
	res := factorial(num)

	fmt.Println("RESULT:", res)
}

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	var n2, n1 = 0, 1
	for i := 2; i <= n; i++ {
		n2, n1 = n1, n1+n2
	}
	print(n1)
	return n1
}

func task12() {
	fmt.Println("TASK 12")

	var num int
	fmt.Scanf("%d", &num)

	res := ""
	for i := 0; i < num; i++ {
		res += strconv.Itoa(fibonacci(i)) + " "
	}

	fmt.Println("RESULT:", res)
}

func task13() {
	fmt.Println("TASK 13")

	var inp string
	_, err := fmt.Scanf("%s", &inp)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	nums := strings.Split(inp[1:len(inp)-1], ",")

	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}

	fmt.Println("RESULT:", nums)
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func task14() {
	fmt.Println("TASK 14")

	var num int
	fmt.Scanf("%d", &num)

	res := ""
	for i := 0; i < num; i++ {
		if isPrime(i) {
			res += strconv.Itoa(i) + " "
		}
	}

	fmt.Println("RESULT:", res)
}

func task15() {
	fmt.Println("TASK 15")

	var inp string
	_, err := fmt.Scanf("%s", &inp)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	nums := strings.Split(inp[1:len(inp)-1], ",")

	var res int
	for i := 0; i < len(nums); i = i + 1 {
		num, _ := strconv.Atoi(nums[i])
		res += num
	}

	fmt.Println("RESULT:", res)
}

func main() {
	// BLOCK 1
	task1()
	task2()
	task3()
	task4()
	task5()

	// BLOCK 2
	task6()
	task7()
	task8()
	task9()
	task10()

	// BLOCK 3
	task11()
	task12()
	task13()
	task14()
	task15()
}
