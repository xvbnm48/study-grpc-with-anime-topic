package repository

import (
	"context"

	"github.com/xvbnm48/study-grpc-with-anime-topic/model"
)

type DatabaseReadWriter interface {
	Add(ctx context.Context, anime model.Anime) error
	Get(ctx context.Context, name string) (model.Anime, error)
}
