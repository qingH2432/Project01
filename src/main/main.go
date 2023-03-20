package main

import "fmt"

func main() {
	const PI float64 = 3.14159
	const daysInWeek = 7
	const message = "Hello World!"
	fmt.Println(PI)
	fmt.Println(daysInWeek)
	fmt.Println(message)
	const (
		Sunday = iota
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
	)
	fmt.Println(Friday)

}
