package main

import (
	"fmt"
	girdled "github.com/go-redis/redis/v7"
	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v7"
	"sync"
	"time"
)

func main() {
	// Create a pool with go-redis (or redigo) which is the pool redisync will
	// use while communicating with Redis. This can also be any pool that
	// implements the `redis.Pool` interface.

	client := girdled.NewClient(&girdled.Options{
		Addr: "127.0.0.1:6379",
	})
	pool := goredis.NewPool(client) // or, pool := redigo.NewPool(...)

	// Create an instance of redisync to be used to obtain a mutual exclusion
	// lock.
	rs := redsync.New(pool)

	// Obtain a new mutex by using the same name for all instances wanting the
	// same lock.

	//
	gNum:=2
	mutexname := "421"
	var wg sync.WaitGroup
	wg.Add(gNum)
	for i:=0;i<gNum;i++{
		go func() {
			defer wg.Done()
			//
			mutex := rs.NewMutex(mutexname)
			fmt.Printf("开始获取redis锁%d\n",i)
			// Obtain a lock for our given mutex. After this is successful, no one else
			// can obtain the same lock (the same mutex name) until we unlock it.
			if err := mutex.Lock(); err != nil {
				panic(err)
			}

			// Do your work that requires the lock.
			fmt.Printf("获取redis锁suc%d\n",i)
			time.Sleep(time.Second*2)
			fmt.Printf("释放redis锁%d\n",i)
			// Release the lock so other processes or threads can obtain a lock.
			if ok, err := mutex.Unlock(); !ok || err != nil {
				panic("unlock failed")
			}
			fmt.Printf("释放redis锁%dsuc\n",i)
		}()
	}
	wg.Wait()


}

