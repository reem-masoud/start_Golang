package main
import "fmt"

func main() {

var ms *mystruct
//fmt.Println(ms)
//ms=&mystruct{foo: 42}
ms=new(mystruct)
//(*ms).foo=42
ms.foo=42
//fmt.Println((*ms).foo)
fmt.Println(ms.foo)
//fmt.Println(ms)

}

type mystruct struct{
foo int
}