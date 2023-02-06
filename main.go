package main

import (
	"bufio"
	"errors"
	"fmt"
)

import (
	"os"
	"strconv"
	"strings"
)

var RomanToArabicMap = map[string]int{
	"I":    1,
	"II":   2,
	"III":  3,
	"IV":   4,
	"V":    5,
	"VI":   6,
	"VII":  7,
	"VIII": 8,
	"IX":   9,
	"X":    10,
}

func IsArabic(str string) bool {
	buffer, _ := strconv.Atoi(str)
	if buffer <= 10 {
		return true
	} else {
		return false
	}
}
func IsRoman(str string) bool {
	_, ok := RomanToArabicMap[str]
	if ok {
		return ok
	} else {
		return false
	}
}
func UserInput() ([]string, error) {
	reader := bufio.NewReader(os.Stdin)

	buffer, _ := reader.ReadString('\n')
	buffer = strings.TrimSpace(buffer)
	s := strings.Split(buffer, " ")

	if len(s) < 3 || len(s) > 3 {
		err := errors.New("UserInput: неверное математическое выражение")
		return nil, err
	} else {
		return s, nil
	}
}
func GetActionSign(str []string) (string, error) {
	switch string(str[1]) {
	case "+":
		return "+", nil
	case "-":
		return "-", nil
	case "*":
		return "*", nil
	case "/":
		return "/", nil
	default:
		NotCorrectSign := errors.New("GetActionSign: неверный знак")
		return "", NotCorrectSign
	}
}
func RomanCalc(str []string, sign string) (string, error) {
	firstNumber := str[0]
	secondNumber := str[2]
	buffer := 0
	switch sign {
	case "+":
		buffer = RomanToArabicMap[firstNumber] + RomanToArabicMap[secondNumber]
	case "-":
		buffer = RomanToArabicMap[firstNumber] - RomanToArabicMap[secondNumber]
	case "*":
		buffer = RomanToArabicMap[firstNumber] * RomanToArabicMap[secondNumber]
	case "/":
		buffer = RomanToArabicMap[firstNumber] / RomanToArabicMap[secondNumber]
	}

	if buffer < 1 {
		err := errors.New("RomanCalc: результат вычисления меньше 1")
		return "", err
	}
	return TranslateArabicToRoman(buffer), nil
}
func ArabicCalc(str []string, sign string) (int, error) {
	firstNumber, err := strconv.Atoi(str[0])
	if err != nil {
		return 0, fmt.Errorf("ArabicCalc: введеный символ не число / попытка операции в разных системах счисления")
	}
	secondNumber, err := strconv.Atoi(str[2])
	if err != nil {
		return 0, fmt.Errorf("ArabicCalc: введеный символ не число / попытка операции в разных системах счисления")
	}
	buffer := 0

	switch sign {
	case "+":
		buffer = int(firstNumber + secondNumber)
	case "-":
		buffer = int(firstNumber - secondNumber)
	case "*":
		buffer = int(firstNumber * secondNumber)
	case "/":
		buffer = int(firstNumber / secondNumber)
	default:
		return -1, nil
	}
	return buffer, err
}
func TranslateArabicToRoman(arabicNum int) string {
	arabicNums := []int{100, 90, 50, 40, 10, 9, 5, 4, 1}
	romanChar := []string{"C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	buffer := ""
	for i, arabicNumsPointer := range arabicNums {
		buffer += strings.Repeat(romanChar[i], arabicNum/arabicNumsPointer)
		arabicNum %= arabicNumsPointer
	}
	return buffer
}

func main() {
	for {
		fmt.Print("Input: ")
		input, err := UserInput()
		if err != nil {
			fmt.Println(err)
			break
		} else {
			sign, err := GetActionSign(input)
			if err != nil {
				fmt.Println(err)
				break
			} else {
				if IsRoman(input[0]) && IsRoman(input[2]) {
					result, err := RomanCalc(input, sign)
					if err != nil {
						fmt.Println(err)
						break
					} else {
						fmt.Println("Output:", result)
					}
				} else if IsArabic(input[0]) && IsArabic(input[2]) {
					result, err := ArabicCalc(input, sign)
					if err != nil {
						fmt.Println(err)
						break
					} else {
						fmt.Println("Output:", result)
					}
				} else {
					fmt.Println("Введеное число более 10")
				}
			}
		}
	}
	fmt.Println("Калькулятор завершил работу")
}
