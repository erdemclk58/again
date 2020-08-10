package main

import(
		"fmt"
		"time"
		"math"
		"math/rand"
		"math/cmplx"
)

var(
	ToBe bool=false;
	MaxInt uint64=1<<64-1
	z complex128=cmplx.Sqrt(-5+12i)
)



var i,j int=1,2

func add(x ,y int) int {
	return x+y
}

func swap(x,y string)(string ,string ){
	return y,x
}

func split(sum int)(x,y int){
	x=sum*4/9
	y=sum-x
	return
}
func main(){
	fmt.Println("Hello, I am Erdem")
	fmt.Println("The time is ",time.Now())
	fmt.Println("The random number is",rand.Intn(16))
	fmt.Printf("Now you have %g problems.\n",math.Sqrt(7))
	fmt.Println("Pi is ",math.Pi)
	fmt.Println("Sum is ",add(42,12))
	a,b:=swap("hello","everyone")	
	fmt.Println(a,b)
	fmt.Println(split(21))
	var c,python,java="true","false","no!"
	k:=3
	fmt.Println(i,j,k,c,python,java)
	fmt.Printf("Type : %T Value %v\n",ToBe,ToBe)
	fmt.Printf("Type : %T Value %v\n",MaxInt,MaxInt)
	fmt.Printf("Type : %T Value %v\n",z,z)
	var d int
	var f float
	var b bool
	var s string
	fmt.Println(d,f,b,s)
}
