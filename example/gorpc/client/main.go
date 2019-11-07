package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net/rpc"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/shohi/gocode/example/gorpc/core"
	"golang.org/x/sync/errgroup"
)

func doRequest(ctx context.Context) error {
	client, err := rpc.Dial("tcp", "localhost:6001")
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	in := bufio.NewReader(os.Stdin)
	for {
		// check context firstly.
		select {
		case <-ctx.Done():
			log.Println("===> context done")
			time.Sleep(5 * time.Second)
			return ctx.Err()
		default:
			// do nothing
		}

		name, _, err := in.ReadLine()
		if err != nil {
			log.Println("===> read line error")
			return err
		}
		req := &core.Request{Name: string(name)}
		var resp core.Response
		err = client.Call(core.HandlerName, req, &resp)
		if err != nil {
			log.Printf("===> response error, err: [%+v], req: [%+v]\n", err, req)
			return err
		}

		fmt.Printf("response===> %v\n", resp.Message)
	}
}

// monitor watchs on signals (SIGINT/SIGTERM) to exit.
func monitor(ctx context.Context) error {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	select {
	case sig := <-signals:
		return fmt.Errorf("signal %v tripped", sig)
	case <-ctx.Done():
		return ctx.Err()
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	eg, egCtx := errgroup.WithContext(ctx)

	// NOTE: both these two goroutines will accept control chars from stdin.
	eg.Go(func() error { return doRequest(egCtx) })
	eg.Go(func() error { return monitor(egCtx) })

	// TODO: Wait all goroutines to shutdown.
	if err := eg.Wait(); err != nil {
		fmt.Printf("error in the group goroutines: %s\n", err)
		os.Exit(1)
	}
	fmt.Println("everything closed successfully")
}
