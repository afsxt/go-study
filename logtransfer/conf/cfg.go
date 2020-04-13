package conf

// LogTransfer 配置
type LogTransferCfg struct {
	KafkaCfg `ini:"kafka"`
	ESCfg    `ini:"es"`
}

// KafkaCfg kafka
type KafkaCfg struct {
	Address string `ini:"address"`
	Topic   string `ini:"topic"`
}

// ESCfg es
type ESCfg struct {
	Address  string `ini:"address"`
	ChanSize int    `ini:"chansize"`
}
