package mysql

import (
	"context"

	"github.com/xvbnm48/study-grpc-with-anime-topic/model"
)

func (a *mysqlReadWriter) Add(ctx context.Context, anime model.Anime) error {
	return nil
}

func (a *mysqlReadWriter) Get(ctx context.Context, name string) (model.Anime, error) {
	return model.Anime{}, nil
}
