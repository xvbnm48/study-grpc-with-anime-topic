package service

import (
	"context"

	"github.com/xvbnm48/study-grpc-with-anime-topic/model"
)

type AnimeService interface {
	AddAnime(ctx context.Context, anime model.Anime) error
	GetAnime(ctx context.Context, name string) (model.Anime, error)
}
