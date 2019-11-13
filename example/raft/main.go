package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

const (
	DefaultHTTPAddr = ":9010"
	DefaultRaftAddr = ":9011"
)

const (
	retainSnapshotCount = 2
	raftTimeout         = 10 * time.Second
)

var (
	httpAddr string
	raftAddr string
	joinAddr string
	nodeID   string
	raftDir  string
)

func init() {
	flag.StringVar(&httpAddr, "haddr", DefaultHTTPAddr, "Set the HTTP bind address")
	flag.StringVar(&raftAddr, "raddr", DefaultRaftAddr, "Set Raft bind address")
	flag.StringVar(&joinAddr, "join", "", "Set join address, if any")
	flag.StringVar(&nodeID, "id", "", "Node ID")
	flag.StringVar(&raftDir, "dir", "", "raft storage directory")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options] \n", os.Args[0])
		flag.PrintDefaults()
	}
}

func main() {
	flag.Parse()

	if raftDir == "" {
		fmt.Fprintf(os.Stderr, "No Raft storage directory specified\n")
		os.Exit(1)
	}

	os.MkdirAll(raftDir, 0700)

	s := newStore()

	s.RaftDir = raftDir
	s.RaftBind = raftAddr

	if err := s.Open(joinAddr == "", nodeID); err != nil {
		log.Fatalf("%v", err)
	}
	s.start()

	service := newService(httpAddr, *s)
	if err := service.Start(); err != nil {
		log.Fatalf("failed to start HTTP service: %s", err.Error())
	}

	if joinAddr != "" {
		if err := join(joinAddr, raftAddr, nodeID); err != nil {
			log.Fatalf("failed to join node at %s: %s", joinAddr, err.Error())
		}
	}

	log.Println("started successfully ...")

	terminate := make(chan os.Signal, 1)
	signal.Notify(terminate, os.Interrupt)
	<-terminate
	log.Println("exiting ...")
}

func join(joinAddr, raftAddr, nodeID string) error {
	queryParams := fmt.Sprintf("addr=%s&id=%s", raftAddr, nodeID)
	url := fmt.Sprintf("http://%s/join?%s", joinAddr, queryParams)
	log.Printf("join ====> %v\n", url)
	resp, err := http.Post(url, "", nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}
