package batch

import (
	"context"
	"golang.org/x/sync/errgroup"
	"sync"
	"time"
)

type user struct {
	ID int64
}


func getOne(id int64) (user) {
	time.Sleep(time.Millisecond * 100)
	return user{ID: id}
}

//func getBatch(n int64, pool int64) (res []user) {
//	var wg sync.WaitGroup
//	var items struct{
//		users []user
//		mu sync.Mutex
//	}
//	sem:=make(chan struct{},pool)
//	for i := 0; i < int(n); i++ {
//		wg.Add(1)
//		go func(i int) {
//			defer wg.Done()
//			sem<-struct{}{}
//			u := getOne(int64(i))
//			<-sem
//			//fmt.Print("user - ",u)
//			items.mu.Lock()
//			items.users= append(items.users,u)
//			items.mu.Unlock()
//		}(i)
//
//	}
//	wg.Wait()
//	return items.users
//}


func getBatch(n int64, pool int64) (res []user) {
	var items struct{
		users []user
		mu sync.Mutex
	}
	errG,_:=errgroup.WithContext(context.Background())
	errG.SetLimit(int(pool))
	for i := 0; i < int(n); i++ {
		func(i int) {
			errG.Go(func() error {
				u := getOne(int64(i))
				items.mu.Lock()
				items.users = append(items.users, u)
				items.mu.Unlock()
				return nil
			})
		}(i)
	}
	err:=errG.Wait()
	if err==nil{
		return items.users
	}
	return nil
}
