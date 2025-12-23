package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	cond := sync.NewCond(&sync.Mutex{})
	playersInGame := 5
	cancel := false
	go waitingForPlayers(cond, &cancel)
	for playerId := 0; playerId < 4; playerId++ {
		go playerHandler(cond, &playersInGame, playerId, &cancel)
		time.Sleep(1 * time.Second)
	}
	time.Sleep(3 * time.Second)
}

func playerHandler(cond *sync.Cond, playersRemaining *int, playerId int, cancel *bool) {
	cond.L.Lock()
	fmt.Println(playerId, ": Connected")
	*playersRemaining--
	if *playersRemaining == 0 {
		cond.Broadcast()
	}
	for *playersRemaining > 0 && !*cancel {
		fmt.Println(playerId, ": Waiting for more players")
		cond.Wait()
	}
	cond.L.Unlock()
	if *cancel {
		fmt.Println(playerId, ": Game cancelled due to timeout")
	} else {
		fmt.Println(playerId, ": All players connected. Ready player")
	}
}

func waitingForPlayers(cnd *sync.Cond, cancel *bool) {
	time.Sleep(1 * time.Second)
	cnd.L.Lock()
	*cancel = true
	cnd.Broadcast()
	cnd.L.Unlock()
}
