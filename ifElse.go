package main

import "fmt"

func main()  {
	if 7 % 2 == 0 {
		fmt.Println("7 is even, 7 是偶数")
	}else {
		fmt.Println("7 is odd, 7 是奇数")
	}

	if 8 % 4 == 0{
		fmt.Println("8 is divisible 4")
	}

	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	}else if num < 10 {
		fmt.Println(num, "has i digit")
	}
}
