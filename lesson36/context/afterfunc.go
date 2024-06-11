package main

import (
	"fmt"
	"time"
)

// time.AfterFunc

/*
	Bu funksiya ma'lum vaqt oralig'idan keyin berilgan funksiyani ishga tushiradi.
*/

func main() {
	time.AfterFunc(3*time.Second, func() {
		fmt.Println("3 soniyadan keyin ishga tushdi")
	})

	time.Sleep(5 * time.Second)
}
