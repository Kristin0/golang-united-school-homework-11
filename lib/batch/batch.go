package batch

import (
	"sync"
	"time"
)

type user struct {
	ID int64
}

func getOne(id int64) user {
	time.Sleep(time.Millisecond * 100)
	return user{ID: id}
}

func getBatch(n int64, pool int64) (res []user) {
	res = make([]user, n)
	syn := make(chan struct{}, pool)

	var wg sync.WaitGroup

	for i := int64(0); i < n; i++ {
		syn <- struct{}{}
		wg.Add(1)
		go func(j int64) {
			res[j] = getOne(j)
			<-syn
			wg.Done()
		}(i)
	}
	wg.Wait()
	return res
}
