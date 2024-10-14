package main

import (
	"bufio"
	"fmt"
	"math"
	"math/rand"
	"os"
	"strconv"
	"time"
)

// Проверка на простоту
// Напишите функцию, которая проверяет, является ли переданное число простым.
// Ваша программа должна использовать циклы для проверки делителей, и если
// число не является простым, выводить первый найденный делитель.
func task1() {
	fmt.Println("TASK 1")

	var num int
	_, err := fmt.Scanf("%d", &num)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	res := 0

	for i := 2; i <= int(math.Sqrt(float64(num))); i++ {
		if num%i == 0 {
			res = i
		}
	}

	if res == 0 {
		fmt.Println("RESULT:", "Простое")
	} else {
		fmt.Println("RESULT:", res)
	}
}

// Наибольший общий делитель (НОД)
// Напишите программу для нахождения наибольшего общего делителя (НОД) двух
// чисел с использованием алгоритма Евклида. Используйте цикл `for` для
// вычислений.
func task2() {
	fmt.Println("TASK 2")

	var num1 int
	_, err := fmt.Scanf("%d", &num1)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	var num2 int
	_, err = fmt.Scanf("%d", &num2)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	res := 0

	for {
		if num2 == 0 {
			res = num1
			break
		}
		temp := num1
		num1 = num2
		num2 = temp % num2
	}

	fmt.Println("RESULT:", res)
}

// Сортировка пузырьком
// Реализуйте сортировку пузырьком для списка целых чисел. Программа должна
// выполнять сортировку на месте и выводить каждый шаг изменения массива.
func task3() {
	fmt.Println("TASK 3")

	arr := []int{7, 6, 5, 10, -12, 1, -24, 66, 1, 0, -5, 76, 666, -1231, 14321}
	fmt.Println("arr:", arr)

	for i := 0; i < len(arr); i++ {
		for j := 1; j < len(arr); j++ {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			}
		}
	}

	fmt.Println("RESULT:", arr)
}

// Таблица умножения в формате матрицы
// Напишите программу, которая выводит таблицу умножения в формате матрицы 10x10.
// Используйте циклы для генерации строк и столбцов.
func task4() {
	fmt.Println("TASK 4")

	for i := 0; i < 10; i++ {
		fmt.Printf("\t%d", i)
		for j := 1; j < 10; j++ {
			if i == 0 {
				fmt.Printf("\t%d", j)
			} else {
				fmt.Printf("\t%d", i*j)
			}
		}
		fmt.Println()
	}
}

func memfib(num int) int {

	memoized := make([]int, num+1, num+2)
	if num < 2 {
		memoized = memoized[0:2]
	}

	memoized[0] = 0
	memoized[1] = 1

	var initial int = 0
	for initial = 2; initial <= num; initial++ {
		memoized[initial] = memoized[initial-1] + memoized[initial-2]
	}

	return memoized[num]
}

// Фибоначчи с мемоизацией
// Напишите функцию для вычисления числа Фибоначчи с использованием мемоизации
// (сохранение ранее вычисленных результатов). Программа должна использовать
// рекурсию и условные операторы.
func task5() {
	fmt.Println("TASK 5")

	var num int
	_, err := fmt.Scanf("%d", &num)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	res := ""
	for i := 0; i < num; i++ {
		res += strconv.Itoa(memfib(i)) + " "
	}

	fmt.Println("RESULT:", res)

}

// Обратные числа
// Напишите программу, которая принимает целое число и выводит его в обратном
// порядке. Например, для числа 12345 программа должна вывести 54321. Используйте
// цикл для обработки цифр числа.
func task6() {
	fmt.Println("TASK 6")

	var num int
	_, err := fmt.Scanf("%d", &num)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Print("RESULT: ")
	for num != 0 {
		fmt.Print(num % 10)
		num /= 10
	}

	fmt.Println()
}

func helper(index int, temp []int) []int {
	arr := make([]int, 0)
	arr = append(arr, 1)
	for i := 1; i < index; i++ {
		arr = append(arr, temp[i]+temp[i-1])
	}
	arr = append(arr, 1)
	return arr
}

func generate(numRows int) [][]int {
	result := make([][]int, 0)
	pascalRow := make([]int, 0)
	pascalRow = append(pascalRow, 1)
	result = append(result, pascalRow)
	for i := 1; i < numRows; i++ {
		pascalRow = helper(i, pascalRow)
		result = append(result, pascalRow)
	}
	return result
}

