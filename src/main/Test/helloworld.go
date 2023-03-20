package main

import (
	"Project01/src/main/myMath"
	"fmt"
)

func main() {
	fmt.Println("Hello World!")
	fmt.Println(mathClass.Add(1, 1))
	fmt.Println(mathClass.Sub(1, 1))
	fmt.Println("Google" + "Runoob")
	var stockcode = 123
	var enddate = "2020-12-31"
	var url = "Code=%d endDate=%s"
	var targetUrl = fmt.Sprintf(url, stockcode, enddate)
	fmt.Printf(url, stockcode, enddate)
	fmt.Println()
	fmt.Println(targetUrl)
	var b bool = true
	fmt.Println(b)
}
