package main

import (
	"flag"
	"fmt"
	"log"
	"runtime"
	"strconv"
	"sync"
	"sync/atomic"
	"time"

	"github.com/zeromicro/go-zero/core/limit"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

const seconds = 5

var (
	rdx     = flag.String("redis", "192.168.188.101:6379", "the redis, default localhost:6379")
	rdxKey  = flag.String("redisKey", "rate", "the redis key, default rate")
	rdxPass = flag.String("redisPass", "redisPassword", "the redis password")
	threads = flag.Int("threads", runtime.NumCPU(), "the concurrent threads, default to cores")
)

func main() {
	flag.Parse()

	store := redis.New(*rdx, redis.WithPass(*rdxPass))
	fmt.Println(store.Ping())
	// 实际上只能针对每个
	lmt := limit.NewPeriodLimit(1, 101, store, *rdxKey)
	timer := time.NewTimer(time.Second * seconds)
	quit := make(chan struct{})
	defer timer.Stop()
	go func() {
		<-timer.C
		close(quit)
	}()

	var allowed, denied int32
	var wait sync.WaitGroup
	now := time.Now()
	for i := 0; i < *threads; i++ {

		wait.Add(1)
		go func(i int) {
			for {
				select {
				case <-quit:
					wait.Done()
					return
				default:
					if v, err := lmt.Take(strconv.FormatInt(int64(i), 10)); err == nil && v == limit.Allowed {
						atomic.AddInt32(&allowed, 1)
					} else if err != nil {
						log.Fatal(err)
					} else {
						atomic.AddInt32(&denied, 1)
					}
				}
			}
		}(i)
	}

	wait.Wait()
	fmt.Printf("allowed: %d, denied: %d, qps: %d,time cost:%v\n", allowed, denied, (allowed+denied)/seconds, time.Since(now))
}
