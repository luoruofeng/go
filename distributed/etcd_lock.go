package distributed

import (
	"fmt"
	"sync"

	"github.com/zieckey/etcdsync"
)

//export GOPROXY=https://goproxy.io

var (
	etcdcounter int64 = 0
)

func incrCounter() {
	m, err := etcdsync.New("/lock", 10, []string{"127.0.0.1"})
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	//lock
	lockerr := m.Lock()
	if lockerr != nil {
		fmt.Println(lockerr)
		panic(err)
	}
	fmt.Println("lock success!")

	//dosomething
	etcdcounter++
	fmt.Println("counter:", etcdcounter)

	//unlock
	unlockerr := m.Unlock()
	if unlockerr != nil {
		panic(err)
	}
	fmt.Println("unlock success!")
}

func EtcdApp() {
	var wg sync.WaitGroup
	for i := 0; i < 14; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			incrCounter()
		}()
	}
	wg.Wait()
}
