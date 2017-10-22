package main

import (
	"fmt"
	"time"

	"golang.org/x/net/context"
)

func main() {
	keyStr := "key"
	valueCtx := context.WithValue(context.Background(), &keyStr, "value")
	dlCtx, dlCancelFunc := context.WithTimeout(valueCtx, 10*time.Second)
	defer dlCancelFunc()
	ctx, doCancelFunc := context.WithCancel(dlCtx)
	defer doCancelFunc()

	fmt.Println(ctx)
}
