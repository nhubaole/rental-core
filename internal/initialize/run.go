package initialize

import "smart-rental/pkg/blockchain"

func Run() {
	LoadConfig()
	InitPostgre()
	InitRedis()
	InitS3()
	InitElasticsearch()
	ESCreateIndexIfNotExist()
	blockchain.InitEthClient()
}
