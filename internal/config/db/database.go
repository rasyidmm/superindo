package db

type Databaselist struct {
	Superindo Database `yaml:"superindo"`
}

type Database struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Dbname   string `yaml:"dbname"`
}
