package config

import (
	"encoding/json"
	"fmt"
	"github.com/rasyidmm/superindo-test/internal/config/db"
	"github.com/rasyidmm/superindo-test/internal/config/redis"
	"github.com/rasyidmm/superindo-test/internal/config/server"
	"github.com/spf13/viper"
	"os"
)

type config struct {
	Server   server.ServerList
	Database db.Databaselist
	Redis    redis.Redislist
}

var cfg config

func init() {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	viper.AddConfigPath(dir + "/internal/config/server")
	viper.SetConfigType("yaml")
	viper.SetConfigName("server.yml")
	err = viper.MergeInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error server config file: %w \n", err))
	}

	viper.AddConfigPath(dir + "/internal/config/db")
	viper.SetConfigType("yaml")
	viper.SetConfigName("database.yml")
	err = viper.MergeInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error db config file: %w \n", err))
	}

	viper.AddConfigPath(dir + "/internal/config/redis")
	viper.SetConfigType("yaml")
	viper.SetConfigName("redis.yml")
	err = viper.MergeInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error redis config file: %w \n", err))
	}
	viper.UnmarshalExact(&cfg)

	fmt.Println("=============================")
	dataByte, _ := json.Marshal(cfg)
	fmt.Println(string(dataByte))
	fmt.Println("=============================")
}

func GetConfig() *config {
	return &cfg
}
