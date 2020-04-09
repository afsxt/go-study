package main

import (
	"fmt"
	"logagent/conf"
	"logagent/etcd"
	"logagent/kafka"
	"logagent/taillog"
	"logagent/utils"
	"sync"
	"time"

	"gopkg.in/ini.v1"
)

var cfg conf.AppConf

// func run() {
// 	// 1. 读取日志
// 	for {
// 		select {
// 		case line := <-taillog.ReadChan():
// 			// 2. 发送到kafka
// 			kafka.SendToKafka(cfg.Topic, line.Text)
// 		default:
// 			time.Sleep(time.Second)
// 		}
// 	}
// }

func main() {
	// 0. 加载配置文件
	// cfg, err := ini.Load("./conf/config.ini")
	err := ini.MapTo(&cfg, "./conf/config.ini")
	if err != nil {
		fmt.Println("load ini failed, err: ", err)
	}
	fmt.Println(cfg)

	// 1. 初使化kafka连接
	err = kafka.Init([]string{cfg.KafkaConf.Address}, cfg.KafkaConf.ChanMaxSize)
	if err != nil {
		fmt.Println("init kafka failed, err: ", err)
		return
	}
	defer kafka.Close()
	fmt.Println("init kafka success")

	// 2. 初使化ETCD
	err = etcd.Init(cfg.EtcdConf.Address, time.Duration(cfg.EtcdConf.Timeout)*time.Second)
	if err != nil {
		panic(err)
	}
	fmt.Println("init etcd success")

	// 2.1从etcd中获取日志收集项的配置信息
	ipStr, err := utils.GetOutboundIP()
	if err != nil {
		panic(err)
	}
	etcdConfKey := fmt.Sprintf(cfg.EtcdConf.Key, ipStr)
	logEntryConf, err := etcd.GetConf(etcdConfKey)
	if err != nil {
		panic(err)
	}
	fmt.Println("get conf from etcd success")
	for index, value := range logEntryConf {
		fmt.Println(index, value)
	}
	// 3 收集日志往kafka
	taillog.Init(logEntryConf)

	newConfChan := taillog.NewConfChan()
	var wg sync.WaitGroup
	wg.Add(1)
	// 监视日志收集项的变化，实现热更新
	go etcd.WatchConf(etcdConfKey, newConfChan)
	wg.Wait()

	// run()
}
