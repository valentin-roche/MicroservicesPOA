package blogPOA

import (
	"context"
	"errors"
	"math/rand"
	"strings"
	"sync"
	"time"
)

type BlogPostService interface {
	GetAllPosts(ctx context.Context) ([]BlogPost, error)
	GetByID(ctx context.Context, id int) (BlogPost, error)
	GetByTitle(ctx context.Context, querry string) ([]BlogPost, error)
	GetByAuthor(ctx context.Context, author string) ([]BlogPost, error)
	Add(ctx context.Context, blogpost BlogPost) (BlogPost, error)
	Update(ctx context.Context, id int, blogpost BlogPost) error
	Delete(ctx context.Context, id int) error
}

var (
	ErrNotAnId     = errors.New("Not an ID")
	ErrNotFound    = errors.New("Post not found")
	ErrDifferentId = errors.New("Ids are different")
)

// Methode servant a la creation du service en memoire
func NewInmemBlogPostService() BlogPostService {
	s := &inmemService{
		blogPosts: map[int]BlogPost{},
		nextID:    0,
	}
	rand.Seed(time.Now().UnixNano())

	s.blogPosts[1234] = BlogPost{
		ID:          1234,
		Title:       "Test",
		Author:      "Valentin ROCHE",
		Content:     "Ca marche",
		PublishedOn: time.Now(),
	}

	return s
}

// structure contenant un mutex et les donnees du service
type inmemService struct {
	sync.RWMutex
	blogPosts map[int]BlogPost
	nextID    int
}

func (s *inmemService) GetAllPosts(ctx context.Context) ([]BlogPost, error) {
	s.RLock()
	defer s.RUnlock()

	posts := make([]BlogPost, 0, len(s.blogPosts))
	for _, post := range s.blogPosts {
		posts = append(posts, post)
	}

	return posts, nil
}

func (s *inmemService) GetByID(ctx context.Context, id int) (BlogPost, error) {
	if post, valid := s.blogPosts[id]; valid {
		return post, nil
	}

	return BlogPost{}, ErrNotAnId
}

func (s *inmemService) GetByTitle(ctx context.Context, querry string) ([]BlogPost, error) {
	s.RLock()
	defer s.RUnlock()

	posts := make([]BlogPost, 0, len(s.blogPosts))
	for _, post := range s.blogPosts {
		if strings.ContainsAny(post.Title, querry) {
			posts = append(posts, post)
		}
	}

	return posts, nil
}

func (s *inmemService) GetByAuthor(ctx context.Context, author string) ([]BlogPost, error) {
	s.RLock()
	defer s.RUnlock()

	posts := make([]BlogPost, 0, len(s.blogPosts))
	for _, post := range s.blogPosts {
		if post.Author == author {
			posts = append(posts, post)
		}
	}

	return posts, nil
}

func (s *inmemService) Add(ctx context.Context, blogpost BlogPost) (BlogPost, error) {
	s.RLock()
	defer s.RUnlock()

	blogpost.ID = s.nextID
	s.nextID = s.nextID + 1
	blogpost.PublishedOn = time.Now()

	s.blogPosts[blogpost.ID] = blogpost
	return blogpost, nil
}

func (s *inmemService) Update(ctx context.Context, id int, blogpost BlogPost) error {
	s.RLock()
	defer s.RUnlock()

	if _, valid := s.blogPosts[id]; !valid {
		return ErrNotAnId
	}

	if id != blogpost.ID {
		return ErrDifferentId
	}

	s.blogPosts[id] = blogpost
	return nil
}

func (s *inmemService) Delete(ctx context.Context, id int) error {
	s.RLock()
	defer s.RUnlock()

	if _, valid := s.blogPosts[id]; !valid {
		return ErrNotAnId
	}

	delete(s.blogPosts, id)
	return nil
}
