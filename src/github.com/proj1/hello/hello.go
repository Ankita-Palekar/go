package main

import (
	"fmt"
)

var isActive bool
var enabled, disabled = true, false

func main() {
	//fmt.Println(string.Reverse("What's up? Ankita the Great!"))
	//const PI float32 = 3.1467
	//
	//fmt.Println(PI)

	//var available string// local variable
	//valid := false      // brief statement of variable
	//available = true    // assign value to variable
	//
	//fmt.Println(available)
	//fmt.Println(valid)
	//
	////var c complex64 = 5+5i
	////output: (5+5i)
	////fmt.Printf("Value is: %v", c)
	//
	//
	//s := "hello"
	//c := []byte(s)
	//c[0]  = 'c'
	//s2 := string(c)
	//fmt.Println(s2)

	//var arr [10]int
	//arr[0] = 42
	//arr[1] = 32

	//fmt.Printf("%d", arr[0])
	//fmt.Println(available)

	//slice := []byte {'a', 'b', 'c', 'd'}
	////fmt.Printf("%d", append(slice, '6'));
	//slice = append(slice, 't')
	//fmt.Println(slice)
	//
	//var copy_slice []byte
	//copy(copy_slice, slice)
	//fmt.Print(copy_slice)

	//	using map

	//numbers := make(map[string]int)
	//numbers["key1"] = 1
	//numbers["key2"] = 2
	//fmt.Println(numbers["key1"])
	//
	//numberSliced := make([]int, 10, 100)
	//numberSliced[0] = 2
	//fmt.Println(numberSliced[2])
	//
	//// understanding the diff between the make and the new
	//
	//p := new(map[string]int)   // p has type: *chan int
	//c := make(map[string]int)  // c has type: chan int
	//k := c
	//k["m"] = 1
	//fmt.Println(k)
	//fmt.Print(p)
	//fmt.Print(c)

	//	understanding of the loops
	//	sum := 0;
	//	for index:=0; index < 10 ; index++ {
	//		sum += index
	//	}
	//	fmt.Println("sum is equal to ", sum)


	//example with the range and the key value pair
//	for k, v := range numbers{
//		fmt.Println("map's key:", k)
//		fmt.Println("map's val:", v)
//	}
//
//	fmt.Println("-----------------")
//
//	myName := "ankita"
////	example if you want to ignore the key value pair
//	for _, v := range myName{
//		//fmt.Println("map's key:", k)
//		fmt.Println("map's val:", v)
//	}


////switch statement
//	i := 10
//	switch i {
//	case 1:
//		fmt.Println("i is equal to 1")
//	case 2, 3, 4:
//		fmt.Println("i is equal to 2, 3 or 4")
//	case 10:
//		fmt.Println("i is equal to 10")
//	default:
//		fmt.Println("All I know is that i is an integer")
//	}


//maxNum := maxNum(0, 1889)
//fmt.Println(maxNum)
//
//
//sum, product := sumAndProduct(1, 2)
//
//fmt.Println("sum = ", sum, "product = ", product)
//
////using variadic funcitons
//
//	addNumbers(1,2,3,4,5,6,7,8,9,10)


	//learning deffer function

	defer doSomeThing()
	doSomeThingElse()
	defer doSomeThing1()
	doSomeThingElse2()
}
//	Learning Functions



func maxNum(a int, b int) int{
	if(a > b){
		return a
	}
	return b
}


func sumAndProduct(a int, b int) (int, int){
	return a + b, a * b
}

//variadic functions

func addNumbers(nums ... int){
	var total int

	for _, num := range nums {
		total += num
	}

	fmt.Println("total sum = ", total)
}


func doSomeThingElse(){
	fmt.Println("doSomeThingElse()")
}

func doSomeThing(){
	fmt.Println("doSomeThing()")
}

func doSomeThingElse2(){
	fmt.Println("doSomeThingElse2()")
}

func doSomeThing1(){
	fmt.Println("doSomething1()")
}
//understanding defer