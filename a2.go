package main
import "fmt"
func add(x float32, y int)int{
return int(x)+y
}
func stringg(s1,s2 string)(string,string){
return "hi "+s1 , " hello "+s2
}
func main() {

 var m float32=3
 var n int=int(m)
  a:="hi"
 b:="hello"
    fmt.Println(add(m,n))
	fmt.Println(stringg(a,b))
}
