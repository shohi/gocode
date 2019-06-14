package time_test

import (
	"context"
	"log"
	"testing"
	"time"
)

func TestTimeoutContext(t *testing.T) {
	ctx, _ := context.WithTimeout(context.Background(), 1*time.Second)

	go func() {
		<-ctx.Done()
		log.Printf("context done without cancel")
	}()

	time.Sleep(2 * time.Second)
}
