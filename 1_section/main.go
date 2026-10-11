package main

import "fmt"

const HOST string = "localhost"

const (
	Sunday    = 0
	Monday    = 1
	Tuesday   = 2
	Wednesday = 3
	Thursday  = 4
	Friday    = 5
	Saturday  = 6
)

type LogLevel int

const (
	LogError LogLevel = 0
	LogWarn  LogLevel = 1
	LogInfo  LogLevel = 2
	LogDebug LogLevel = 3
)

var levelNames = []string{"Error", "Warn", "Info", "Debug"}

func (l LogLevel) String() string {
	if l < LogError || l > LogWarn {
		return "LogError"
	}
	return levelNames[l]
}

func printLogLevel(level LogLevel) {
	fmt.Printf("Log Level: %d %s\n", level, level.String())
}

func main() {
	fmt.Println("Hello " + "World\n")
	fmt.Println(1 + 1)
	fmt.Println(true, false)
	fmt.Println([]int{1, 2, 3})

	// variables
	var greeting string // zero-value is an empty string
	greeting = "Hello"
	fmt.Println(greeting)

	var count int
	count = 10
	fmt.Println(count)

	var isRunning bool
	isRunning = true
	fmt.Println(isRunning)

	var firstname, lastname string
	firstname = "Doe"
	lastname = "Joh"
	fmt.Println(firstname, lastname)

	email := "sample@sample.com" // infer type when this code is compiled
	fmt.Println(email)

	var year = 2025
	fmt.Println(year)

	fmt.Println(HOST)

	// enum in go
	fmt.Println(Sunday)
	fmt.Println(Monday)
	fmt.Println(Tuesday)
	fmt.Println(Wednesday)
	fmt.Println(Thursday)
	fmt.Println(Friday)
	fmt.Println(Saturday)

	fmt.Println(LogError)
	fmt.Println(LogWarn)
	fmt.Println(LogInfo)
	fmt.Println(LogDebug)

	fmt.Println("-----------\n")

	printLogLevel(LogError)
	printLogLevel(LogDebug)
	printLogLevel(LogInfo)
	printLogLevel(LogWarn)
}
