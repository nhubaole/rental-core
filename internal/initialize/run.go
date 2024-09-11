package initialize

func Run() {
	LoadConfig()
	InitPostgre()
	InitRedis()
	InitS3()
}
