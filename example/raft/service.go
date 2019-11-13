package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type Service struct {
	addr  string
	store Store
}

func newService(addr string, store Store) *Service {
	return &Service{
		addr:  addr,
		store: store,
	}
}

func (s *Service) Start() error {
	router := mux.NewRouter()

	router.HandleFunc("/join", s.handleJoin)
	router.HandleFunc("/apply", s.handleApply)
	router.HandleFunc("/leader", s.handleLeader)
	router.HandleFunc("/leave", s.handleLeave)

	go func() {
		log.Printf("listening on %s ...", s.addr)
		http.ListenAndServe(s.addr, router)
	}()
	return nil
}
func (s *Service) handleApply(w http.ResponseWriter, r *http.Request) {
	msg := r.URL.Query().Get("msg")

	err := s.store.Apply([]byte(msg), 2*time.Second)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (s *Service) handleLeader(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(strconv.Itoa(s.store.IsLeader())))
}

func (s *Service) handleLeave(w http.ResponseWriter, r *http.Request) {
	err := s.store.Leave()
	if err != nil {
		msg := fmt.Sprintf("failed to leave, err: %v", err)
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Leave successfully"))
}

func (s *Service) handleJoin(w http.ResponseWriter, r *http.Request) {
	m := r.URL.Query()
	remoteAddr := m.Get("addr")
	if len(remoteAddr) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	nodeID := m.Get("id")
	if len(nodeID) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := s.store.Join(nodeID, remoteAddr); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
