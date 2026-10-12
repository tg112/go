package main

import (
	"errors"
	"fmt"
)

// Composition --> Has a relationship
// Car is composed of many parts (Engine, Doors)

type Address struct {
	Street string
	City   string
	State  string
	Zip    string
}

func (a Address) FullAddress() string {
	if a.Street == "" && a.City == "" {
		return "No address provided"
	}
	return fmt.Sprintf("%s %s %s %s", a.Street, a.City, a.State, a.Zip)
}

type Customer struct {
	CustomerID      int
	Name            string
	Email           string
	BillingAddress  Address // embedded
	ShippingAddress Address
}

func (c Customer) PrintDetails() {
	fmt.Printf("CustomerID: %d\n", c.CustomerID)
	fmt.Printf("Name: %s\n", c.Name)
	fmt.Printf("Email: %s\n", c.Email)
	fmt.Printf("BillingAddress: %s\n", c.BillingAddress.FullAddress())
	fmt.Printf("ShippingAddress: %s\n", c.ShippingAddress.FullAddress())
}

func buildCustomer() Customer {
	cust1 := Customer{
		CustomerID: 1,
		Name:       "John Doe",
		Email:      "sample@sample.com",
		BillingAddress: Address{
			Street: "123 street",
			City:   "123 City",
			State:  "123 State",
			Zip:    "123 Zip",
		},
		ShippingAddress: Address{
			Street: "123 street",
			City:   "123 City",
			State:  "123 State",
			Zip:    "123 Zip",
		},
	}
	return cust1
}

type ContactInfo struct {
	Email string
	Phone string
}

func (ci ContactInfo) DisplayContact() string {
	return fmt.Sprintf("Email: %s, Phone: %s", ci.Email, ci.Phone)
}

type Company struct {
	Name string
	Address
	ContactInfo
	BusinessType string
}

func (c Company) GetProfile() {
	fmt.Printf("Company Name: %s\n", c.Name)

	fmt.Printf("Location: %s\n", c.FullAddress())
	fmt.Printf("Street (promoted): %s\n", c.Street)

	fmt.Printf("Email (promoted): %s\n", c.Email)
	fmt.Printf("Business Type: %s\n", c.BusinessType)
}

type CompanyWithOwnEmail struct {
	Name string
	Address
	ContactInfo
	Email string // This Email field shadows the one in ContactInfo
}

func structEmbedded() {
	fmt.Println("--- Struct Embedding-----")

	comp := Company{
		Name:         "Innovate Solutions Inc.",
		Street:       "789 Innovation Drive",
		City:         "Techville",
		State:        "TS",
		Zip:          "12345",
		Email:        "contact@innovate.com",
		Phone:        "555-0100",
		BusinessType: "Technology",
	}

	comp.GetProfile()

	fmt.Printf("\nDirect access to comp.City: %s\n", comp.City) // Promoted from Address
	fmt.Printf("Direct access to comp.Phone: %s\n", comp.Phone) // Promoted from ContactInfo

	fmt.Printf("Embedded Address struct: %+v\n", comp.Address)
	fmt.Printf("Embedded ContactInfo struct: %+v\n", comp.ContactInfo)

}

// Sample Project
type Account struct {
	AccountNumber string
	Balance       float64
	OwnerName     string
}

func (acc *Account) Deposit(amount float64) error {
	if amount < 0 {
		return errors.New("deposit must be positive")
	}
	acc.Balance += amount
	fmt.Printf("Account Balance: %.2f\n", acc.Balance)
	return nil
}

func (acc *Account) Withdraw(amount float64) error {
	if amount < 0 {
		return errors.New("withdraw must be positive")
	}
	if amount > acc.Balance {
		return errors.New("withdraw must be greater than balance")
	}
	acc.Balance -= amount
	fmt.Printf("Account Balance: %.2f\n", acc.Balance)
	return nil
}

func (acc *Account) GetBalance() float64 {
	return acc.Balance
}

func (acc *Account) String() string {
	return fmt.Sprintf("Account Number: %s, Balance: %.2f, Owner: %s", acc.AccountNumber, acc.Balance, acc.OwnerName)
}

type SavingsAccount struct {
	Account
	InterestRate float64
}

func (sa *SavingsAccount) AddInterest() {
	interest := sa.Balance * sa.InterestRate
	fmt.Printf("Adding interest $%.2f to savings account %s. ", interest, sa.AccountNumber)
	err := sa.Deposit(interest) // Uses promoted Deposit method
	if err != nil {
		fmt.Printf("AddInterst: Error depositing $%.2f to savings account. %v\n", interest, err)
	}
}

type OverdraftAccount struct {
	Account        // Embed Account struct
	OverdraftLimit float64
}

func (oa *OverdraftAccount) Withdraw(amount float64) error {
	if amount <= 0 {
		return errors.New("withdrawal amount must be positive")
	}
	// Allow withdrawal up to Balance + OverdraftLimit
	if (oa.Balance + oa.OverdraftLimit) < amount {
		return fmt.Errorf("withdrawal of $%.2f exceeds overdraft limit for %s. Available including overdraft: $%.2f",
			amount, oa.AccountNumber, oa.Balance+oa.OverdraftLimit)
	}
	oa.Balance -= amount // Balance can go negative
	fmt.Printf("Withdrew $%.2f from overdraft account %s. New Balance: $%.2f\n", amount, oa.AccountNumber, oa.Balance)
	return nil
}

func bankInfo() {
	fmt.Println("--- Bank Account System ---")

	savAcc := SavingsAccount{
		// Initialize the embedded Account
		AccountNumber: "SAV001",
		Balance:       1000.00,
		OwnerName:     "Alice Saver",
		InterestRate:  0.02, // 2%
	}
	fmt.Println("\n--- Savings Account Operations ---")
	fmt.Println(savAcc.Account.String())

	err := savAcc.Deposit(200.00)
	if err != nil {
		fmt.Printf("Error depositing $%.2f to savings account. %v\n", 200.00, err)
	}
	savAcc.AddInterest()
	err = savAcc.Withdraw(50.00)
	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println("Final Savings Details:", savAcc.Account.String())

	ovdAcc := OverdraftAccount{
		AccountNumber:  "OVD002",
		Balance:        100.00,
		OwnerName:      "Bob Spender",
		OverdraftLimit: 200.00,
	}

	fmt.Println("\n--- Overdraft Account Operations ---")
	fmt.Println(ovdAcc.Account.String())

	err = ovdAcc.Deposit(50.00)
	if err != nil {
		fmt.Println("Error:", err)
	}
	err = ovdAcc.Withdraw(200.00)
	if err != nil {
		fmt.Println("Error:", err)
	}

	err = ovdAcc.Withdraw(100.00)
	if err != nil {
		fmt.Println("Error:", err)
	}
	err = ovdAcc.Withdraw(100.00)
	if err != nil {
		fmt.Println("Error (expected for overdraft limit):", err)
	}

	fmt.Println("Final Overdraft Details:", ovdAcc.Account.String())
}

func main() {
	customer := buildCustomer()
	customer.PrintDetails()
	structEmbedded()

}
