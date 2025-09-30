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

func SimulateUsers(n int, serverUrl string) {

	var wg sync.WaitGroup

	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(uid int) {
			defer wg.Done()
			randelay := rand.Intn(990) + 10
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
