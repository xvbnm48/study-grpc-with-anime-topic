package endpoint

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/xvbnm48/study-grpc-with-anime-topic/model"
	"github.com/xvbnm48/study-grpc-with-anime-topic/service"
)

type AnimeEndpoint struct {
	Add endpoint.Endpoint
	Get endpoint.Endpoint
}

func MakeEndpoint(s service.AnimeService) AnimeEndpoint {
	return AnimeEndpoint{
		Add: AddAnimeEndpoint(s),
		Get: GetAnimeEndpoint(s),
	}
}

// func NewAnimeEndpoints(svc service.AnimeService) AnimeEndpoint {
// 	return AnimeEndpoint{svc}
// }

func AddAnimeEndpoint(svc service.AnimeService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(model.Anime)
		result := svc.AddAnime(ctx, req)
		return result, nil
	}
}

func GetAnimeEndpoint(svc service.AnimeService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(string)
		result, _ := svc.GetAnime(ctx, req)
		return result, nil
	}
}
