package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	var name, lastname string

	fmt.Println("Enter your name and lastname: ")
	_, err := fmt.Scanln(&name, &lastname)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input %v\n", err)
	}

	fmt.Println("Enter your age: ")
	scanner.Scan()

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input %v\n", err)
	}

	age := scanner.Text()

	ageInt, _ := strconv.Atoi(age)

	fmt.Printf("Приятно познакомиться, %s. Я 5 лет назад познакомился с человеком, у которого тоже фамилия %s, вам тогда было %d. Как молоды мы были!\n", name, lastname, ageInt)
}
