package main

import (
	"fmt"

	"golang.org/x/net/context"
)

func main() {
	ctx, doCancelFunc := context.WithCancel(context.WithValue(context.Background(), "key", "value"))
	defer doCancelFunc()

	fmt.Println(ctx)
}
