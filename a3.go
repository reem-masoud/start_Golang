package main

import "fmt"
import "os"


func main(){
var x int
var y int

fmt.Scan(&x)
fmt.Scan(&y)

var a=x+y
var b=x-y
var d=x*y

x++
y--
fmt.Fprintf(os.Stdout,"sum is %d \n",a)
fmt.Fprintf(os.Stdout,"subtract is %d \n",b)

fmt.Fprintf(os.Stdout,"Multiply is %d \n",d)

fmt.Println(x)
fmt.Println(y)



}