package main

import "fmt"

type contactInfo struct {
	email string
	zipCode int
}

type person struct {
	firtName string
	lastName string
	contactInfo
}

func main() {
	jim := person{
		firtName: "Jim",
		lastName: "Party",
		contactInfo: contactInfo{
			email: "jim@gmail.com",
			zipCode: 94000,
			},
	}
	// fmt.Println(jim)
	// fmt.Printf("%+v", jim)
	jimPointer := &jim // &variable = give me the memory address of the value this variable is pointing at 
	jimPointer.updateName("Jimmy")
	jim.print()
}

// *address = give me the value this memory address is pointing at(Turn address into value)
// &value = give me the memory address of the value this variable is pointing at(Turn value into address)

func (pointerToperson *person) updateName(newFirstName string) {
	(*pointerToperson).firtName = newFirstName // *pointer = give me the value this memory address is pointing at
}

func (p person) print() {
	fmt.Printf("%+v", p)
}