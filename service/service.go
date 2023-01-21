package service

import (
	"context"
	"fmt"
	"log"

	"github.com/xvbnm48/study-grpc-with-anime-topic/model"
	repo "github.com/xvbnm48/study-grpc-with-anime-topic/repository"
)

type service struct {
	logger log.Logger
}

func NewService(logger log.Logger) repo.DatabaseReadWriter {
	return &service{logger}
}

func (s *service) Add(ctx context.Context, anime model.Anime) error {
	fmt.Println(" -- service.Add() -- ")

	return nil
}

func (s *service) Get(ctx context.Context, name string) (model.Anime, error) {
	fmt.Println(" -- service.Get() -- ")
}
