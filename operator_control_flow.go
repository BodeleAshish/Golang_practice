package main

import "fmt"

func main() {
	//comparison operators
	var city string = "mumbai"
	var city_1 string = "bombay"
	fmt.Println(city == city_1)
	fmt.Println(city != city_1)
	var a, b int = 10, 10
	fmt.Println(a > b)
	fmt.Println(a >= b)
	fmt.Println(a <= b)
	fmt.Println(a < b)
	//arithmetic operators
	var c, d string = "foo", "bar"
	fmt.Println(c + d)
	//fmt.Println(c-d) //fail
	//unary operator
	var i int = 1
	i++
	fmt.Println(i)
	i--
	fmt.Println(i)
	var f float64 = 5.9
	f++
	fmt.Println(f)
	f--
	fmt.Println(f)
	//var e string = "vi"
	//var g int = 45
	//fmt.Println(r + g)  //error
	//var h float64 = 70, 66
	//fmt.Println(g + h)  //error
	var e, g float64 = 24.4, 3.0
	fmt.Println(e / g)
	fmt.Println(int(e) % int(g))

	//if else and else if stmt
	var h string = "happy"
	if h == "happy" {
		fmt.Println(h)
	}
}
