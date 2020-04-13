package main

// log transfer
// 将日志数据从kafka取出来发往ES

import (
	"fmt"
	"logtransfer/conf"
	"logtransfer/es"
	"logtransfer/kafka"

	"gopkg.in/ini.v1"
)

func main() {
	// 0. 加载配置文件
	var cfg conf.LogTransferCfg
	err := ini.MapTo(&cfg, "./conf/cfg.ini")
	if err != nil {
		panic(err)
	}
	fmt.Println(cfg)

	err = es.Init(cfg.ESCfg.Address, cfg.ESCfg.ChanSize)
	if err != nil {
		panic(err)
	}
	fmt.Println("init es success")

	err = kafka.Init([]string{cfg.KafkaCfg.Address}, cfg.KafkaCfg.Topic)
	if err != nil {
		panic(err)
	}
	fmt.Println("init kafka success")

	select {}
}
