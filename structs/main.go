package main

import "fmt"

type person struct {
	firstName string
	lastName string
	contactInfo

}

type contactInfo struct {
	email string
	zipCode int
}

func main(){
	/*alex := person{firstName:"Alex", lastName:"Anderson"}
	frank := person{"Frank", "Sinatra"}
	fmt.Println(alex)
	fmt.Println(frank)

	var john person
	fmt.Printf("%+v", john)*/
	
	jim := person{
		firstName: "Jim",
		lastName: "Party",
		contactInfo: contactInfo{
			email: "jim@gmail.com",
			zipCode: 94306,
		},
	}
	// This would not work since we are passing 
	// a value and not a reference
	jim.updateName("jimmy")
	jim.print()
	// This works along with changing the receiver 
	// of update name to *person
	jimPointer := &jim
	jimPointer.updateName("jimmy")
	jim.print()
	// or
	// Dont update the method to have a receiver of type
	// Golang implicitly makes it a pointer if the receiver 
	// is a pointer
	jim.updateName("jimmyBoy")
	jim.print()

	// Example of pass by reference in golang
	mySlice := []string{"Hi", "There", "How", "Are", "You"}
	updateSlice(mySlice)
	// Example showing slices are pass by reference
	fmt.Println(mySlice)
	// Golang reference types vs value types
	/*                          ||         
	 * int                      ||         Slices
	 * float                    ||         Maps
	 * string                   ||         Channels
	 * bool                     ||         Pointers
	 * structs                  ||         Functions
	 * 
	 * 
	*/
}	

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (p *person) updateName(newFirstName string) {
	p.firstName = newFirstName
}

func updateSlice(s []string) {
	s[0] = "Bye"
}