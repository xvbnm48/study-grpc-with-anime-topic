package mysql

import (
	"context"
	"errors"

	"github.com/sirupsen/logrus"
	"github.com/xvbnm48/study-grpc-with-anime-topic/model"
)

var insertAnime = `INSERT INTO anime (name, genre, studio, status) VALUES (?, ?, ?, ?, ?)`

func (a *mysqlReadWriter) Add(ctx context.Context, anime model.Anime) error {
	sqlCheck := `select exists(select 1 from anime where name = ?)`
	err := a.db.QueryRowContext(ctx, insertAnime, anime.Name, anime.Description, anime.Genre, anime.Studio, anime.Year).Scan(sqlCheck)

	if err != nil {
		logrus.SetFormatter(&logrus.JSONFormatter{})
		logrus.WithFields(logrus.Fields{
			"error": err,
		}).Error("Error when inserting anime")
	}

	stmt, err := a.db.PrepareContext(ctx, insertAnime)
	if err != nil {
		logrus.SetFormatter(&logrus.JSONFormatter{})
		logrus.WithFields(logrus.Fields{
			"error": err,
		}).Error("Error when preparing statement")
		return errors.New("Error when preparing statement")
	}
	defer stmt.Close()
	_, err = stmt.ExecContext(ctx, anime.Name, anime.Description, anime.Genre, anime.Studio, anime.Year)
	if err != nil {
		logrus.SetFormatter(&logrus.JSONFormatter{})
		logrus.WithFields(logrus.Fields{
			"error": err,
		}).Error("Error when inserting anime")
		return errors.New("Error when inserting anime")
	}
	return nil
}

func (a *mysqlReadWriter) Get(ctx context.Context, name string) (model.Anime, error) {
	return model.Anime{}, nil
}
