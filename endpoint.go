package blog

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type BlogPostEndpoints struct {
	GetAllPostsEndpoint endpoint.Endpoint
	GetByIDEndpoint     endpoint.Endpoint
	GetByTitleEndpoint  endpoint.Endpoint
	GetByAuthorEndpoint endpoint.Endpoint
	AddEndpoint         endpoint.Endpoint
	UpdateEndpoint      endpoint.Endpoint
	DeleteEndpoint      endpoint.Endpoint
}

func MakeBlogPostEndpoints(s BlogPostService) BlogPostEndpoints {
	return BlogPostEndpoints{
		GetAllPostsEndpoint: MakeGetAllPostsEndpoints(s),
		GetByIDEndpoint:     MakeGetByIDEndpoint(s),
		GetByTitleEndpoint:  MakeGetByTitleEndpoint(s),
		GetByAuthorEndpoint: GetByAuthorEndpoint(s),
		AddEndpoint:         MakeAddEndpoint(s),
		UpdateEndpoint:      MakeUpdateEndpoint(s),
		DeleteEndpoint:      MakeDeleteEndpoint(s),
	}
}

type GetAllPostsRequest struct {
}

type GetAllPostsResponse struct {
	BlogPosts []BlogPost `json:"blogposts"`
}

func MakeGetAllPostsEndpoints(s BlogPostService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		blogposts, err := s.GetAllPosts(ctx)
		return GetAllPostsResponse{blogposts}, err
	}
}

type GetByIDRequest struct {
	ID int
}

type GetByIDResponse struct {
	BlogPost BlogPost `json:"blogpost"`
}

func MakeGetByIDEndpoint(s BlogPostService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetByIDRequest)
		blogpost, err := s.GetByID(ctx, req.ID)
		return GetByIDResponse{blogpost}, err
	}
}
