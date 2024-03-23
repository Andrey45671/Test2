package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var a int                                             // длина строки введенной пользователем
	var flag1, flag2, flag3, flag4 int                    // переменные операции введеной пользователем
	var operatorpos int                                   // положение оператора
	var arabicnum1, romannum1, arabicnum2, romannum2 bool //определяем систему счисления
	var num1string, num2string string
	var num1, num2 int
	var errorflag bool
	var qtyflag int
	var check2 bool

	for {
		fmt.Println("Введите значение")
		text, _ := reader.ReadString('\n') //Ждет ввода данных в формате строки
		text = strings.TrimSpace(text)     //Очищает все пустоты (пробелы, табуляцию)

		a = len(text) //определяем длину строки
		codeArr := make([]int, a, a)
		stringArr := make([]string, a, a)

		for i := 0; i < len(text); i++ { //создаем два массива с кодами и значениями
			codeArr[i] = int(text[i])
			stringArr[i] = string(codeArr[i])
		}

		//ищем оператор операции и взводим флаг операции 1 - сложение, 2 - вычитание, 3 - умножение, 4 - деление, проверяем что он один
		for i := 0; i < a; i++ {
			if stringArr[i] == "+" {
				flag1 = 1
				_ = flag1
				operatorpos = i
				_ = operatorpos
				qtyflag = qtyflag + 1
			}
			if stringArr[i] == "-" {
				flag2 = 1
				_ = flag2
				operatorpos = i
				_ = operatorpos
				qtyflag = qtyflag + 1
			}
			if stringArr[i] == "*" {
				flag3 = 1
				_ = flag3
				operatorpos = i
				_ = operatorpos
				qtyflag = qtyflag + 1
			}
			if stringArr[i] == "/" {
				flag4 = 1
				_ = flag4
				operatorpos = i
				_ = operatorpos
				qtyflag = qtyflag + 1
			}
		}
		if qtyflag <= 0 || qtyflag > 1 {
			errorflag = true
		}
		//обрабатываем левую часть вырвжениея
		for i := 0; i < operatorpos; i++ {
			if codeArr[i] >= 48 && codeArr[i] <= 57 {
				arabicnum1 = true
				_ = arabicnum1
				num1string = num1string + stringArr[i]
			}
			if codeArr[i] == 73 || codeArr[i] == 86 || codeArr[i] == 88 {
				romannum1 = true
				_ = romannum1
				num1string = num1string + stringArr[i]
			}
			if arabicnum1 == true && romannum1 == true {
				errorflag = true
			}
			errorflag = checksimvol(codeArr[i])

		}
		if romannum1 == true {
			num1string = romantoarabic(num1string)
		}
		num1, _ = strconv.Atoi(num1string)
		//fmt.Println(num1)
		if num1 == 0 {
			check2 = true
		}
		_ = check2
		if num1 <= 0 || num1 > 10 {
			errorflag = true
		}
		//обрабатываем правую часть вырвжениея
		for i := operatorpos + 1; i < a; i++ {
			if codeArr[i] >= 48 && codeArr[i] <= 57 {
				arabicnum2 = true
				_ = arabicnum2
				num2string = num2string + stringArr[i]
			}
			if codeArr[i] == 73 || codeArr[i] == 86 || codeArr[i] == 88 {
				romannum2 = true
				_ = romannum2
				num2string = num2string + stringArr[i]
			}
			if arabicnum2 == true && romannum2 == true {
				errorflag = true
			}
			errorflag = checksimvol(codeArr[i])
		}
		if romannum2 == true {
			num2string = romantoarabic(num2string)
		}
		if (arabicnum1 == true && romannum2 == true) || (arabicnum2 == true && romannum1 == true) {
			errorflag = true
		}
		num2, _ = strconv.Atoi(num2string)
		//fmt.Println(num2)
		if num2 == 0 {
			check2 = true
		}
		if num2 <= 0 || num2 > 10 {
			errorflag = true
		}
		//вызываем функцию калькулятора
		if qtyflag <= 0 || qtyflag > 1 {
			errorflag = true
		}
		if check2 == true {
			errorflag = true
		}
		if errorflag == false {
			if flag1 == 1 {
				addition(num1, num2, romannum1)
			}
			if flag2 == 1 {
				subtraction(num1, num2, romannum1)
			}
			if flag3 == 1 {
				multiplication(num1, num2, romannum1)

			}
			if flag4 == 1 {
				divisions(num1, num2, romannum1)

			}
		}
		if errorflag == true {
			makePanic()
		}
		flag1, flag2, flag3, flag4 = 0, 0, 0, 0
		num1string, num2string = "", ""
		errorflag = false
		arabicnum1, romannum1, arabicnum2, romannum2 = false, false, false, false
		qtyflag = 0

	}

}

func addition(fnum1, fnum2 int, sistem bool) {

	if sistem == false {
		fmt.Println(fnum1 + fnum2)
	}
	if sistem == true {
		fmt.Println(arabictoroman(fnum1 + fnum2))
	}

}

func subtraction(fnum1, fnum2 int, sistem bool) {
	if sistem == false {
		fmt.Println(fnum1 - fnum2)
	}
	if sistem == true && (fnum1-fnum2) > 0 {
		fmt.Println(arabictoroman(fnum1 - fnum2))
	}
	if sistem == true && (fnum1-fnum2) < 0 {
		makePanic()
	}
}

func divisions(fnum1, fnum2 int, sistem bool) {
	if sistem == false {
		fmt.Println(fnum1 / fnum2)
	}
	if sistem == true {
		fmt.Println(arabictoroman(fnum1 / fnum2))
	}
}

func multiplication(fnum1, fnum2 int, sistem bool) {
	if sistem == false {
		fmt.Println(fnum1 * fnum2)
	}
	if sistem == true {
		fmt.Println(arabictoroman(fnum1 * fnum2))
	}
}

func checksimvol(fsimvol int) bool {
	var lenArr int
	var result bool
	result = true
	checkArr := []int{48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 73, 86, 88}
	_ = checkArr
	lenArr = len(checkArr)
	for i := 0; i < lenArr; i++ {
		if checkArr[i] == fsimvol {
			result = false
		}
	}
	return result
}

func romantoarabic(froman string) string {
	var result string
	switch froman {
	case "I":
		result = "1"
	case "II":
		result = "2"
	case "III":
		result = "3"
	case "IV":
		result = "4"
	case "V":
		result = "5"
	case "VI":
		result = "6"
	case "VII":
		result = "7"
	case "VIII":
		result = "8"
	case "IX":
		result = "9"
	case "X":
		result = "10"
	}
	return result
}

func arabictoroman(farabic int) string {
	var result string
	var lenArr int
	spisok1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 20, 30, 40, 50, 60, 70, 80, 90, 100, 500, 1000}
	_ = spisok1
	spisok2 := []string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC", "C", "D", "M"}
	_ = spisok2
	lenArr = len(spisok1)
	_ = lenArr
	for i := lenArr - 1; i > 0; i-- {
		if farabic > 0 && farabic >= spisok1[i] {
			result = result + spisok2[i]
			farabic = farabic - spisok1[i]

		}
	}

	return result
}

func makePanic() {
	defer func() {
		panicValue := recover()
		fmt.Println(panicValue)
	}()
	panic("Паника")
}
