package main

import (
	"fmt"
	"logagent/conf"
	"logagent/etcd"
	"logagent/kafka"
	"logagent/taillog"
	"time"

	"gopkg.in/ini.v1"
)

var cfg conf.AppConf

func run() {
	// 1. 读取日志
	for {
		select {
		case line := <-taillog.ReadChan():
			// 2. 发送到kafka
			kafka.SendToKafka(cfg.Topic, line.Text)
		default:
			time.Sleep(time.Second)
		}
	}
}

func main() {
	// 0. 加载配置文件
	// cfg, err := ini.Load("./conf/config.ini")
	err := ini.MapTo(&cfg, "./conf/config.ini")
	if err != nil {
		fmt.Println("load ini failed, err: ", err)
	}
	fmt.Println(cfg)

	// 1. 初使化kafka连接
	err = kafka.Init([]string{cfg.KafkaConf.Address})
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
	logEntryConf, err := etcd.GetConf(cfg.EtcdConf.Key)
	if err != nil {
		panic(err)
	}
	fmt.Println("get conf from etcd success")
	for index, value := range logEntryConf {
		fmt.Println(index, value)
	}
	// 2.2 监视日志收集项的变化

	// 2. 打开日志文件准备收集日志
	// err = taillog.Init(cfg.FileName)
	// if err != nil {
	// 	fmt.Println("init taillog failed, err: ", err)
	// 	return
	// }
	// fmt.Println("init taillog success")

	// run()
}
