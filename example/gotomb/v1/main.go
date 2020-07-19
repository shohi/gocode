package main

import (
	"fmt"
	"log"
	"sync"
	"time"

	"gopkg.in/tomb.v1"
)

type foo struct {
	tomb tomb.Tomb
	wg   sync.WaitGroup
}

func (f *foo) task(id int) {
	for i := 0; i < 10; i++ {
		select {
		case <-time.After(1e9):
			log.Printf("task %d tick\n", id)
		case <-f.tomb.Dying():
			log.Printf("task %d stopping\n", id)
			f.wg.Done()
			return
		}
	}
}

func (f *foo) Run() {
	f.wg.Add(10)
	for i := 0; i < 10; i++ {
		go f.task(i)
	}
	go func() {
		f.wg.Wait()
		fmt.Printf("===> waitgroup return\n")
		f.tomb.Done()
	}()
}

func (f *foo) Stop() error {
	f.tomb.Kill(nil)
	return f.tomb.Wait()
}

func main() {
	var f foo
	f.Run()
	time.Sleep(3.5e9)
	// log.Printf("====> start sleep\n")
	// time.Sleep(3*time.Second + 500*time.Millisecond)
	// log.Printf("====> end sleep\n")
	log.Printf("calling stop\n")
	err := f.Stop()
	log.Printf("all done, reason: %v\n", err)
}
