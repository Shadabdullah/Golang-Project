// Array > Creation | built in | working with arrays
// slices > Creation | built in | working with slicex

package main

import(
	"fmt"
)

func main(){
// Array declare the size also
grades:=[3] int { 65,34,34}
fmt.Print(grades)
var student[3] string

fmt.Print(student)
student[0]="SHAD"
student[1]="shad"
fmt.Print(student)
//built in len function
fmt.Print("lenght of array ", len(student))

// matring in Go
var identyMatrix [3][3]int
identyMatrix[0] =[3] int {1,0,1}
identyMatrix[1] =[3] int {1,1,1}
identyMatrix[2] =[3] int {0,0,0}

fmt.Print(identyMatrix)

// different then other language example python in python it points to same array

a:= [...] int{1,2,3}
b:= a
b[1]= 5

fmt.Println(a)
fmt.Println(b)

// but we can point to same data using &

c:=[...] int {1,2,4}
d:= &c
d[0]= 10
fmt.Print(c,d)

// Slice is bit different then array we can declare using square bracket and it have one
// more function cap and its point to same array changing in one will reflect in another
e:=[] int {1,2,4}
f:=e
f[0]= 10
fmt.Print(e,f)
fmt.Printf("length: %v  capsity %v",len(e),cap(f))







}