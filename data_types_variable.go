package main

import (
	"fmt"
	"reflect"
	"strconv"
)

var P int = 25 //global variable

func main() {
	fmt.Println("Hello world...!")
	name := "ashish" //implicite declaration
	fmt.Println(name)

	var greet string = "welcome"
	fmt.Println(greet)

	var city string = "mumbai"
	fmt.Print(greet, "\n") //new line

	fmt.Println("Welcome to ", greet, ", ", city)

	fmt.Printf("welcome to %s ", city) //format print
	fmt.Printf("\n")
	//%v formats the value in default format
	fmt.Printf("Welcome to the city %v\n", city)

	//%d formats decimal integres
	var grades int = 50
	var sub string = "phy"
	fmt.Printf("grade achived in test is %d\n", grades)
	fmt.Printf("grade achived in %v test is %d\n", sub, grades)

	var user string
	user = "Ashish"
	fmt.Println(user)

	//short hand
	var sa, t string = "foo", "bar"
	fmt.Println(sa, t)
	var (
		s string = "foo"
		i int    = 5
	)
	fmt.Println(s)
	fmt.Println(i)
	//shorthand
	pech := "Hi Ashish"
	fmt.Println(pech)
	pech = "Hi Ajay"
	fmt.Println(pech)

	city = "pune" //local variable
	{
		country := "india"
		fmt.Println(country)
		fmt.Println(city)
	}

	fmt.Println(city)
	//fmt.Println(country) //fail due to inner block decl.
	//zero value
	/*
		bool - false
		int - 0
		float64 - 0.0
		string - ""
	*/
	var a int
	fmt.Println(a)
	//user input
	var sir string
	fmt.Println("Enter your sir name")
	//fmt.Scanf("%s", &sir)
	fmt.Println("hey there, ", sir)

	//multiple input
	// var sirname string
	// var is_boy bool
	fmt.Print("Enter sirname and are you boy?\n")
	//fmt.Scanf("%s %t", &sirname, &is_boy)
	//fmt.Scanf("%v %t", &sirname, &is_boy)
	//fmt.Println(sirname, is_boy)

	//multiple input
	// var x string
	// var y int
	// fmt.Print("Enter a string and a number: ")
	// count, err := fmt.Scanf("%s %d", &x, &y)

	// fmt.Println("count: ", count)
	// fmt.Println("error: ", err)
	// fmt.Println("x: ", x)
	// fmt.Println("y: ", y)

	/*
		Enter a string and a number: ash ash
		count:  1
		error:  expected integer
		x:  ash
		y:  0
		----------------
		Enter a string and a number: ash 87
		count:  2
		error:  <nil>
		x:  ash
		y:  87
	*/
	// types of variable using %T
	var (
		z       int     = 10
		msg     string  = "here"
		isCheck bool    = true
		amount  float32 = 4566.44
	)
	// fmt.Printf("varible z=%v is of type %T\n", z, z)
	// fmt.Printf("varible msg=%v is of type %T\n", msg, msg)
	// fmt.Printf("varible isCheck=%v is of type %T\n", isCheck, isCheck)
	// fmt.Printf("varible amount=%v is of type %T\n", amount, amount)

	fmt.Printf("varible z=%v \n", reflect.TypeOf(z))
	fmt.Printf("varible msg=%v \n", reflect.TypeOf(msg))
	fmt.Printf("varible isCheck=%v \n", reflect.TypeOf(isCheck))
	fmt.Printf("varible amount=%v \n", reflect.TypeOf(amount))

	//typecasting
	//int to float
	// var b int = 90
	// var bf float64 = float64(b)
	//fmt.Printf("%0.2f\n", bf)
	//float to int
	var bf float64 = 90.23
	var b int = int(bf)
	fmt.Printf("%v\n", b) //loose the data

	//strconv package
	//int to string
	var g int = 49
	var h string = strconv.Itoa(g)
	fmt.Printf("%q", h)

	//string to int
	// var j string = "20"
	// k, err := strconv.Atoi(j)
	// fmt.Printf("%v, %T \n", k, k)
	// fmt.Println("%v, %T", err, err)

	// var j string = "20abc"
	// k, err := strconv.Atoi(j)
	// fmt.Printf("%v, %T \n", k, k)
	// fmt.Println("%v, %T", err, err)  //error

}
