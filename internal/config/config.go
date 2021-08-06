package config

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Constants struct {
	Port        string `json:"port"`
	LogLevel    string `json:"log_level"`
	DatabaseURL string `json:"database_url"`
}

type PsqlInstance struct {
	DB *gorm.DB
}

type Config struct {
	Constants
	Psql PsqlInstance
	Log  *logrus.Logger
}

var ServiceConfig Config

func NewServiceConfig() *Config {
	err := initEnv()
	if err != nil {
		log.Println("initEnv err")
	}
	c := Config{}
	// Load constants
	c.Constants = Constants{
		Port:        os.Getenv("PORT"),
		LogLevel:    os.Getenv("LOG_LEVEL"),
		DatabaseURL: os.Getenv("DATABASE_URL"),
	}

	// create logger
	c.Log = logrus.New()
	c.Log.SetFormatter(&logrus.TextFormatter{
		DisableColors: true,
		FullTimestamp: true,
	})
	c.Log.SetOutput(os.Stdout)
	logLvl, err := logrus.ParseLevel(c.Constants.LogLevel)
	if err != nil {
		logLvl = 4
	}
	c.Log.SetLevel(logLvl)

	return &c
}

func getDSN(url string) string {
	var host string
	var user string
	var password string
	var dbname string
	var port string

	s1 := strings.Split(url, "://")
	s2 := strings.Split(s1[1], ":")
	user = s2[0]
	s3 := strings.Split(s2[1], "@")
	password = s3[0]
	host = s3[1]
	s4 := strings.Split(s2[2], "/")
	port = s4[0]
	dbname = s4[1]

	return fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable", host, user, password, dbname, port)
}

func initEnv() error {
	err := godotenv.Load()
	if err != nil {
		log.Println("initEnv > Load err")
		return err
	}
	return nil
}
