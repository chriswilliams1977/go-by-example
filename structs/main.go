package main

func main() {
	bob := Person{
		FirstName: "Bob",
		LastName:  "Smith",
		Age:       42,
	}

	pat := Mother{Person{
		FirstName: "Pat",
		LastName:  "Smith",
		Age:       73,
	},
	}

	bella := Dog{}

	doSomething(&bob)
	doSomething(&pat)
	doSomething(&bella)
}
