package kafka

// Conf Configuration for connecting to Kakfa
type Conf struct {
	BrokerURL   string `conf:"env:KAFKA_BROKER_URL,required"`
	TopicPrefix string `conf:"env:KAFKA_TOPIC_PREFIX,required"`
	Username    string `conf:"env:KAFKA_USERNAME,required"`
	Password    string `conf:"env:KAFKA_PASSWORD,required"`
	GroupID     string `conf:"env:KAFKA_GROUP_ID"`
	DisableTLS  bool   `conf:"env:KAFKA_DISABLE_TLS"`
	DisableSASL bool   `conf:"env:KAFKA_DISABLE_SASL"`
}
