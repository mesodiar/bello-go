package main

import "fmt"

func main() {
	prime(18)
}

func prime(n int) {
	for i:=2; i<=n; i++ {
		count := 0
		for j:=1; j<=i; j++ {
			if i%j==0 {
				count++
			}
		}
		if count == 2{
			fmt.Println(i)
		}
	}
}

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
