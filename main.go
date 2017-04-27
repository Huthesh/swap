package main 

import (
"fmt"
)

func main() {
	a:=10
	b:=20
	Swap(&a,&b)
	fmt.Println(a)
	fmt.Println(b)
}

func Swap(a,b *int){
	var temp int =*a
	*a=*b
	*b=temp
}