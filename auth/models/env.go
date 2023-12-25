package models

type Environment struct {
	Datasources struct {
		MongoDB struct {
			Uri string `mapstructure:"uri"`
		} `mapstructure:"mongodb"`
		Redis struct {
			Uri      string `mapstructure:"uri"`
			Password string `mapstructure:"password"`
		} `mapstructure:"redis"`
		Neo4j struct {
			Uri      string `mapstructure:"uri"`
			Username string `mapstructure:"username"`
			Password string `mapstructure:"password"`
		} `mapstructure:"neo4j"`
	} `mapstructure:"datasources"`
	Services struct {
		Notifications string `mapstructure:"notifications"` // address of the service
	} `mapstructure:"services"`
}
