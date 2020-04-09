package conf

type AppConf struct {
	KafkaConf `ini:"kafka"`
	// TaillogConf `ini:"taillog"`
	EtcdConf `ini:"etcd"`
}

type EtcdConf struct {
	Address string `ini:"address"`
	Timeout int    `ini:"timeout"`
	Key     string `ini:"collectLogKey"`
}

type KafkaConf struct {
	Address string `ini:"address"`
	Topic   string `ini:"topic"`
}

type TaillogConf struct {
	FileName string `ini:"filename"`
}
