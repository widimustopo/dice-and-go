package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	OneValuePassDice    = 1
	SixValueGetOnePoint = 6
)

type Players struct {
	player      string
	dicesInHand []int
	point       int
	movingDice  int
}

type FinalPoint struct {
	player string
	point  int
}

var onlyOnce sync.Once

// prepare the dice
var dice = []int{1, 2, 3, 4, 5, 6}

func rollDice() int {

	onlyOnce.Do(func() {
		rand.Seed(time.Now().UnixNano()) // only run once
	})

	return dice[rand.Intn(len(dice))]
}

func countProcess(players []Players, newFinalPoint []FinalPoint, try int) ([]Players, []FinalPoint, int) {

	var playerInGame []Players

	resPoint := checkPoint(players)
	playerInGame = resPoint
	fmt.Println("Get 1 point if value is 6 :", playerInGame)
	result, finalPointsFromResPoint := checkGameOver(resPoint, newFinalPoint)
	for _, resPoint := range finalPointsFromResPoint {
		newFinalPoint = append(newFinalPoint, FinalPoint{
			player: resPoint.player,
			point:  resPoint.point,
		})
	}
	playerInGame = result

	movedice := moveDice(playerInGame)
	playerInGame = movedice
	fmt.Println("move die if value is  1   :", playerInGame)

	result, finalPointsMoving := checkGameOver(playerInGame, newFinalPoint)
	for _, resPoint := range finalPointsMoving {
		newFinalPoint = append(newFinalPoint, FinalPoint{
			player: resPoint.player,
			point:  resPoint.point,
		})
	}

	playerInGame = result

	if len(playerInGame) != 0 && len(playerInGame) != 1 {
		for _, new := range playerInGame {
			for i := 0; i < len(new.dicesInHand); i++ {
				new.dicesInHand[i] = rollDice()
			}
		}
		fmt.Printf("after new roll #%v         : %v", try+1, playerInGame)
		fmt.Println()
	}

	return playerInGame, newFinalPoint, try + 1

}

func checkPoint(players []Players) []Players {
	var newPlayerDice []Players

	for _, player := range players {

		for i := 0; i < len(player.dicesInHand); i++ {

			if player.dicesInHand[i] == SixValueGetOnePoint {
				player.point += 1
				player.dicesInHand = append(player.dicesInHand[:i], player.dicesInHand[i+1:]...)
				i--
				continue
			}

			if player.dicesInHand[i] == OneValuePassDice {
				player.movingDice += 1
				player.dicesInHand = append(player.dicesInHand[:i], player.dicesInHand[i+1:]...)
				i--
				continue
			}

		}

		newPlayerDice = append(newPlayerDice, Players{
			player:      player.player,
			dicesInHand: player.dicesInHand,
			point:       player.point,
			movingDice:  player.movingDice,
		})
	}

	return newPlayerDice
}

func moveDice(players []Players) []Players {

	if len(players) != 1 && len(players) != 0 {
		for i := 0; i < len(players); i++ {
			if players[i].movingDice != 0 {
				if i == len(players)-1 {
					for j := 0; j < players[i].movingDice; j++ {
						players[0].dicesInHand = append(players[0].dicesInHand, 1)
					}
				} else {
					for j := 0; j < players[i].movingDice; j++ {
						players[i+1].dicesInHand = append(players[i+1].dicesInHand, 1)
					}
				}
			}
			players[i].movingDice = 0
		}
	}

	return players
}

func checkGameOver(players []Players, finals []FinalPoint) ([]Players, []FinalPoint) {
	var newFinalPoint []FinalPoint

	for i := 0; i < len(players); i++ {
		if len(players[i].dicesInHand) == 0 {

			newFinalPoint = append(newFinalPoint, FinalPoint{
				player: players[i].player,
				point:  players[i].point,
			})

			players = append(players[:i], players[i+1:]...)
			i--
			continue
		}
	}

	return players, newFinalPoint
}

func displayWinner(finalPoint []FinalPoint) {
	fmt.Println("===========================")
	fmt.Println("Final Result              :", finalPoint)
	biggestPoint := finalPoint[0]
	for _, win := range finalPoint {
		if win.point > biggestPoint.point {
			biggestPoint = win
		}
	}
	fmt.Println("===========================")
	fmt.Printf("Winner is %v , with points of %v", biggestPoint.player, biggestPoint.point)
	fmt.Println()
	fmt.Println("===========================")
}

func start(player int, dice int) []Players {
	var newPlayer []Players

	//startig player
	for i := 0; i < player; i++ {
		var hisDice []int

		//starting dice
		for i := 0; i < dice; i++ {
			hisDice = append(hisDice, rollDice())
		}
		newPlayer = append(newPlayer, Players{
			player:      fmt.Sprintf("Player #%v", i+1),
			dicesInHand: hisDice,
			point:       0,
			movingDice:  0,
		})
	}

	return newPlayer
}

func main() {
	player := 3
	dice := 4

	initPlayers := start(player, dice)
	fmt.Println("after roll #1             :", initPlayers)
	var resPlayer []Players
	var finalPoint []FinalPoint
	var tried int
	resLeft, _, count := countProcess(initPlayers, nil, 1)
	tried = count
	resPlayer = resLeft
	for len(finalPoint) != player && len(resPlayer) != 0 && len(resPlayer) != 1 {
		result, final, count := countProcess(resPlayer, finalPoint, tried)
		resPlayer = result
		finalPoint = final
		tried = count
	}

	if len(resPlayer) > 0 {
		for _, last := range resPlayer {
			finalPoint = append(finalPoint, FinalPoint{
				player: last.player,
				point:  last.point,
			})
		}

	}

	displayWinner(finalPoint)

}
