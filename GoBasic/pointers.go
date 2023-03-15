package main

import "fmt"

// creating pointers |Deferencing pointers| The new function | Working with nil | Types with internal pointers 
func main2(){

a:= 43
fmt.Println(a)

c:= 45
d := &c
fmt.Print(c,d ,"\n")
fmt.Print(c,*d)
}

