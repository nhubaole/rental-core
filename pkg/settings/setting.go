package settings

import _ "github.com/spf13/viper"

type Config struct {
	DB            Database      `mapstructure:"db"`
	Redis         Redis         `mapstructure:"redis"`
	Server        Server        `mapstructure:"server"`
	S3            AWS           `mapstructure:"s3"`
	JWT           JWT           `mapstructure:"security"`
	NodeServer    NodeServer    `mapstructure:"node_server"`
	Infura        Infura        `mapstructure:"infura"`
	Kafka         Kafka         `mapstructure:"kafka"`
	ElasticSearch ElasticSearch `mapstructure:"elasticsearch"`
	SmartContract SmartContract `mapstructure:"smart_contract"`
}

type Database struct {
	DBPort     int    `mapstructure:"port"`
	DBPassword string `mapstructure:"password"`
	DBHost     string `mapstructure:"host"`
	DBUser     string `mapstructure:"user"`
	DBName     string `mapstructure:"name"`
}

type Server struct {
	Port int    `mapstructure:"port"`
	Host string `mapstructure:"host"`
}

type Redis struct {
	Port int    `mapstructure:"port"`
	Host string `mapstructure:"host"`
}

type AWS struct {
	Region      string `mapstructure:"region"`
	AccessKeyID string `mapstructure:"aws_access_key_id"`
	SecretKey   string `mapstructure:"aws_secret_access_key"`
}

type JWT struct {
	SecretKey string `mapstructure:"secret_key"`
	AESKey    string `mapstructure:"aes_key"`
}

type NodeServer struct {
	Url string `mapstructure:"url"`
}
type Infura struct {
	APIKey string `mapstructure:"api_key"`
}

type Kafka struct {
	Port int `mapstructure:"port"`
}
type ElasticSearch struct {
	Port int `mapstructure:"port"`
}
type SmartContract struct {
	ListingContractAddress         string `mapstructure:"listing_contract"`
	LeaseAgreementProducerContract string `mapstructure:"lease_agreement_producer_contract"`
}
