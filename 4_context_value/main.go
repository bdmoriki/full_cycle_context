package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "token", "123")
	bookHotel(ctx, "Bru")
}

func bookHotel(ctx context.Context, nome string) {
	token := ctx.Value("token")
	fmt.Printf("O token do hotel %s é %s \n", nome, token)
}
