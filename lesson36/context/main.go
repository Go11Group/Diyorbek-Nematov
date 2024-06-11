package main

import (
	"context"
	"fmt"
	"time"
)

// context.WithDeadlineCause()
/*
	Bu funksiya mavjud kontekstga belgilangan muddatni qo'shadi. 
	Shuningdek, kontekst tugashi uchun sababni ham o'rnatish imkonini beradi. 
	Sabab tugallanganida kontekstda xato sifatida qaytariladi.
*/

func main() {
	ctx := context.Background()
	deadline := time.Now().Add(5 * time.Second)
	ctx, cancel := context.WithDeadlineCause(ctx, deadline, fmt.Errorf("Custom cause"))
	defer cancel()

	select {
	case <-time.After(10 * time.Second):
		fmt.Println("Operatsiya tugadi")
	case <-ctx.Done():
		fmt.Println("Kontekst tugadi:", ctx.Err())
	}
}
