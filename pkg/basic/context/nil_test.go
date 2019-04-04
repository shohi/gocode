package context_test

import (
	"context"
	"log"
	"testing"
)

func TestNilContext(t *testing.T) {
	var ctx context.Context
	kCtx := context.WithValue(ctx, "key", "val")
	log.Printf("kCtx: %v", kCtx)

	log.Printf("value: %v", kCtx.Value("key"))
}
