package repository

import (
	"context"
	"io"

	"github.com/xvbnm48/study-grpc-with-anime-topic/model"
)

type DatabaseReadWriter interface {
	io.Closer
	Add(ctx context.Context, anime model.Anime) error
	Get(ctx context.Context, name string) (model.Anime, error)
}
