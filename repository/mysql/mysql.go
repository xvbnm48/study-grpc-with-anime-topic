package mysql

import (
	"database/sql"
	"fmt"

	"github.com/pressly/goose"
	"github.com/xvbnm48/study-grpc-with-anime-topic/model"
	"github.com/xvbnm48/study-grpc-with-anime-topic/repository"
	"go.uber.org/zap"
)

type mysqlReadWriter struct {
	db *sql.DB
}

const driverName = "mysql"

func NewMysqlReadWriter(conf model.Configuration) (repository.DatabaseReadWriter, error) {
	confURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s", conf.DBUser, conf.DBPassword, conf.DBHost, conf.DBPort, conf.DBName, conf.DBOptions)

	logger, _ := zap.NewProduction()
	defer logger.Sync()
	logger.Info("Connecting to database", zap.String("url", confURL))

	db, err := sql.Open(driverName, confURL)

	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	err = goose.Up(db, "repository/mysql/migrations")
	if err != nil {
		return nil, err
	}

	return &mysqlReadWriter{db: db}, nil
}

func (a *mysqlReadWriter) Close() error {
	return a.db.Close()
}
