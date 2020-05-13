package main

import "fmt"


func main() {
var pass int=1234
var x int
var i int
fmt.Print("please enetr password  ")
fmt.Scanln(&x)
if(x==pass){
fmt.Println("welcome")

}else{
for i=0;i<2;i++{
fmt.Print("please try again ")
fmt.Scanln(&x)
if(x==pass){
fmt.Println("welcome")
break
}
}
}


}