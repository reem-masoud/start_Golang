package main

import "fmt"
import "os"


func main(){
var a int
var b int

fmt.Scan(&a, &b)
if (a>=b){
var c=a/b
fmt.Fprintf(os.Stdout, "c is %d",c)
}else{
var c=b
fmt.Fprintf(os.Stdout, "c is %d",c)
}



}