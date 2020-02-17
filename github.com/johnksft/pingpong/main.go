package main

import (
	"fmt"
	"math/rand"
)

const MAX_POINT = 7
const REL_POINT = 6

func isPoint() int {
	isPoint := rand.Intn(REL_POINT)
	switch isPoint {
	case 1:
		return 1
	default:
		return 0
	}
}

func score(pointsA *int, pointsB *int) {
	fmt.Print("Score: Ping: ", *pointsA, " VS Pong: ", *pointsB, "\n")
}

func ping(ball chan<- int, action chan<- string, player string ) {
	ball <- 1
	action <- "ping "+player
}

func pong(ball chan<- int, action chan<- string, player string) {
	ball <- 2
	action <- "pong "+player
}

func referee(action <-chan string, pointsA *int, pointsB *int) {
	for {
		fmt.Println("Action: ", <-action)
	}
}
func pingpong() {
	
	var player1 string
	var player2 string
	fmt.Println("Ingrese nombre de Jugador 1: ")
	fmt.Scanf("%s",&player1)
	fmt.Println("Ingrese nombre de Jugador 2: ")
	fmt.Scanf("%s",&player2)
	ball := make(chan int)
	action := make(chan string)
	pointsA := 0
	pointsB := 0
	go referee(action, &pointsA, &pointsB)
	go ping(ball, action, player1)
	for pointsA < MAX_POINT && pointsB < MAX_POINT {
		value := <-ball
		switch value {
		case 1:
			go pong(ball, action, player2)
			pointsB = pointsB + isPoint()
		case 2:
			go ping(ball, action, player1)
			pointsA = pointsA + isPoint()
		}
	}
	defer func(){
		score(&pointsA, &pointsB)
		if pointsA > pointsB {
			fmt.Printf("%s es el Ganador \n",player1)
		} else {
			fmt.Printf("%s es el Ganador \n",player2)
		}
	}()
}

func main() {
	pingpong()
}
