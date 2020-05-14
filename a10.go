package main
import "fmt"

func main() {  
	a := [5] int {1,2,3,4,5}
	a_slice := a[2:4]
	b := [5] string {"a","b","c","d","e"}
	b_slice := b[0:3]

    fmt.Println("a_slice:", a_slice)
    fmt.Println("b_slice:", b_slice)
    fmt.Println("Length a_slice:", len(a_slice))
    fmt.Println("Length b_slice:", len(b_slice))

  
}