package blogPOA

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
		GetByAuthorEndpoint: MakeGetByAuthorEndpoint(s),
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

type GetByTitleRequest struct {
	Query string
}

type GetByTitleResponse struct {
	BlogPosts []BlogPost `json:"blogposts"`
}

func MakeGetByTitleEndpoint(s BlogPostService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetByTitleRequest)
		blogposts, err := s.GetByTitle(ctx, req.Query)
		return GetByTitleResponse{blogposts}, err
	}
}

type GetByAuthorRequest struct {
	Author string
}

type GetByAuthorResponse struct {
	BlogPosts []BlogPost `json:"blogposts"`
}

func MakeGetByAuthorEndpoint(s BlogPostService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetByAuthorRequest)
		blogpost, err := s.GetByAuthor(ctx, req.Author)
		return GetByAuthorResponse{blogpost}, err
	}
}

type AddRequest struct {
	BlogPost BlogPost
}

type AddResponse struct {
	BlogPost BlogPost `json:"blogpost"`
}

func MakeAddEndpoint(s BlogPostService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(AddRequest)
		blogpost, err := s.Add(ctx, req.BlogPost)
		return AddResponse{blogpost}, err
	}
}

type UpdateRequest struct {
	ID       int
	BlogPost BlogPost
}

type UpdateResponse struct {
}

func MakeUpdateEndpoint(s BlogPostService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateRequest)
		err := s.Update(ctx, req.ID, req.BlogPost)
		return UpdateResponse{}, err
	}
}

type DeleteRequest struct {
	ID int
}

type DeleteResponse struct {
}

func MakeDeleteEndpoint(s BlogPostService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteRequest)
		err := s.Delete(ctx, req.ID)
		return DeleteResponse{}, err
	}
}
