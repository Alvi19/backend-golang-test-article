package usecase

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/Alvi19/backend-golang-test/internal/domain"
	"github.com/Alvi19/backend-golang-test/internal/repository"
)

var (
	ErrNotFound     = errors.New("article not found")
	ErrInvalidInput = errors.New("invalid input")
)

type ArticleUsecase interface {
	Create(ctx context.Context, req *CreateArticleRequest) (*domain.Article, error)
	GetByID(ctx context.Context, id uint) (*domain.Article, error)
	List(ctx context.Context, limit, offset int) ([]domain.Article, error)
	Update(ctx context.Context, id uint, req *UpdateArticleRequest) (*domain.Article, error)
	Delete(ctx context.Context, id uint) error
}

type articleUsecase struct {
	repo repository.ArticleRepository
}

func NewArticleUsecase(r repository.ArticleRepository) ArticleUsecase {
	return &articleUsecase{repo: r}
}

// --- Request DTOs ---
type CreateArticleRequest struct {
	Title    string `json:"title" validate:"required,min=20"`
	Content  string `json:"content" validate:"required,min=200"`
	Category string `json:"category" validate:"required,min=3"`
	Status   string `json:"status" validate:"required"`
}

type UpdateArticleRequest struct {
	Title    *string `json:"title" validate:"omitempty,min=20"`
	Content  *string `json:"content" validate:"omitempty,min=200"`
	Category *string `json:"category" validate:"omitempty,min=3"`
	Status   *string `json:"status" validate:"omitempty"`
}

// --- Helpers ---
func normalizeStatus(s string) (domain.ArticleStatus, error) {
	switch strings.ToLower(strings.TrimSpace(s)) {
	case "publish":
		return domain.StatusPublish, nil
	case "draft":
		return domain.StatusDraft, nil
	case "trash":
		return domain.StatusTrash, nil
	default:
		return "", errors.New("invalid status, must be one of: Publish, Draft, Trash")
	}
}

// --- Validation ---
func (u *articleUsecase) validateArticle(req *CreateArticleRequest) error {
	if len(strings.TrimSpace(req.Title)) < 20 {
		return errors.New("title must be at least 20 characters")
	}
	if len(strings.TrimSpace(req.Content)) < 200 {
		return errors.New("content must be at least 200 characters")
	}
	if len(strings.TrimSpace(req.Category)) < 3 {
		return errors.New("category must be at least 3 characters")
	}
	return nil
}

// --- Implementations ---
func (u *articleUsecase) Create(ctx context.Context, req *CreateArticleRequest) (*domain.Article, error) {
	if err := u.validateArticle(req); err != nil {
		return nil, err
	}

	status, err := normalizeStatus(req.Status)
	if err != nil {
		return nil, err
	}

	a := &domain.Article{
		Title:       req.Title,
		Content:     req.Content,
		Category:    req.Category,
		Status:      status,
		CreatedDate: time.Now(),
		UpdatedDate: time.Now(),
	}

	if err := u.repo.Create(ctx, a); err != nil {
		return nil, err
	}
	return a, nil
}

func (u *articleUsecase) GetByID(ctx context.Context, id uint) (*domain.Article, error) {
	a, err := u.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if a == nil {
		return nil, ErrNotFound
	}
	return a, nil
}

func (u *articleUsecase) List(ctx context.Context, limit, offset int) ([]domain.Article, error) {
	return u.repo.List(ctx, limit, offset)
}

func (u *articleUsecase) Update(ctx context.Context, id uint, req *UpdateArticleRequest) (*domain.Article, error) {
	a, err := u.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if a == nil {
		return nil, ErrNotFound
	}

	if req.Title != nil {
		if len(*req.Title) < 20 {
			return nil, errors.New("title must be at least 20 characters")
		}
		a.Title = *req.Title
	}
	if req.Content != nil {
		if len(*req.Content) < 200 {
			return nil, errors.New("content must be at least 200 characters")
		}
		a.Content = *req.Content
	}
	if req.Category != nil {
		if len(*req.Category) < 3 {
			return nil, errors.New("category must be at least 3 characters")
		}
		a.Category = *req.Category
	}
	if req.Status != nil {
		status, err := normalizeStatus(*req.Status)
		if err != nil {
			return nil, err
		}
		a.Status = status
	}

	a.UpdatedDate = time.Now()
	if err := u.repo.Update(ctx, a); err != nil {
		return nil, err
	}
	return a, nil
}

func (u *articleUsecase) Delete(ctx context.Context, id uint) error {
	a, err := u.repo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if a == nil {
		return ErrNotFound
	}
	return u.repo.Delete(ctx, id)
}
