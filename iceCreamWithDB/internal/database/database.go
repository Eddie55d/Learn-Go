package database

import (
	"encoding/json"
	"fmt"

	"os"
	"path/filepath"
	"strconv"

	"database/sql"

	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

type Config struct {
	Database struct {
		Host           string `json:"host"`
		Port           string `json:"port"`
		Name           string `json:"name"`
		User           string `json:"user"`
		Password       string `json:"password"`
		SslMode        string `json:"sslmode"`
		ConnectTimeout string `json:"connect-timeout"`
	} `json:"database"`
}

func loadConfig(file string) (Config, error) {
	var config Config
	configFile, err := os.Open(file)
	if err != nil {
		return config, err
	}
	defer configFile.Close()

	dec := json.NewDecoder(configFile)
	err = dec.Decode(&config)
	return config, err
}

var db *sql.DB

func ConnectDB() *sql.DB {
	var err error
	pathFile := filepath.Join("configs", "configDB.json")

	conf, err := loadConfig(pathFile)
	if err != nil {
		logrus.Error(err)
	}

	portDB, err := strconv.Atoi(conf.Database.Port)
	if err != nil {
		logrus.Error(err)
	}

	connTimeDB, err := strconv.Atoi(conf.Database.ConnectTimeout)
	if err != nil {
		logrus.Error(err)
	}

	dbinfo := fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s sslmode=%s connect_timeout=%d",
		conf.Database.Host, portDB, conf.Database.Name, conf.Database.User, conf.Database.Password, conf.Database.SslMode, connTimeDB)

	db, err = sql.Open("postgres", dbinfo)
	if err != nil {
		logrus.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		logrus.Error(err)
	}
	logrus.Info("Database connection successful.")

	return db
}
