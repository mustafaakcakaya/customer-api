package configs

var ac AppConfig

type AppConfig struct {
	MongoConnectionURI     string
	CustomerDBName         string
	CustomerCollectionName string
}

func SetConfigs(env string) {
	ac = AppConfig{
		MongoConnectionURI:     Configs[env].MongoConnectionURI,
		CustomerDBName:         Configs[env].CustomerDBName,
		CustomerCollectionName: Configs[env].CustomerCollectionName,
	}
}

func GetConfigs() *AppConfig {
	return &ac
}
