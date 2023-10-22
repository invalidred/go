package main

import "fmt"

func mayPanic() {
	panic("a problem")
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered. Error:\n", r)
		}
	}()

	mayPanic()

	// Since mayPanic() panics, the line below will not be
	// executed. Since the panic is recovered in the defer func()
	// the excution will continue in the deffered closure.
	fmt.Println("After mayPanic()")
}
