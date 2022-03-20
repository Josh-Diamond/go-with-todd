package main

import "fmt"

func main() {
	fmt.Println("Josh: Remy is the coolest ever!")
	woof()
	fmt.Println("Josh: Let's go get a cookie")

	for i := 0; i < 100; i++ {
		if i%2 != 0 {
			fmt.Println(i)
		}
	}
	cookie()
}

func woof() {
	fmt.Println("Remy: Woof woof!!")
}

func cookie() {
	fmt.Println("Narrator: and then Remy ate a cookie...")
}

// control flow: (how a computer reads code, order of execution and evaluation)
// 1. sequence
// 2. loop: iterative
// 3. conditional
