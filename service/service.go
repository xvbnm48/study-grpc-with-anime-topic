package service

import (
	repo "github.com/xvbnm48/study-grpc-with-anime-topic/repository"
)

type animeService struct {
	repo repo.DatabaseReadWriter
}

func NewAnimeService(repo repo.DatabaseReadWriter) animeService {
	return animeService{repo: repo}
}
