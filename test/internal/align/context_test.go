package align_test

import (
	"context"
	"log"
	"testing"
	"time"
	"unsafe"

	"go.uber.org/goleak"
)

func TestContextDoneAlign(t *testing.T) {
	defer goleak.VerifyNone(t)

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	ch := ctx.Done()

	log.Printf("context done channel - size: %v, align %v", unsafe.Sizeof(ch), unsafe.Alignof(ch))
}
