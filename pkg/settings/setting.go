package settings

import _ "github.com/spf13/viper"

type Config struct {
	DB     Database `mapstructure:"db"`
	Server Server   `mapstructure:"server"`
	S3     AWS      `mapstructure:"s3"`
}

type Database struct {
	DBPort     int    `mapstructure:"port"`
	DBPassword string `mapstructure:"db_password"`
	DBHost     string `mapstructure:"db_host"`
	DBUser     string `mapstructure:"db_user"`
	DBName     string `mapstructure:"db_name"`
}

type Server struct {
	Port int `mapstructure:"port"`
}

type AWS struct {
	Region      string `mapstructure:"region"`
	AccessKeyID string `mapstructure:"aws_access_key_id"`
	SecretKey   string `mapstructure:"aws_secret_access_key"`
}
