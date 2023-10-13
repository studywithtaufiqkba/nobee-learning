package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
    PINGPONG APPS
    2 player => 2 goroutine
    kondisi kalah : saat flag/counter/random number, habis dibagi 11 (random % 11 == 0 THAN break)


    Contoh :
    Player A = Hit 1 // counter 8
    Player B = Hit 2 // counter 3
    Player A = Hit 3 // counter 24
    Player B = Hit 4 // counter 33
    Player B kalah, total hit : 4, kalah di nomor 33

   Player A = Hit 1 // counter 8
   Player B = Hit 2 // counter 9
   Player A = Hit 3 // counter 22

   Player A kalah, total hit : 3, kalah di nomor 22

   Requirement :
   - Struct Player {
       Name string
       Hit int
   }

   a := Player{Name: A, Hit:0}
   b := Player{Name: B, Hit:0}

   a.Hit++
   b.Hit++

*/

const setPoint = 11

func main() {
	player := make(chan *Player)
	done := make(chan *Player)

	players := []string{"Player A", "Player B", "Player C", "Player D"}
	for _, p := range players {
		go PlayPingPong(p, player, done)
	}
	player <- new(Player)

	finish(done)
	//	Output example
	// Player D lost, total hit:  13 with Value:  99
}

type Player struct {
	Name string
	Hit  int
}

func PlayPingPong(Name string, Hit, Done chan *Player) {
	rand.Seed(time.Now().UnixNano())
	for {
		select {
		case h := <-Hit:

			v := rand.Intn(100-1) + 1
			h.Hit++
			h.Name = Name
			fmt.Println(h.Name, "= Hit", h.Hit, "with Value: ", v)

			if v%setPoint == 0 {
				fmt.Println(h.Name, "lost, total hit: ", h.Hit, "with value: ", v)
				Done <- h
				break
			}
			Hit <- h
		}
	}
}

func finish(Done chan *Player) {
	for {
		select {
		case _ = <-Done:
			return
		}
	}
}
