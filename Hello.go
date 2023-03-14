package main

import "fmt"

// we can declare variable in var block

var (

	name string = "shad"
	age int = 24 
	desig string = "student"
)

// All the variable must be used / if they are declared at function level
// use lowercase for short term variable use uppercase for constant
// we can decalre variable and assign it latter
// there are two types of declartaion level package level and function level 
// for conversion string to int we need to import package 

func main() {
	
	// without var keyword

	greet := "hi"
	time := 12.10
	// declare variable 
	var work string 
	work = "Programming"
	// other way to declare

	var job string = "programmmer"

	fmt.Printf( "Name: %s  want to learn  GO at %d who is a student %s  ", name, age, desig )
	fmt.Printf("%s %s  its %f of morning get up and starting working  on %s don't forget you are a %s",greet,name ,time,work,job)
}

// Variables  declaration



