package main

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

func calculateArea(width, height float64) float64 {
	if width < 0 || height < 0 {
		return 0
	}
	return width * height
}

func factorial(x int) int {
	if x <= 1 {
		return 1
	}
	return x * factorial(x-1)
}

// Closure
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

// Variadic function

func sum(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func splitName(fullName string) (firstname, lastname string) {
	parts := strings.Split(fullName, " ")
	firstname = parts[0]
	lastname = parts[1]
	return firstname, lastname
}

// custom error type
var ErrorDivisionByZero = errors.New("division by zero")
var ErrorNumTooLarge = errors.New("number too large")

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, ErrorDivisionByZero
	}
	if a > 100 {
		return 0, ErrorNumTooLarge
	}
	return a / b, nil
}

type OpError struct {
	Op      string
	Code    int
	Message string
	Time    time.Time
}

func (op OpError) Error() string {
	return op.Message
}

func NewOpError(op string, code int, message string, t time.Time) *OpError {
	return &OpError{op, code, message, t}
}

func doSomething() error {
	return NewOpError("doSomething", 100, "message", time.Now())
}

// defer
func simpleDefer() {
	fmt.Println("Function simpleDefer: Start")
	defer fmt.Println("Function simpleDefer: deferred")
	fmt.Println("Function simpleDefer: Middle")
	fmt.Println("Function simpleDefer: Middle")
	fmt.Println("Function simpleDefer: Middle")
	fmt.Println("Function simpleDefer: Middle")
	fmt.Println("Function simpleDefer: Middle")
}

func lifoSimpleDefer() {
	fmt.Println("Function lifosimpleDefer: Start")
	defer fmt.Println("First: deferred")
	defer fmt.Println("Second : deferred")
	fmt.Println("Function lifosimpleDefer: Middle")
}

// panic and recover
func mightPanic(shouldPanic bool) {
	if shouldPanic {
		panic("something went wrong in might panic")
	}

	fmt.Println("This function executed without panic")
}

func recoverable() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic", r)
		}
	}()
	mightPanic(true)
}

type MathError struct {
	Operation string
	InputA    int
	InputB    int
	Message   string
}

const (
	division       = "Division"
	divisionErrMsg = "division by zero"
)

func (e MathError) Error() string {
	var inputs []string
	if e.Operation == division {
		inputs = append(inputs, fmt.Sprintf("a=%d", e.InputA))
		inputs = append(inputs, fmt.Sprintf("a=%d", e.InputB))
	}
	return fmt.Sprintf("Math error %s %s %s",
		e.Operation, strings.Join(inputs, ","), e.Message)
}

func Sum(numbers ...int) int {
	defer fmt.Println("Sum finished")

	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func safeDivision(a, b int) (int, error) {
	if b == 0 {
		return 0, &MathError{
			Operation: "div",
			InputA:    a,
			InputB:    b,
			Message:   divisionErrMsg,
		}
	}
	return a / b, nil
}

func main() {
	area := calculateArea(10, 5)
	fmt.Println(area)
	result := factorial(5)
	fmt.Println(result)

	nextInt := intSeq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	logger := func(msg string) {
		fmt.Println(msg)
	}
	logger("Hello World")

	fmt.Println(sum(1, 2, 3, 4, 5))

	fmt.Println(divide(10, 5))
	val, err := divide(5, 0)
	if err != nil {
		if errors.Is(err, ErrorDivisionByZero) {
			fmt.Println("Divide by zero")
		} else if errors.Is(err, ErrorNumTooLarge) {
			fmt.Println("number too large")
		} else {
			fmt.Println(val)
		}
	}

	firstname, lastname := splitName("test name")
	println(firstname, lastname)
	simpleDefer()
	lifoSimpleDefer()

	fmt.Println(Sum(1, 2, 3, 4, 5))
	fmt.Println(safeDivision(10, 5))
	fmt.Println(safeDivision(10, 0))
	recoverable()

}
