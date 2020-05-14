package main

import "fmt"


func main() {
 var name[10] string
 var i int
 
 name[0]="hi"
 name[1]="hii"
 name[2]="hiii"
 
 //fmt.Println(name)
 fmt.Println(len(name))
 fmt.Println(name[0])
 
 for i=0;i<10;i++{
 fmt.Println(name[i])
 }
 
 for i=0;i<len(name);i++{
 fmt.Println(name[i])
 }
 
}