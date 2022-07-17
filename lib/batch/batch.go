package batch

import (
	"time"
	"sync"
)

type user struct {
	ID int64
}

func getOne(id int64) user {
	time.Sleep(time.Millisecond * 100)
	return user{ID: id}
}

func getBatch(n int64, pool int64) (res []user) {
	const paralel = 5
	var wg sync.WaitGroup
	sem := make(chan struct{}, paralel)
	for i := 0; int64(i) < n; i++ {
		wg.Add(1)
		sem <- struct{}{}
		go func (){
			for j := 0; int64(j) < n; j++ {
				getOne(int64(j))
			}
		  <- sem
		  wg.Done()
		}()
	}
	wg.Wait()
	return nil
}
