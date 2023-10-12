package main

import (
	"fmt"
	"log"
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
			//log.Println(h)
			h.Hit++
			h.Name = Name
			fmt.Println(h.Name, "= Hit", h.Hit)

			if v%setPoint == 0 {
				log.Println(Name, "fail in value", v)
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
		case d := <-Done:
			fmt.Println(d.Name, "kalah, total hit: ", d.Hit)
			return
		}
	}
}
