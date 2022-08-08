package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Hello! I can solve your math problem ;)")
	fmt.Print("Please input your math problem and press Enter: ")
	reader := bufio.NewReader(os.Stdin)
	// ReadString will block until the delimiter is entered
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("An error occured while reading input. Please try again", err)
		return
	}

	var a, b, res float64
	var oper string
	a, oper, b = Scanner1(input)

	switch oper {
	case "+":
		res = a + b
	case "-":
		res = a - b
	case "*":
		res = a * b
	case "/":
		res = a / b
	}

	fmt.Println(res)

}

//The Scanner funcion reads an argument from the command line
//as a strinfg and splits it into two arguments and an operator
func Scanner1(s string) (float64, string, float64) {

	operFlag := false
	arg1Str := ""
	arg2Str := ""
	var arg1, arg2 float64
	oper := ""
	err3 := true
	for _, el := range s {
		if (el >= '0' && el <= '9') && !operFlag {
			arg1Str += string(el)
		} else if MathOper(el) {
			oper += string(el)
			err3 = false
			operFlag = true
		} else if (el >= '0' && el <= '9') && operFlag {
			arg2Str += string(el)
		}
	}
	if err3 {
		fmt.Println("Please input correct math operator such as: + - * / ")
		return 0, "Error", 0
	}

	arg1, err1 := strconv.ParseFloat(arg1Str, 64)
	arg2, err2 := strconv.ParseFloat(arg2Str, 64)
	if err1 != nil || err2 != nil {
		fmt.Println(err1)
		fmt.Println("Error: Please input valid number")
		return 0, "Error", 0
	}
	return float64(arg1), oper, float64(arg2)
}

//MathOper function checks if a given rune is a math operator
func MathOper(r rune) bool {
	operArr := "+-*/"
	for _, el := range operArr {
		if el == r {
			return true
		}
	}
	return false
}