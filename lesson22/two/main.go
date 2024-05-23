package main

import "fmt"

func moveZeroes(nums []int) {
    writePointer := 0

    for i := 0; i < len(nums); i++ {
        if nums[i] != 0 {

            nums[writePointer], nums[i] = nums[i], nums[writePointer]
            writePointer++
        }
    }
}

func addStrings(num1 string, num2 string) string {
    dic := map[rune]int{
        '0': 0, '1': 1, '2': 2, '3': 3, '4': 4, '5': 5, '6': 6, '7': 7, '8': 8, '9': 9,
    }

    ans := ""
    carry := 0

    for len(num1) > 0 || len(num2) > 0 || carry > 0 {
        sum := 0
        if len(num1) > 0 {
            digit1 := dic[rune(num1[len(num1)-1])]
            num1 = num1[:len(num1)-1]
            sum += digit1
        }
        if len(num2) > 0 {
            digit2 := dic[rune(num2[len(num2)-1])]
            num2 = num2[:len(num2)-1]
            sum += digit2
        }
        
        sum += carry
        digit := sum % 10
        carry = sum / 10

        ans = string(digit+'0') + ans
    }

    return ans
}

func main() {
    num1 := "123"
    num2 := "456"
    result := addStrings(num1, num2)
    fmt.Println(result)

    nums := []int{0, 1, 0, 3, 12}
    moveZeroes(nums)
    fmt.Println(nums)
}
