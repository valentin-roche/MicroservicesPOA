package blogPOA

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/transport"
	httptransport "github.com/go-kit/kit/transport/http"
)

var (
	// ErrBadRouting is returned when an expected path variable is missing.
	// It always indicates programmer error.
	ErrBadRouting = errors.New("inconsistent mapping between route and handler (programmer error)")
)

func MakeHTTPHandler(s BlogPostService, logger log.Logger) http.Handler {
	r := mux.NewRouter()
	e := MakeBlogPostEndpoints(s)
	options := []httptransport.ServerOption{
		httptransport.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
		httptransport.ServerErrorEncoder(encodeError),
	}

	// GET		/posts/ 		fetch all posts
	// GET		/posts/:id		fetch post by id
	// GET		/posts/:query 	fetch posts containing title pattern
	// GET		/posts/:author	fetch all post of specified author
	// POST 	/posts/			adds a post
	// PUT  	/posts/:id      update post with this id
	// DELETE 	/posts/:id		deletes post with the given id

	r.Methods("GET").Path("/posts/").Handler(httptransport.NewServer(
		e.GetAllPostsEndpoint,
		decodeGetAllPostsRequest,
		encodeResponse,
		options...,
	))

	r.Methods("GET").Path("/posts/{id}").Handler(httptransport.NewServer(
		e.GetByIDEndpoint,
		decodeGetByIDRequest,
		encodeResponse,
		options...,
	))

	r.Methods("GET").Path("/posts/{query}").Handler(httptransport.NewServer(
		e.GetByTitleEndpoint,
		decodeGetByTitleRequest,
		encodeResponse,
		options...,
	))

	r.Methods("GET").Path("/posts/{author}").Handler(httptransport.NewServer(
		e.GetByAuthorEndpoint,
		decodeGetByAuthorRequest,
		encodeResponse,
		options...,
	))

	r.Methods("POST").Path("/posts/").Handler(httptransport.NewServer(
		e.AddEndpoint,
		decodeAddRequest,
		encodeResponse,
		options...,
	))

	r.Methods("PUT").Path("/posts/update/{id}").Handler(httptransport.NewServer(
		e.UpdateEndpoint,
		decodeUpdateRequest,
		encodeResponse,
		options...,
	))

	r.Methods("DELETE").Path("/posts/{id}").Handler(httptransport.NewServer(
		e.DeleteEndpoint,
		decodeDeleteRequest,
		encodeResponse,
		options...,
	))

	return r
}

func decodeGetAllPostsRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var req GetAllPostsRequest
	return req, nil
}

func decodeGetByIDRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	vars := mux.Vars(r)
	idparam, ok := vars["id"]
	if !ok {
		return nil, ErrBadRouting
	}
	id, err := strconv.Atoi(idparam)
	if err != nil {
		return nil, ErrBadRouting
	}
	return GetByIDRequest{ID: id}, nil
}

func decodeGetByTitleRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	vars := mux.Vars(r)
	query, ok := vars["query"]
	if !ok {
		return nil, ErrBadRouting
	}
	return GetByTitleRequest{Query: query}, nil
}

func decodeAddRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var req AddRequest
	if e := json.NewDecoder(r.Body).Decode(&req.blogPost); e != nil {
		return nil, e
	}
	return req, nil
}

func decodeGetByAuthorRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	vars := mux.Vars(r)
	author, ok := vars["author"]
	if !ok {
		return nil, ErrBadRouting
	}
	return GetByAuthorRequest{Author: author}, nil
}

func decodeUpdateRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var req UpdateRequest
	ePost := json.NewDecoder(r.Body).Decode(&req.BlogPost)
	vars := mux.Vars(r)
	idparam, ok := vars["id"]
	if !ok {
		return nil, ErrBadRouting
	}

	id, err := strconv.Atoi(idparam)
	if err != nil {
		return nil, ErrBadRouting
	}

	if ePost != nil {
		return nil, ePost
	}

	req.ID = id

	return req, nil
}

func decodeDeleteRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	vars := mux.Vars(r)
	idparam, ok := vars["id"]
	if !ok {
		return nil, ErrBadRouting
	}
	id, err := strconv.Atoi(idparam)
	if err != nil {
		return nil, ErrBadRouting
	}
	return DeleteRequest{ID: id}, nil
}

type errorer interface {
	error() error
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if e, ok := response.(errorer); ok && e.error() != nil {
		// Not a Go kit transport error, but a business-logic error.
		// Provide those as HTTP errors.
		encodeError(ctx, e.error(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if err == nil {
		panic("encodeError with nil error")
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(codeFrom(err))
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}

func codeFrom(err error) int {
	switch err {
	case ErrNotFound:
		return http.StatusNotFound
	case ErrNotAnId, ErrNotFound, ErrDifferentId:
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}
