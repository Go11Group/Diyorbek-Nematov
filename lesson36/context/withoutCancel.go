package main

import (
	"context"
	"fmt"
	"time"
)

// context.WithoutCancel

/*
	Bu funksiya mavjud kontekstdan bekor qilish signalini olib tashlaydi.
	Yangi kontekst asosiy kontekst tugatilganda tugatilmaydi.
*/

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	newCtx := context.WithoutCancel(ctx)

	cancel()

	select {
	case <-newCtx.Done():
		fmt.Println("Yangi kontekst bekor qilindi")
	case <-time.After(2 * time.Second):
		fmt.Println("Yangi kontekst bekor qilinmadi")
	}
}
