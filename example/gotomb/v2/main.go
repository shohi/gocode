package main

import (
	"errors"
	"fmt"
	"time"

	"gopkg.in/tomb.v2"
)

func main() {
	var tb tomb.Tomb
	tb.Go(func() error {
		count := 0
		for {
			select {
			case <-tb.Dying():
				return tomb.ErrDying
			default:
				fmt.Printf("hello: %v\n", count)
				time.Sleep(100 * time.Millisecond)
				count++
			}
		}

	})

	time.Sleep(2 * time.Second)
	tb.Kill(errors.New("main error"))
	err := tb.Wait().Error()

	fmt.Printf("====> tomb error: %v\n", err)

}
