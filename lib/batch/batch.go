package batch

import (
	"sync"
	"time"

)

type user struct {
	ID int64
}

var mu sync.Mutex
func getOne(id int64) user {
	time.Sleep(time.Millisecond * 100)
	return user{ID: id}
}

func getBatch(n int64, pool int64) (res []user) {
	var wg sync.WaitGroup
	var items struct{
		users []user
		mu sync.Mutex
	}
	sem:=make(chan struct{},pool)
	for i := 0; i < int(n); i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			sem<-struct{}{}
			u := getOne(int64(i))
			<-sem
			//fmt.Print("user - ",u)
			items.mu.Lock()
			items.users= append(items.users,u)
			items.mu.Unlock()
		}(i)

	}
	wg.Wait()
	return items.users
}
