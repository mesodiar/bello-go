package main

import (
	"fmt"
	foobarqix "github.com/mesodiar/bello-go/foobar"
)



func split(text string) []string {
	var result []string
	for i:= 0;i<= len(text)-2; i++ {

		if i % 2 == 0 {
			result = append(result, text[i:i+2])
		}
	}
	return result
}


func main() {
	var s []int
	text := "ABCDEFG"

	fmt.Println(split(text))
	// fmt.Println(len(text))
	// x := [...]int{4, 5, 6, 7, 8}
	// s = x[1:3]
	// s = make([]int, 3, 5)

	// s = x[1:3]
	// s = append(s, 1, 2, 3)
	// x[1] = 15
	fmt.Println(s)
	fmt.Println(len(s), cap(s))

	// fmt.Println(s == nil)
}

func mainOfArray() {
	x := [...]int{4, 5, 6, 7}

	x[1] = 12

	fmt.Printf("%T %v\n", x, x)
	for _, v:= range x{
		fmt.Println(v)
	}
}




func mainImport() {
	fmt.Println(foobarqix.Title)
}

func mainPointer() {
	// var p *int
	var p = new(int)

	fmt.Println(p == nil)
	fmt.Println(*p)

	if p == nil {
		fmt.Println(*p)
	}

	i := 42
	p = &i
	
	fmt.Printf("%x %x\n", *p, i)
	fmt.Println(*p, i)

	*p = 43
	fmt.Println(*p, i)

}


// func prime(n int) {
// 	for i:=2; i<=n; i++ {
// 		count := 0
// 		for j:=1; j<=i; j++ {
// 			if i%j==0 {
// 				count++
// 			}
// 		}
// 		if count == 2{
// 			fmt.Println(i)
// 		}
// 	}
// }

func mainOfBasic() {
	var name string
	name = "Mils"
	
	fmt.Println("Hello,", name)
	fmt.Printf("Hello, %q\n", name)

	var age int = 27
	fmt.Printf("I'm %d years old\n", age)

	var nickName = "Mil"
	fmt.Printf("type is %T\n", nickName)

	sex := "female"
	fmt.Printf("%v\n", sex)

	fmt.Printf("1 + 2 = %d\n", add(1,2))

	a, b := 1, 2
	a, b = swap(a, b)
	fmt.Println(a, b)

}

func div (a int, b int) {
	return 
}
// func greeting(name string) {} //greeting("Yod") => "hello, Yod"

func add(a int, b int) int {
	return a + b
}

func swap(a, b int) (int, int) {
	return b, a
}

func greeting (name string) {
	fmt.Printf("Hello %s\n", name)
}
