package settings

import _ "github.com/spf13/viper"

type Config struct {
	DB     Database `mapstructure:"db"`
	Server Server   `mapstructure:"server"`
	S3     AWS      `mapstructure:"s3"`
	JWT    JWT      `mapstructure:"security"`
	NodeServer NodeServer `mapstructure:"node_server"`
}

type Database struct {
	DBPort     int    `mapstructure:"port"`
	DBPassword string `mapstructure:"password"`
	DBHost     string `mapstructure:"host"`
	DBUser     string `mapstructure:"user"`
	DBName     string `mapstructure:"name"`
}

type Server struct {
	Port int `mapstructure:"port"`
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