package distributed

import (
	"fmt"
	"github.com/samuel/go-zookeeper/zk"
	"sync"
	"time"
)

var (
	counter  int64 = 0
)

func incr_counter() {
	c, _, err := zk.Connect([]string{"127.0.0.1"}, time.Second*100)
	if err != nil {
		panic(err)
	}

	lock := zk.NewLock(c, "/lock", zk.WorldACL(zk.PermAll))

	//lock
	lockerr := lock.Lock()
	if lockerr != nil {
		panic(err)
	}
	fmt.Println("lock success!")

	//dosomething
	counter++
	fmt.Println("counter:", counter)
	time.Sleep(time.Millisecond * 1000)
	//unlock
	unlockerr := lock.Unlock()
	if unlockerr != nil {
		panic(err)
	}
	fmt.Println("unlock success!")
}

func ZKLockAPP() {
	var wg sync.WaitGroup
	for i := 0; i < 144; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			incr_counter()
		}()
	}
	wg.Wait()
}