// Треугольник Паскаля
// Напишите программу, которая выводит треугольник Паскаля до заданного уровня.
// Для этого используйте цикл и массивы для хранения предыдущих значений строки
// треугольника.
func task7() {
	fmt.Println("TASK 7")

	var num int
	_, err := fmt.Scanf("%d", &num)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	res := generate(num)
	for i := 0; i < len(res); i++ {
		fmt.Println(res[i])
	}
}

// Число палиндром
// Напишите программу, которая проверяет, является ли число палиндромом (одинаково
// читается слева направо и справа налево). Не используйте строки для решения этой
// задачи — работайте только с числами.
func task8() {
	fmt.Println("TASK 8")

	var num int
	_, err := fmt.Scanf("%d", &num)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	dels := []int{}
	for num != 0 {
		dels = append(dels, num%10)
		num /= 10
	}

	isp := true
	for i := 0; i < len(dels)/2; i++ {
		if dels[i] != dels[len(dels)-i-1] {
			isp = false
			break
		}
	}

	if isp == true {
		fmt.Println("RESULT: Палиндром")
	} else {
		fmt.Println("RESULT: Не палиндром")
	}
}

// Нахождение максимума и минимума в массиве
// Напишите функцию, которая принимает массив целых чисел и возвращает одновременно
// максимальный и минимальный элемент с использованием одного прохода по массиву.
func task9() {
	fmt.Println("TASK 9")

	arr := []int{7, 6, 5, 10, -12, 1, -24, 66, 1, 0, -5, 76, 666, -1231, 14321}
	fmt.Println("arr:", arr)

	max := arr[0]
	min := arr[0]
	for _, i := range arr {
		if i > max {
			max = i
		}

		if i < min {
			min = i
		}
	}

	fmt.Println("RESULT:", "MIN = ", min, " | MAX = ", max)
}

// Игра "Угадай число"
// Напишите программу, которая загадывает случайное число от 1 до 100, а пользователь
// пытается его угадать. Программа должна давать подсказки "больше" или "меньше" после
// каждой попытки. Реализуйте ограничение на количество попыток.
func task10() {
	fmt.Println("TASK 10")

	attempts := 3

	rand := 1 + rand.Intn(100-1)

	for attempts > 0 {
		var num int
		_, err := fmt.Scanf("%d", &num)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		if num == rand {
			fmt.Println("You win!")
			break
		} else if num > rand {
			fmt.Println("Less")
		} else if num < rand {
			fmt.Println("More")
		}

		attempts--
	}

	if attempts == 0 {
		fmt.Println("You lose!")
	}
}

// Числа Армстронга
// Напишите программу, которая проверяет, является ли число числом Армстронга (число
// равно сумме своих цифр, возведённых в степень, равную количеству цифр числа).
// Например, 153 = 1³ + 5³ + 3³.
func task11() {
	fmt.Println("TASK 11")

	var num int
	_, err := fmt.Scanf("%d", &num)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	len := len(strconv.Itoa(num))

	save := num

	sum := 0
	for num != 0 {
		sum += int(math.Pow(float64(num%10), float64(len)))
		num /= 10
	}

	if sum == save {
		fmt.Println("RESULT: Да")
	} else {
		fmt.Println("RESULT: Нет")

	}
}

// Подсчет слов в строке
// Напишите программу, которая принимает строку и выводит количество уникальных
// слов в ней. Используйте `map` для хранения слов и их количества.
func task12() {
	fmt.Println("TASK 12")

	in := bufio.NewReader(os.Stdin)
	line, err := in.ReadString('\n')
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	uniq := make(map[string]bool)
	count := 0
	ppos := 0
	for i := 0; i < len(line); i++ {
		if line[i] == ' ' || line[i] == '\n' {
			if !uniq[line[ppos:i]] {
				uniq[line[ppos:i]] = true
			}
			ppos = i + 1
			count++
		}
	}

	fmt.Println("RESULT:", count)
	for k, _ := range uniq {
		fmt.Print(k, " ")
	}
	fmt.Println()
}

