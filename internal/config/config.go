package config

import (
	"flag"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env         string        `yaml:"env" env-default:"local"` //env-default - если не будет указан параметр то значение по умолчанию
	StoragePath string        `yaml:"storage_path"`            //env-required - если не будет указан параметр то программа не запустится
	TokenTTL    time.Duration `yaml:"token_ttl" env-required:"true"`
	GRPC        GRPCConfig    `yaml:"grpc"`
}
type GRPCConfig struct {
	Port    int           `yaml:"port"`
	Timeout time.Duration `yaml:"timeout"`
}

func MustLoad() *Config { //Must используют когда не возвращают ошибку
	path := fetchConfigPath()
	if path == "" {
		panic("config path is empty")
	}
	if _, err := os.Stat(path); os.IsNotExist(err) { //os.Stat- - Проверка существует ли файл, os.IsNotExist-если нет то
		panic("config file does not exist: " + path)
	}
	var cfg Config

	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		panic("failet to read config" + err.Error())
	}
	return &cfg
}

func fetchConfigPath() string {
	var res string

	flag.StringVar(&res, "config", "", "path") //"config" - имя флага (-config) "path" - описание для справки
	flag.Parse()

	if res == "" {
		res = os.Getenv("CONFIG_PATH")
	}
	return res
}

//для запуска
//go run ./cmd/sso/main.go --config=./config/local.yml
