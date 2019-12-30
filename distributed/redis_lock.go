package distributed

import (
	"fmt"
	"github.com/go-redis/redis"
	"sync"
	"time"
)

//docker run --name some-redis -p 6379:6379 -d redis redis-server --appendonly yes

const (
	counterkey = "counterkey"
	lockkey    = "lockkey"
)

func incr() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // use default Addr
		Password: "",               // no password set
		DB:       0,                // use default DB
	})

	//lock
	resp := rdb.SetNX(lockkey, 1, time.Second*5)
	lockSuccess, err := resp.Result()
	if err != nil || !lockSuccess {
		fmt.Println(err, "lock result:", lockSuccess)
		return
	} else {
		fmt.Println("locked!!")
	}

	//counter
	val := rdb.Get(counterkey)
	v, err := val.Int64()
	if err == nil {
		v++
	} else {
		v = 0
	}
	counterresp := rdb.Set(counterkey, v, 0)
	_, countererr := counterresp.Result()
	if countererr != nil {
		println("set value error!")
	}
	println("counter:", v)

	unlockval := rdb.Del(lockkey)
	unlocksuccess, err := unlockval.Result()

	if err == nil && unlocksuccess > 0 {
		println("unlock success")
	} else {
		println("unlock err , ", err)
	}
}

func RedisLockApp() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			incr()
		}()
	}
	wg.Wait()
}
