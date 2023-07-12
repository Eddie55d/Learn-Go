package database

import (
	"fmt"
	config "ice-cream-app/internal/configs"

	"github.com/pkg/errors"

	"path/filepath"
	"strconv"

	"database/sql"

	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

func ConnectDB() (*sql.DB, error) {

	pathFile := filepath.Join("configs", "configDB.json")

	conf, err := config.LoadConfig(pathFile)
	if err != nil {
		logrus.Error(err)
		return nil, errors.New("config file not load")
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

	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		logrus.Fatal(err)
		return nil, errors.Wrap(err, "database connection fail")
	}

	if err = db.Ping(); err != nil {
		logrus.Error(err)
		return nil, errors.New("database connection fail")
	}
	logrus.Info("Database connection successful.")

	return db, err
}
