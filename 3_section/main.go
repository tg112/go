package main

import "fmt"

func array_example() {
	var numbers [2]int
	fmt.Printf("%+v\n", numbers)

	numbers[0] = 10
	numbers[1] = 20
	fmt.Printf("%+v\n", numbers)

	primes := [4]int{2, 3, 5, 7}
	fmt.Printf("%+v\n", primes)
	primes[3] = 11

	for i := range len(primes) {
		fmt.Printf("%d\n", primes[i])
	}

	var matrix [2][3]int
	matrix[0][0] = 1
	matrix[0][1] = 2
	matrix[1][0] = 3
	fmt.Printf("%+v\n", matrix)
}

func slice_example() {
	names := []string{"John", "Paul", "George", "Ringo"}
	fmt.Printf("%+v\n", names)
	items := make([]int, 3, 5)
	items = append(items, 1)
	items = append(items, 2)
	items = append(items, 3)
	items = append(items, 4)
	items[0] = 5
	fmt.Printf("Items: %+v Len: %d Cap: %d\n", items, len(items), cap(items))

	fmt.Println("--------- Advanced slice-example---------------")
	original := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s1 := original[2:5] // low:high (excluding high)
	s2 := original[:4]  // low:high (high is not included)
	s3 := original[6:]  // low:high (including low)
	s4 := original[:]   // copy
	fmt.Printf("Original: %+v Len: %d, Cap: %d\n", original, len(original), cap(original))
	fmt.Printf("s1: %+v Len: %d, Cap: %d\n", s1, len(s1), cap(s1))
	fmt.Printf("s2: %+v Len: %d, Cap: %d\n", s2, len(s2), cap(s2))
	fmt.Printf("s3: %+v Len: %d, Cap: %d\n", s3, len(s3), cap(s3))
	fmt.Printf("s4: %+v Len: %d, Cap: %d\n", s4, len(s4), cap(s4))
}

func map_example() {
	studentsGrades := map[string]int{
		"James":      32,
		"Moneypenny": 40,
		"Alice":      80,
	}
	fmt.Printf("%+v\n", studentsGrades)
	studentsGrades["James"] = 78
	fmt.Printf("%+v\n", studentsGrades)
	key := "Alice"
	alice, ok := studentsGrades[key]
	if ok {
		fmt.Printf("%s %+v\n", key, alice)
	}

	configs := make(map[string]int)
	fmt.Printf("%+v %T\n", configs, configs)
	if configs == nil {
		fmt.Println("Configs is nil")
	}
}

func modifyValue(val int) {
	val = val * 10
	fmt.Printf("modifyVal: %+v\n", val)
}
func modifyPointer(val *int) {
	if val == nil {
		fmt.Println("val is nil")
		return
	}
	*val = *val * 10 // dereferencing
	fmt.Printf("modifyVal: %+v\n", *val)
}

func pointer_example() {
	//age := 10
	//agePtr := &age
	//fmt.Printf("%+v\n", *agePtr)
	//fmt.Printf("agePort: %d\n", agePtr)
	//fmt.Printf("agePort: %d\n", &age)
	modifyValue(10)
	num := 30
	modifyPointer(&num)
	fmt.Println(num)
}

type Contact struct {
	ID    int
	Name  string
	Email string
	Phone string
}

var contactList = []Contact{}
var contactIndexByName map[string]int
var nextID = 1

func init() {
	contactList = make([]Contact, 0)
	contactIndexByName = make(map[string]int)
}

func addContact(name, email, phone string) {
	if _, exist := contactIndexByName[name]; exist {
		fmt.Printf("Contact already exists: %+v\n", name)
		return
	}

	newContact := Contact{
		ID:    nextID,
		Name:  name,
		Email: email,
		Phone: phone,
	}
	nextID++
	contactList = append(contactList, newContact)
	contactIndexByName[name] = newContact.ID
	fmt.Printf("Contact added: %+v\n", name)
}

func findContact(name string) *Contact {
	index, exist := contactIndexByName[name]
	if exist {
		return &contactList[index]
	}
	return nil
}

func ListContacts() {
	fmt.Println("----ListContacts----")
	if len(contactList) == 0 {
		fmt.Println("ListContacts is empty")
		return
	}
	for _, contact := range contactList {
		fmt.Printf("%+v\n", contact)
	}
}

func sample_project() {
	addContact("Alice", "alice@sample.com", "111-2222")
	addContact("Bob", "bob@sample.com", "111-2222")
	addContact("Tom", "tom@sample.com", "111-2222")
	addContact("One", "one@sample.com", "111-2222")

	ListContacts()

	bob := findContact("Bob")
	if bob == nil {
		fmt.Println("No Bob contact found")
	} else {
		fmt.Println("Bob contact found")
	}
}

func main() {
	array_example()
	slice_example()
	map_example()
	pointer_example()
	sample_project()
}
