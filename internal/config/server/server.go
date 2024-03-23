package server

type ServerList struct {
	Rest Server `yaml:"rest"`
}

type Server struct {
	TLS     bool   `yaml:"tls"`
	Host    string `yaml:"host"`
	Port    string `yaml:"port"`
	Timeout int    `yaml:"timeout"`
}
