package main

import (
	"fmt"
	"strings"
)

func iteration() {
	// for -- only way to loop
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// while-style
	k := 3
	for k > 0 {
		fmt.Println(k)
		k--
	}

	//infinite loop
	//for {}

	items := [3]string{"go", "python", "java"}
	for index, value := range items {
		fmt.Println(index, value)
	}
}

func condition() {
	temp := 25
	if temp == 30 {
		fmt.Println(temp)
	} else if temp > 30 {
		fmt.Println("greater than 30")
	} else {
		fmt.Println("less than 30")
	}

	userAccess := map[string]bool{
		"jane": true,
		"jhon": false,
	}

	if hasAccess, ok := userAccess["jane"]; ok && hasAccess {
		fmt.Println("She can acceess")
	}
}

func switchCondition() {
	day := "Sunday"
	switch day {
	case "Monday", "Tuesday", "Wednesday", "Thursday":
		fmt.Println("It's a weekday")
	case "Saturday", "Sunday":
		fmt.Println("It's a weekend")
	default:
		fmt.Println("It's a Friday")
	}

	checkType := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Printf("I'm a bool: %t\n", t)
		case int:
			fmt.Printf("I'm an int: %d\n", t)
		case string:
			fmt.Printf("I'm a string %s\n", t)
		default:
			fmt.Printf("Unknown type: %T\n", t)
		}
	}
	checkType(21)
	checkType(true)
	checkType(1.5)
	checkType("apple")
}

var productPrices = map[string]float64{
	"T-SHIRT": 20.00,
	"MUG":     12.50,
	"HAT":     18.00,
	"BOOK":    25.50,
	"TEST":    0.00,
}

func calculateItemPrice(itemCode string) (float64, bool) {
	basePrice, found := productPrices[itemCode]
	if !found {
		if before, ok := strings.CutSuffix(itemCode, "_SALE"); ok {
			originalItemCode := before
			basePrice, found = productPrices[originalItemCode]
			if found {
				salePrice := basePrice * 0.90
				fmt.Printf("Item %s (Sale! Original: $%.2f, Sale price: %.2f",
					originalItemCode, basePrice, salePrice)
				return salePrice, true
			}
		}
		fmt.Printf("Item: %s (Product not found)\n)", itemCode)
		return 0.0, false
	}
	return basePrice, true
}

func main() {
	iteration()
	condition()
	switchCondition()

	orderItems := []string{
		"HAT_SALE", "T-SHIRT", "BOOK", "MUG", "COFFEE",
	}

	for _, orderItem := range orderItems {
		fmt.Println(calculateItemPrice(orderItem))
	}
}
