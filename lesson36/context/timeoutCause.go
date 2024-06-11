package main

import (
	"context"
	"fmt"
	"time"
)

//context.WithTimeoutCause

/*
	Bu funksiya context.WithDeadlineCause kabi ishlaydi, faqat muddat o'rniga vaqt oralig'ini belgilaydi.
	Shuningdek, kontekstni tugatish sababini ham qo'shadi.
*/

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeoutCause(ctx, 5*time.Second, fmt.Errorf("Timeout cause"))
	defer cancel()

	select {
	case <-time.After(10 * time.Second):
		fmt.Println("Operatsiya tugadi")
	case <-ctx.Done():
		fmt.Println("Kontekst tugadi:", ctx.Err())
	}
}
