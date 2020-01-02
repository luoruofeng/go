package etcd

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/coreos/etcd/client"
)

//etcdctl get /configs/remote_config.json

//etcdctl -o json get /configs/remote_config.json
// {
// 	"action": "get",
// 	"node": {
// 		"key": "/configs/remote_config.json",
// 		"value": "{\"addr\":\"addr\",\"aes_key\":\"\",\"https\":false,\"secret\":\"\",\"private_key_path\":\"\",\"cert_file_path\":\"\"}",
// 		"nodes": null,
// 		"createdIndex": 47,
// 		"modifiedIndex": 47
// 	},
// 	"prevNode": null
// }

var configPath = `/configs/remote_config.json`
var kapi client.KeysAPI

type ConfigStruct struct {
	Addr           string `json:"addr"`
	AesKey         string `json:"aes_key"`
	HTTPS          bool   `json:"https"`
	Secret         string `json:"secret"`
	PrivateKeyPath string `json:"private_key_path"`
	CertFilePath   string `json:"cert_file_path"`
}

var appConfig ConfigStruct

func Init() {
	cfg := client.Config{
		Endpoints:               []string{"http://127.0.0.1:2379"},
		Transport:               client.DefaultTransport,
		HeaderTimeoutPerRequest: time.Second,
	}

	if c, err := client.New(cfg); err != nil {
		log.Fatal(err)
	} else {
		kapi = client.NewKeysAPI(c)

		cs := &ConfigStruct{
			Addr:           "a",
			AesKey:         "a",
			HTTPS:          false,
			Secret:         "s",
			PrivateKeyPath: "p",
			CertFilePath:   "c",
		}

		b, err := json.Marshal(cs)
		if err != nil {
			log.Fatal(err)
		}
		kapi.Set(context.Background(), configPath, string(b), nil)

		log.Println("init value is: ", string(b))

		// initConfig()
	}
}

// func initConfig() {
// 	if resp, err := kapi.Get(context.Background(), configPath, nil); err != nil {
// 		log.Fatal(err)
// 	} else {
// 		err = json.Unmarshal([]byte(resp.Node.Value), &appConfig)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 	}
// }

func watchAndUpdate() {
	w := kapi.Watcher(configPath, nil)
	go func() {
		for {
			resp, err := w.Next(context.Background())
			if err != nil {
				log.Fatal(err)
			}

			log.Println("new value is: ", resp.Node.Value)

			err = json.Unmarshal([]byte(resp.Node.Value), &appConfig)
			if err != nil {
				log.Fatal(err)
			}

			log.Println("appConfig is: ", appConfig)
		}
	}()
}

func getConfig() ConfigStruct {
	return appConfig
}

func UpdateConfigApp() ConfigStruct {
	//init
	Init()
	watchAndUpdate()

	//change value for test watcher
	time.Sleep(time.Second * 4)
	appConfig.Addr = "addr"
	appConfig.CertFilePath = "CertFilePath"
	b, _ := json.Marshal(appConfig)
	kapi.Set(context.Background(), configPath, string(b), nil)

	//你可以手动修改
	//etcdctl set /configs/remote_config.json \{\"addr\"\:\"addr2\",\"aes_key\"\:\"\",\"https\"\:false,\"secret\"\:\"\",\"private_key_path\"\:\"\",\"cert_file_path\"\:\"CertFilePath3\"\}
	time.Sleep(time.Second * 1000)
	return getConfig()
}
