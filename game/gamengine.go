package game

import (
	"fmt"
	"sync"
)

type UserResponse struct {
	UserId    int  `json:"user_id"`
	IsCorrect bool `json:"is_correct"`
}

type GameEngine struct {
	winnerUserId         int
	winnerDeclared       bool
	correctAnswerCount   int
	inCorrectAnswerCount int
	mu                   sync.Mutex
	Notify               chan UserResponse
}

func NewGameEngine() *GameEngine {
	gamengine := &GameEngine{
		winnerUserId: -1,
		Notify:       make(chan UserResponse, 1000),
	}

	go gamengine.listen()

	return gamengine
}

func (g *GameEngine) listen() {

	for resp := range g.Notify {

		g.mu.Lock()

		if g.winnerDeclared {
			if resp.IsCorrect {
				g.correctAnswerCount++
			} else {
				g.inCorrectAnswerCount++
			}
			g.mu.Unlock()
			continue
		}

		if resp.IsCorrect {
			g.winnerUserId = resp.UserId
			g.winnerDeclared = true
			fmt.Printf("---> Winner's user id is %d\n", resp.UserId)
			g.correctAnswerCount++
		} else {
			g.inCorrectAnswerCount++
		}
		g.mu.Unlock()
	}
}

func (g *GameEngine) Metrics() (int, int, int) {
	g.mu.Lock()
	defer g.mu.Unlock()
	return g.winnerUserId, g.correctAnswerCount, g.inCorrectAnswerCount
}

func (g *GameEngine) Close() {
	close(g.Notify)
}
