package main

//-----------------------------------------------------------------------------

import (
	"context"
	"time"

	"go.etcd.io/etcd/clientv3"
)

func romanToInt(s string) int {
	mapRoman := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	var result int
	next := 0
	for i := len(s) - 1; i >= 0; i-- {
		v := mapRoman[s[i]]
		if v < next {
			result -= v
		} else {
			result += v
		}
		next = v
	}
	return result
}

func main() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		panic(err)
	}
	defer cli.Close()

	// ch := cli.Watch(context.Background(), "testKey")
	// if err != nil {
	// 	panic(err)
	// }

	// for resp := range ch {
	// 	for _, ev := range resp.Events {
	// 		fmt.Printf("type:%v, key:%v, value:%v\n", ev.Type, string(ev.Kv.Key), string(ev.Kv.Value))
	// 	}
	// }

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	// _, err = cli.Put(ctx, "/logagent/192.168.5.6/collec_config", `[{"path":"/home/sxt/nginx.log","topic":"web_log"}]`)
	_, err = cli.Put(ctx, "/logagent/192.168.5.6/collec_config", `[{"path":"/home/sxt/nginx.log","topic":"web_log"},{"path":"/home/sxt/redis.log","topic":"redis_log"},{"path":"/home/sxt/mysql.log","topic":"mysql_log"}]`)
	cancel()
	if err != nil {
		panic(err)
	}

	// // fmt.Println(resp)
	// ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	// resp, err := cli.Get(ctx, "sample_key")
	// cancel()
	// if err != nil {
	// 	panic(err)
	// }

	// for _, ev := range resp.Kvs {
	// 	fmt.Printf("key:%v, value:%v\n", string(ev.Key), string(ev.Value))
	// }
}
