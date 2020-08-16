package model

// model
type Config struct {
	KafkaConf
	ESConf
}

type KafkaConf struct {
	Address string `init:"address"`
	Topic   string `ini:"topic"`
}

type ESConf struct {
	Address string `ini:"address"`
	Index   string `ini:"index"`
}
