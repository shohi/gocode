package context_test

import (
	"context"
	"testing"

	goctx "golang.org/x/net/context"
)

func golangContext(ctx goctx.Context) {

}

func stdContext(ctx context.Context) {

}

func TestGolangContextWithStd(t *testing.T) {
	ctx := context.Background()
	golangContext(ctx)
}

func TestStdContextWithGolang(t *testing.T) {
	ctx := goctx.Background()
	stdContext(ctx)
}

func TestCtx_Cancel_Idempotent(t *testing.T) {
	_, cancel := context.WithCancel(context.Background())
	cancel()
	cancel()

}
