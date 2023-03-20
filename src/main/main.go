package main

import "fmt"

type Weekday int

func main() {
	const (
		Sunday Weekday = iota
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
	)
	var today Weekday = Friday
	fmt.Println(today)
}
