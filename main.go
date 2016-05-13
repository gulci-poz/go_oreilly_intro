package main

import (
	"fmt"
)

// globalne, może być też var
const (
	greeting = "hello %s\n"
	person = "world"
	message = "%d %d\n"
	answer1 = iota
	answer2
)

func main() {
	// var message string
	// message = "hello\n"
	// lub

	// greeting := "hello %s\n"
	// person := "world"
	fmt.Printf(greeting, person)
	fmt.Printf(message, answer1, answer2)

	var pi float64 = 3.14
	// można rzutować
	// pi := float64(3.14)
	fmt.Printf("%.2f\n", pi)

	isTrue := true
	fmt.Printf("%t\n", isTrue)

	b := byte(65)
	// pokaże wartość hex
	fmt.Printf("%x\n", b)

	num := uint(9)
	fmt.Printf("%d\n", num)
}

// GOPATH bez src, bin oraz pkg również będą widoczne
// możemy użyć go install
// go install go_oreilly_intro
// w tym folderze mamy plik main.go z funkcją main

// dodatkowo, poprawne formatowanie kodu
// go fmt go_oreilly_intro

// pomoc za pośrednictwem godoc <pakiet funkcja>
