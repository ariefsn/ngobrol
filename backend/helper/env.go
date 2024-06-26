package helper

import (
	"fmt"
	"os"
	"strconv"

	"github.com/ariefsn/ngobrol/logger"
	"github.com/joho/godotenv"
)

type envApp struct {
	Name string
	Host string
	Port string
}

type envDb struct {
	Host     string
	Port     string
	User     string
	Password string
	Db       string
	DbIndex  int
}

type envUrls struct {
	Public string
}

type Env struct {
	App   envApp
	Debug bool
	Mongo envDb
	Urls  envUrls
}

type envValue struct {
	value    string
	fallback interface{}
}

func (e envValue) String() string {
	if e.value == "" && e.fallback != nil {
		return fmt.Sprintf("%s", e.fallback)
	}
	return e.value
}

func (e envValue) Int() int {
	if e.value == "" && e.fallback != nil {
		return e.fallback.(int)
	}
	v, err := strconv.Atoi(e.value)
	if err != nil {
		logger.Error(err)
	}
	return v
}

func (e envValue) Bool() bool {
	if e.value == "" {
		return false
	}

	v, err := strconv.ParseBool(e.value)
	if err != nil {
		logger.Error(err)
	}
	return v
}

var _env *Env

func fromEnv(key string, fallback ...interface{}) envValue {
	var fb interface{}
	if len(fallback) > 0 {
		fb = fallback[0]
	}
	return envValue{
		value:    os.Getenv(key),
		fallback: fb,
	}
}

func InitEnv(envFile ...string) {
	err := godotenv.Load(envFile...)
	if err != nil {
		logger.Warning(err.Error())
	}
	_env = &Env{
		App: envApp{
			Name: fromEnv("APP_NAME", "App").String(),
			Host: fromEnv("APP_HOST", "0.0.0.0").String(),
			Port: fromEnv("APP_PORT", "6001").String(),
		},
		Debug: fromEnv("Debug", true).Bool(),
		Mongo: envDb{
			Host:     fromEnv("MONGO_HOST").String(),
			Port:     fromEnv("MONGO_PORT").String(),
			User:     fromEnv("MONGO_USER").String(),
			Password: fromEnv("MONGO_PASSWORD").String(),
			Db:       fromEnv("MONGO_DB").String(),
		},
		Urls: envUrls{
			Public: fromEnv("URL_PUBLIC").String(),
		},
	}
}

func GetEnv() *Env {
	if _env == nil {
		InitEnv()
	}

	return _env
}
