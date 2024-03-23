package redis

type Redislist struct {
	Superindo Redis `yaml:"superindo"`
}
type Redis struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}
