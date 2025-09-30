package mockuser

import (
	"bytes"
	"encoding/json"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

type SimUser struct {
	UserId    int  `json:"user_id"`
	IsCorrect bool `json:"is_correct"`
}

// function to simulate the mock users
func SimulateUsers(n int, serverUrl string) {

	var wg sync.WaitGroup

	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(uid int) {
			defer wg.Done()
			randelay := rand.Intn(991) + 10
			// fmt.Printf("for %d userid creating %d milliseconds of delay\n", uid, randelay)
			time.Sleep(time.Duration(randelay) * time.Millisecond)
			su := SimUser{
				UserId:    uid,
				IsCorrect: rand.Intn(2) == 0,
			}
			body, _ := json.Marshal(su)
			http.Post(serverUrl+"/submit", "application/json", bytes.NewReader(body))
		}(i)
	}

	wg.Wait()
}
