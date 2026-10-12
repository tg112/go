package main

import (
	"fmt"
	"time"
)

type Employee struct {
	ID        int
	FirstName string
	LastName  string
	Position  string
	Salary    int
	IsActive  bool
	JoinedAt  time.Time
}

func NewEmployee(id int, firstName string, lastName string, position string, salary int, isActive bool) *Employee {
	return &Employee{
		ID:        id,
		FirstName: firstName,
		LastName:  lastName,
		Position:  position,
		Salary:    salary,
		IsActive:  isActive,
	}
}

func (e *Employee) fullName() string {
	return e.FirstName + " " + e.LastName
}

func (e *Employee) deactivate() {
	e.IsActive = false
}

// methods and receiver
func sampleReceiver() {
	joe := NewEmployee(2, "Joe", "Smith", "Server", 20000, true)
	fmt.Println(joe.fullName())
	fmt.Println(joe.IsActive)
	joe.deactivate()
	fmt.Println(joe.IsActive)
}

func structSample() {
	jane := Employee{
		ID:        1,
		FirstName: "Jane",
		LastName:  "Doe",
		Position:  "Night",
		Salary:    10,
		IsActive:  true,
		JoinedAt:  time.Now(),
	}
	fmt.Println("%+v\n", jane)
	fmt.Println(jane.ID)
	fmt.Println(jane.FirstName)

	joe := NewEmployee(2, "Joe", "Smith", "Server", 20000, true)
	fmt.Println(*joe)
}

// Interface
type Person interface {
	GetName() string
	GetId() int
}

type Student struct {
	ID   int
	Name string
}

func (e Student) GetName() string {
	return e.Name
}

func (e Student) GetId() int {
	return e.ID
}

func displayPerson(s Person) {
	fmt.Printf("%d: %s\n", s.GetId(), s.GetName())
}

// Stringer Interface
func (s Student) String() string {
	return fmt.Sprintf("%d: %s\n", s.GetId(), s.GetName())
}

// Generics
func Sum[T int | float64 | float32](numbers ...T) T {
	var total T
	for _, number := range numbers {
		total += number
	}
	return total
}

// Sample Project
type Payable interface {
	fmt.Stringer
	CalculatePay() float64 // Calculates monthly pay
	// We'll also rely on fmt.Stringer, so our types should implement String()
}

type SalariedEmployee struct {
	Name         string
	AnnualSalary float64
}

func (se SalariedEmployee) CalculatePay() float64 {
	return se.AnnualSalary / 12.0
}

func (se SalariedEmployee) String() string {
	return fmt.Sprintf("Salaried: %s (Annual: $%.2f)", se.Name, se.AnnualSalary)
}

type HourlyEmployee struct {
	Name        string
	HourlyRate  float64
	HoursWorked float64 // Hours worked in the month
}

func (he HourlyEmployee) CalculatePay() float64 {
	return he.HourlyRate * he.HoursWorked
}

func (he HourlyEmployee) String() string {
	return fmt.Sprintf("Hourly: %s (Rate: $%.2f/hr, Hours: %.1f)", he.Name, he.HourlyRate, he.HoursWorked)
}

type CommissionEmployee struct {
	Name           string
	BaseSalary     float64 // Monthly base
	CommissionRate float64 // e.g., 0.05 for 5%
	SalesAmount    float64
}

func (ce CommissionEmployee) CalculatePay() float64 {
	return ce.BaseSalary + (ce.CommissionRate * ce.SalesAmount)
}

func (ce CommissionEmployee) String() string {
	return fmt.Sprintf("Commission: %s (Base: $%.2f, CommRate: %.2f%%, Sales: $%.2f)",
		ce.Name, ce.BaseSalary, ce.CommissionRate*100, ce.SalesAmount)
}

func PrintEmployeeSummary[P fmt.Stringer](employee P) {
	fmt.Printf("  - Processing: %s\n", employee) // Relies on String() method
}

func ProcessPayroll(employees []Payable) {
	fmt.Println("\n--- Processing Payroll ---")
	totalPayroll := 0.0
	for _, emp := range employees {
		PrintEmployeeSummary(emp) // Generic function call
		pay := emp.CalculatePay()
		fmt.Printf("    Monthly Pay: $%.2f\n", pay)
		totalPayroll += pay
	}
	fmt.Printf("\nTotal Monthly Payroll: $%.2f\n", totalPayroll)
	fmt.Println("--------------------------")
}

func main() {
	structSample()
	sampleReceiver()
	bob := Student{1, "Bob"}
	displayPerson(bob)
	fmt.Println(bob.String())
	v := Sum[float64](10.5, 10.1, 20.4)
	fmt.Println(v)

	fmt.Println("Welcome to the Payroll Processor!")

	salEmp := SalariedEmployee{Name: "Alice Wonderland", AnnualSalary: 72000.00}
	hrEmp := HourlyEmployee{Name: "Bob The Builder", HourlyRate: 25.00, HoursWorked: 160.0}
	comEmp := CommissionEmployee{Name: "Charlie Chaplin", BaseSalary: 2000.00, CommissionRate: 0.10, SalesAmount: 15000.00}

	payrollList := []Payable{
		salEmp,
		hrEmp,
		comEmp,
		HourlyEmployee{Name: "Diana Prince", HourlyRate: 30.00, HoursWorked: 150.0},
	}

	ProcessPayroll(payrollList)
}
