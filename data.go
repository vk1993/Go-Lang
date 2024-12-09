package main

import (
	"errors"
	"fmt"
)

const secondInput = 200

func main() {
	var stringResult, intysy, err3 = errorExample("error_Test")
	var i, i_1 int

	if err3 != nil {
		fmt.Println(err3.Error())
		fmt.Println(stringResult)
		fmt.Println(intysy)
	}

	fmt.Println("Let's start 1st sum function")
	fmt.Print("Type 1st Number : ")
	fmt.Scan(&i)
	fmt.Print("Type 2nd Number : ")
	fmt.Scan(&i_1)
	result, err := add_two_number(i, i_1)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Sum of above : ", result)

	result_with_const, err_2 := add_two_number_with_const(i)

	if err_2 != nil {
		fmt.Println(err_2)
	}
	fmt.Printf("Sum of above : %v", result_with_const)
	// handleRequests()
}

func add_two_number(a int, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("intput b is 0 error")
	}
	return a + b, nil
}

func add_two_number_with_const(input int) (int, error) {
	if input == 0 {
		return 0, errors.New("entered intput is 0 error")
	}

	return input + secondInput, nil
}

func init() {
	fmt.Println("sdfsdf")
}

func foo(format int, name string) {
	switch format {
	case 0:
		fmt.Printf("anta hi arambh he %v", name)
	case 1:
		println("subha se sham tak")
	default:
		println("go out of india")
	}
}

func errorExample(inputString string) (string, int, error) {
	fmt.Println(inputString)
	var err error = errors.New(inputString)
	return "datat", 0, err
}

func init() {
	var n int = 0
	var data = "ggshggsh"
	for i := 0; i < 10; i++ {
		fmt.Println(n)
		fmt.Println(data)
	}

	// var stroingdata string =
	// fmt.Print(stroingdata);

	var1, var3 := "visal", 2
	fmt.Println(var1, var3)
	foo(0, "string")

	if n != 0 {
		foo(3, "visal")
	} else if n == 1 {
		foo(4, "indea")
	}
}
