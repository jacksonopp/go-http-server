package main

import (
	"log"
	"net/http"
)

type InMemoryPlayerStore struct{}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return 134
}

func (i *InMemoryPlayerStore) RecordWin(name string) {}

func main() {
	server := &PlayerServer{store: &InMemoryPlayerStore{}}
	log.Fatal(http.ListenAndServe(":5000", server))
}
