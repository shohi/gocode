package basic

import (
	"fmt"
	"testing"
	"time"

	"golang.org/x/net/context"
)

type key string

func TestContext(t *testing.T) {
	keyStr := key("key")
	valueCtx := context.WithValue(context.Background(), keyStr, "value")

	dlCtx, dlCancelFunc := context.WithTimeout(valueCtx, 10*time.Second)
	defer dlCancelFunc()

	ctx, doCancelFunc := context.WithCancel(dlCtx)
	defer doCancelFunc()

	fmt.Println(ctx)
}
