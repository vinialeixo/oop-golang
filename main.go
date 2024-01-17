package main

import "fmt"

// // type definition
// type data int

// // Defining a metho with non-struct type receiver
// func (d1 data) multiply(d2 data) data {
// 	return d1 * d2
// }

// func main() {
// 	value1 := data(10)
// 	value2 := data(20)
// 	res := value1.multiply(value2)
// 	fmt.Println("Final result: ", res)
// }

type author struct {
	name      string
	branch    string
	particles int
}

// method with a receiver of author
func (a *author) show(abranch string) {
	(*a).branch = abranch
}

func main() {
	//Initializinh the values of the author structure
	res := author{
		name:   "Vinicius",
		branch: "CSA",
	}

	fmt.Println("Author's name: ", res.name)
	fmt.Println("Branch Name (before): ", res.branch)

	//creating a pointer
	ponteiro := &res

	//calling the show method

	ponteiro.show("ABC")
	fmt.Println("Author's name: ", res.name)
	fmt.Println("Branch Name (after): ", res.branch)
}
