package configs

var Configs = map[string]AppConfig{
	"local-qa": {
		"mongodb://localhost:27017",
		"CustomersDb",
		"Customers",
	},
}
