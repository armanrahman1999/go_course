package main

import (
	"context"
	"fmt"
)

func main() {

	todoContext := context.TODO()

	contextBkg := context.Background()

	ctx := context.WithValue(todoContext, "Name", "John")
	fmt.Println(ctx)
	fmt.Println(ctx.Value("Name"))

	ctx1 := context.WithValue(contextBkg, "city", "New York")
	fmt.Println(ctx1)
	fmt.Println(ctx1.Value("city"))
}