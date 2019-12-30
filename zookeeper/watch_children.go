package zookeeper

import (
	"fmt"
	"time"

	"github.com/samuel/go-zookeeper/zk"
)

func ZkChildWatchTest() {
	fmt.Printf("ZkChildWatchTest")

	var hosts = []string{"localhost:2181"}
	conn, _, err := zk.Connect(hosts, time.Second*5)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	// try create root path
	var root_path = "/test_root"

	// check root path exist
	exist, _, err := conn.Exists(root_path)
	if err != nil {
		fmt.Println(err)
		return
	}

	if !exist {
		fmt.Printf("try create root path: %s\n", root_path)
		var acls = zk.WorldACL(zk.PermAll)
		p, err := conn.Create(root_path, []byte("root_value"), 0, acls)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("root_path: %s create\n", p)
	}

	// try create child node
	cur_time := time.Now().Unix()
	ch_path := fmt.Sprintf("%s/ch_%d", root_path, cur_time)
	var acls = zk.WorldACL(zk.PermAll)
	p, err := conn.Create(ch_path, []byte("ch_value"), zk.FlagEphemeral, acls)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("ch_path: %s create\n", p)

	// watch the child events
	children, s, child_ch, err := conn.ChildrenW(root_path)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("root_path[%s] child_count[%d]\n", root_path, len(children))
	for idx, ch := range children {
		fmt.Printf("%d, %s\n", idx, ch)
	}

	fmt.Printf("watch children result state[%s]\n", ZkStateString(s))

	for {
		select {
		case ch_event := <-child_ch:
			{
				fmt.Println("path:", ch_event.Path)
				fmt.Println("type:", ch_event.Type.String())
				fmt.Println("state:", ch_event.State.String())

				if ch_event.Type == zk.EventNodeCreated {
					fmt.Printf("has node[%s] detete\n", ch_event.Path)
				} else if ch_event.Type == zk.EventNodeDeleted {
					fmt.Printf("has new node[%d] create\n", ch_event.Path)
				} else if ch_event.Type == zk.EventNodeDataChanged {
					fmt.Printf("has node[%d] data changed", ch_event.Path)
				}
			}
		}

		time.Sleep(time.Millisecond * 10)
	}
}