// Игра "Жизнь" (Conway's Game of Life)
// Реализуйте клеточный автомат "Жизнь" Конвея для двухмерного массива. Каждая клетка
// может быть либо живой, либо мертвой. На каждом шаге состояния клеток изменяются
// по следующим правилам:
//   - Живая клетка с двумя или тремя живыми соседями остаётся живой, иначе умирает.
//   - Мёртвая клетка с тремя живыми соседями оживает.
//
// Используйте циклы для обработки клеток.
const (
	width          = 10
	height         = 10
	sleepIteration = 500
	ansiEscapeSeq  = "\033c\x0c"
	brownSquare    = "\xF0\x9F\x9F\xAB"
	greenSquare    = "\xF0\x9F\x9F\xA9"
)

type World [][]bool

func (w World) Display() {
	for _, row := range w {
		for _, cell := range row {
			switch {
			case cell:
				fmt.Printf(greenSquare)
			default:
				fmt.Printf(brownSquare)
			}
		}
		fmt.Printf("\n")
	}
}

func (w World) Seed() {
	for _, row := range w {
		for i := range row {
			if rand.Intn(4) == 1 {
				row[i] = true
			}
		}
	}
}

func (w World) Alive(x, y int) bool {
	y = (height + y) % height
	x = (width + x) % width
	return w[y][x]
}

func (w World) Neighbors(x, y int) int {
	var neighbors int

	for i := y - 1; i <= y+1; i++ {
		for j := x - 1; j <= x+1; j++ {
			if i == y && j == x {
				continue
			}
			if w.Alive(j, i) {
				neighbors++
			}
		}
	}
	return neighbors
}

func (w World) Next(x, y int) bool {
	n := w.Neighbors(x, y)
	alive := w.Alive(x, y)
	if n < 4 && n > 1 && alive {
		return true
	} else if n == 3 && !alive {
		return true
	} else {
		return false
	}
}

func Step(a, b World) {
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			b[i][j] = a.Next(j, i)
		}
	}
}

func MakeWorld() World {
	w := make(World, height)
	for i := range w {
		w[i] = make([]bool, width)
	}
	return w
}

func task13() {
	fmt.Println("TASK 13")

	fmt.Println(ansiEscapeSeq)
	newWorld := MakeWorld()
	nextWorld := MakeWorld()
	newWorld.Seed()
	for {
		newWorld.Display()
		Step(newWorld, nextWorld)
		newWorld, nextWorld = nextWorld, newWorld
		time.Sleep(sleepIteration * time.Millisecond)
		fmt.Println(ansiEscapeSeq)
	}
}

// Цифровой корень числа
// Напишите программу, которая вычисляет цифровой корень числа. Цифровой корень — это
// рекурсивная сумма цифр числа, пока не останется только одна цифра. Например, цифровой
// корень числа 9875 равен 2, потому что 9+8+7+5=29 → 2+9=11 → 1+1=2.
func task14() {
	fmt.Println("TASK 14")

	var num int
	_, err := fmt.Scanf("%d", &num)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for len(strconv.Itoa(num)) != 1 {
		temp := num
		sum := 0
		for temp != 0 {
			sum += temp % 10
			temp /= 10
		}
		num = sum
	}

	fmt.Println("RESULT:", num)
}

var values = []struct {
	decVal int
	symbol string
}{
	{1000, "M"}, {900, "CM"}, {500, "D"}, {400, "CD"},
	{100, "C"}, {90, "XC"}, {50, "L"}, {40, "XL"},
	{10, "X"}, {9, "IX"}, {5, "V"}, {4, "IV"}, {1, "I"},
}

// Римские цифры
// Напишите функцию, которая преобразует арабское число (например, 1994) в римское
// (например, "MCMXCIV"). Программа должна использовать циклы и условные операторы для
// создания римской записи.
func task15() {
	fmt.Println("TASK 15")

	var num int
	_, err := fmt.Scanf("%d", &num)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	res := ""
	for num != 0 {
		for _, pair := range values {
			if num >= pair.decVal {
				res += pair.symbol
				num = num - pair.decVal
			}
		}
	}

	fmt.Println("RESULT:", res)
}

func main() {
	task1()
	task2()
	task3()
	task4()
	task5()
	task6()
	task7()
	task8()
	task9()
	task10()
	task11()
	task12()
	task13()
	task14()
	task15()
}
