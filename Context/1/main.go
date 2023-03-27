package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*6)
	defer cancel()
	BookHotel(ctx)
}

func BookHotel(ctx context.Context) {
	select {
	case <-ctx.Done():
		fmt.Println("Hotel booking  canceled. Timeout reached")
		return
	case <-time.After(5 * time.Second):
		fmt.Println("Hotel booking")
		return
	}
}
