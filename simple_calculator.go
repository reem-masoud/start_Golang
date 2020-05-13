package main

import "fmt"


func main() {

var a int
var b int
var op string

fmt.Println("enter first number")
fmt.Scanln(&a)
fmt.Println("enter operator")
fmt.Scanln(&op)
fmt.Println("enter secon number")
fmt.Scanln(&b)
 var result int
switch op {
case "+":
result=a+b
fmt.Printf("result = %d",result )

case "-":
result=a-b
fmt.Printf("result = %d",result )

case "*":
result=a*b
fmt.Printf("result = %d",result )

case "/":
if b==0{
fmt.Printf("error division by zero ")
}else{
result=a/b
fmt.Printf("result = %d",result )
}

case "%":
result=a%b
fmt.Printf("result = %d",result )
default:
fmt.Println("unavailable operator")
}

}