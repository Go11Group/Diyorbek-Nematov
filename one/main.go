package main

import "fmt"

func addDigits(num int) int {
    for num > 9 {
        sum := 0
        for num != 0 {
            remainder := num % 10
            sum += remainder
            num /= 10
        }
        num = sum
    }
    return num
}

func main() {
	fmt.Println(addDigits(38))
	fmt.Println("Chenged from main branch")
}
