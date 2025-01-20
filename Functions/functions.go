package main

import "fmt"

func add(num1 int, num2 int) {

	sum := num1 + num2
	fmt.Println(sum)

}

func newadd(num1 int, num2 int) int{
    sum1 := num1 + num2 
		return sum1 
}


func multiple(num1 int, num2 int) (int,int){
	sum1 := num1 + num2
	mul1 := num1 * num2 

	return sum1, mul1  
} 


func printName(name string){
	fmt.Println("This is", name) 
}

func main() {
	a := 10
	b := 20
	add(a, b)
	add(2, 3)
	add(4, 5) 
  fmt.Println(newadd(12,12)) 

	p, q := multiple(2,3) 

	fmt.Println(p)
	fmt.Println(q)

  printName("Reza")
}